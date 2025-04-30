package partyservices

import (
	"context"
	"encoding/json"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	"github.com/cloudfresco/sc-dcsa/internal/config"
	commonproto "github.com/cloudfresco/sc-dcsa/internal/protogen/common/v1"
	partyproto "github.com/cloudfresco/sc-dcsa/internal/protogen/party/v1"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/contrib/propagators/aws/xray"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/encoding/gzip"
)

// UserService - For accessing user services
type UserService struct {
	log           *zap.Logger
	DBService     *common.DBService
	RedisService  *common.RedisService
	MailerService common.MailerIntf
	JWTOptions    *config.JWTOptions
	UserOptions   *config.UserOptions
	ServerOptions *config.ServerOptions
	partyproto.UnimplementedUserServiceServer
}

// NewUserService - Create User Service
func NewUserService(log *zap.Logger, dbOpt *common.DBService, redisOpt *common.RedisService, mailerOpt common.MailerIntf, jwtOptions *config.JWTOptions, userOpt *config.UserOptions, serverOpt *config.ServerOptions) *UserService {
	return &UserService{
		log:           log,
		DBService:     dbOpt,
		RedisService:  redisOpt,
		MailerService: mailerOpt,
		JWTOptions:    jwtOptions,
		UserOptions:   userOpt,
		ServerOptions: serverOpt,
	}
}

var log *zap.Logger

// StartUserServer - Start User Server
func StartUserServer(log *zap.Logger, isTest bool, pwd string, dbOpt *config.DBOptions, redisOpt *config.RedisOptions, mailerOpt *config.MailerOptions, serverOpt *config.ServerOptions, grpcServerOpt *config.GrpcServerOptions, jwtOpt *config.JWTOptions, oauthOpt *config.OauthOptions, userOpt *config.UserOptions, uptraceOpt *config.UptraceOptions, dbService *common.DBService, redisService *common.RedisService, mailerService common.MailerIntf) {
	common.SetJWTOpt(jwtOpt)

	creds, err := common.GetSrvCred(log, isTest, pwd, grpcServerOpt)
	if err != nil {
		os.Exit(1)
	}

	// tracer, _ := interceptors.NewJaegerTracer(log, jaegerTracerOpt, jaegerTracerOpt.UserServiceName)

	var srvOpts []grpc.ServerOption

	srvOpts = append(srvOpts, grpc.Creds(creds))

	srvOpts = append(srvOpts, grpc.StatsHandler(otelgrpc.NewServerHandler()))

	userService := NewUserService(log, dbService, redisService, mailerService, jwtOpt, userOpt, serverOpt)

	lis, err := net.Listen("tcp", grpcServerOpt.GrpcUserServerPort)
	if err != nil {
		log.Error("Error", zap.Int("msgnum", 9109), zap.Error(err))
		os.Exit(1)
	}

	// Create a HTTP server for prometheus.

	srv := grpc.NewServer(srvOpts...)
	partyproto.RegisterUserServiceServer(srv, userService)

	if err := srv.Serve(lis); err != nil {
		log.Error("Error", zap.Int("msgnum", 9414), zap.Error(err))
		os.Exit(1)
	}
}

// GetUsers - Get users
func (u *UserService) GetUsers(ctx context.Context, in *partyproto.GetUsersRequest) (*partyproto.GetUsersResponse, error) {
	url := "https://" + u.ServerOptions.Auth0Domain + "/api/v2/users"
	respBody, err := common.SendRequest("GET", url, nil, "Bearer "+u.ServerOptions.Auth0MgmtToken)
	if err != nil {
		return nil, err
	}

	jsonDataReader := strings.NewReader(string(respBody))
	decoder := json.NewDecoder(jsonDataReader)
	var userResp []map[string]interface{}
	err = decoder.Decode(&userResp)
	if err != nil {
		return nil, err
	}
	users := []*partyproto.User{}
	for _, userResult := range userResp {
		user := partyproto.User{}
		user.Id = userResult["user_id"].(string)
		user.Email = userResult["email"].(string)
		user.Picture = userResult["picture"].(string)
		user.Name = userResult["name"].(string)
		users = append(users, &user)
	}

	usersResponse := partyproto.GetUsersResponse{}
	usersResponse.Users = users
	return &usersResponse, nil
}

// GetUserByEmail - Get user details by email
func (u *UserService) GetUserByEmail(ctx context.Context, in *partyproto.GetUserByEmailRequest) (*partyproto.GetUserByEmailResponse, error) {
	url := "https://" + u.ServerOptions.Auth0Domain + "/api/v2/users-by-email?email=" + in.Email
	respBody, err := common.SendRequest("GET", url, nil, "Bearer "+u.ServerOptions.Auth0MgmtToken)
	if err != nil {
		return nil, err
	}

	jsonDataReader := strings.NewReader(string(respBody))
	decoder := json.NewDecoder(jsonDataReader)
	var userResp []map[string]interface{}
	err = decoder.Decode(&userResp)

	userResult := userResp[0]
	if err != nil {
		return nil, err
	}

	user := partyproto.User{}
	user.Id = userResult["user_id"].(string)
	user.Email = userResult["email"].(string)
	user.Picture = userResult["picture"].(string)
	user.Name = userResult["name"].(string)

	userResponse := partyproto.GetUserByEmailResponse{}
	userResponse.User = &user
	return &userResponse, nil
}

// GetUser - used to get user by Id
func (u *UserService) GetUser(ctx context.Context, inReq *partyproto.GetUserRequest) (*partyproto.GetUserResponse, error) {
	in := inReq.GetRequest
	url := "https://" + u.ServerOptions.Auth0Domain + "/api/v2/users/" + in.Id
	respBody, err := common.SendRequest("GET", url, nil, "Bearer "+u.ServerOptions.Auth0MgmtToken)
	if err != nil {
		return nil, err
	}

	jsonDataReader := strings.NewReader(string(respBody))
	decoder := json.NewDecoder(jsonDataReader)
	var userResp map[string]interface{}
	err = decoder.Decode(&userResp)
	if err != nil {
		return nil, err
	}

	user := partyproto.User{}
	user.Id = userResp["user_id"].(string)
	user.Email = userResp["email"].(string)
	user.Picture = userResp["picture"].(string)
	user.Name = userResp["name"].(string)

	userResponse := partyproto.GetUserResponse{}
	userResponse.User = &user
	return &userResponse, nil
}

// ChangePassword - used to update password
func (u *UserService) ChangePassword(ctx context.Context, in *partyproto.ChangePasswordRequest) (*partyproto.ChangePasswordResponse, error) {
	url := "https://" + u.ServerOptions.Auth0Domain + "/dbconnections/change_password"

	payload := strings.NewReader(`{"client_id":"` + u.ServerOptions.Auth0ClientId + `","email":"` + in.Email + `", "connection":"` + u.ServerOptions.Auth0Connection + `"}`)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/json")
	req.Header.Add("cache-control", "no-cache")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	return &partyproto.ChangePasswordResponse{}, nil
}

// DeleteUser - used to get user by Id
func (u *UserService) DeleteUser(ctx context.Context, in *partyproto.DeleteUserRequest) (*partyproto.DeleteUserResponse, error) {
	url := "https://" + u.ServerOptions.Auth0Domain + "/api/v2/users/" + in.UserId
	_, err := common.SendRequest("DELETE", url, nil, "Bearer "+u.ServerOptions.Auth0MgmtToken)
	if err != nil {
		return nil, err
	}

	return &partyproto.DeleteUserResponse{}, nil
}

// RefreshToken - used for Refresh Token
// this needs to convert to auth0 code
func (u *UserService) RefreshToken(ctx context.Context, form *partyproto.RefreshTokenRequest) (*partyproto.RefreshTokenResponse, error) {
	return nil, nil
}

func (u *UserService) UpdateUser(ctx context.Context, form *partyproto.UpdateUserRequest) (*partyproto.UpdateUserResponse, error) {
	return nil, nil
}

// GetAuthUserDetails - used to get auth user details
func (u *UserService) GetAuthUserDetails(ctx context.Context, in *partyproto.GetAuthUserDetailsRequest) (*partyproto.GetAuthUserDetailsResponse, error) {
	resp, err := u.RedisService.Get(in.TokenString)
	if err != nil {
		u.log.Error("Error", zap.Int("msgnum", 1583), zap.Error(err))
	}

	v := partyproto.GetAuthUserDetailsResponse{}
	if resp == "" {
		url := "https://" + u.ServerOptions.Auth0Domain + "/userinfo"

		respBody, err := common.SendRequest("GET", url, nil, "Bearer "+in.TokenString)
		if err != nil {
			return nil, err
		}

		jsonDataReader := strings.NewReader(string(respBody))
		decoder := json.NewDecoder(jsonDataReader)
		var userDetail map[string]interface{}
		err = decoder.Decode(&userDetail)
		if err != nil {
			return nil, err
		}

		v.Email = userDetail["email"].(string)
		v.UserId = userDetail["sub"].(string)

	} else {

		err = json.Unmarshal([]byte(resp), &v)
		if err != nil {
			u.log.Error("Error", zap.Int("msgnum", 266), zap.Error(err))
		}

	}

	v.RequestId = common.GetRequestID()

	return &v, nil
}

// CreateRole - used to create Role
func (u *UserService) CreateRole(ctx context.Context, in *partyproto.CreateRoleRequest) (*partyproto.CreateRoleResponse, error) {
	inReq := in.CreateRole
	inReq.Auth0Domain = u.ServerOptions.Auth0Domain
	inReq.Auth0MgmtToken = u.ServerOptions.Auth0MgmtToken
	role, err := common.CreateRoleResp(ctx, inReq)
	if err != nil {
		log.Error("Error", zap.String("user", inReq.UserEmail), zap.String("reqid", inReq.RequestId), zap.Error(err))
		return nil, err
	}
	roleResponse := partyproto.CreateRoleResponse{}
	roleResponse.Role = role

	return &partyproto.CreateRoleResponse{}, nil
}

// GetRole - used to get Role
func (u *UserService) GetRole(ctx context.Context, in *partyproto.GetRoleRequest) (*partyproto.GetRoleResponse, error) {
	inReq := in.GetRole
	inReq.Auth0Domain = u.ServerOptions.Auth0Domain
	inReq.Auth0MgmtToken = u.ServerOptions.Auth0MgmtToken
	role, err := common.GetRoleResp(ctx, inReq)
	if err != nil {
		log.Error("Error", zap.String("user", inReq.UserEmail), zap.String("reqid", inReq.RequestId), zap.Error(err))
		return nil, err
	}
	roleResponse := partyproto.GetRoleResponse{}
	roleResponse.Role = role
	return &roleResponse, nil
}

// DeleteRole - used to delete Role
func (u *UserService) DeleteRole(ctx context.Context, in *partyproto.DeleteRoleRequest) (*partyproto.DeleteRoleResponse, error) {
	inReq := in.DeleteRole
	inReq.Auth0Domain = u.ServerOptions.Auth0Domain
	inReq.Auth0MgmtToken = u.ServerOptions.Auth0MgmtToken
	err := common.DeleteRoleResp(ctx, inReq)
	if err != nil {
		log.Error("Error", zap.String("user", inReq.UserEmail), zap.String("reqid", inReq.RequestId), zap.Error(err))
		return nil, err
	}
	return &partyproto.DeleteRoleResponse{}, nil
}

// UpdateRole - used to update Role
func (u *UserService) UpdateRole(ctx context.Context, in *partyproto.UpdateRoleRequest) (*partyproto.UpdateRoleResponse, error) {
	inReq := in.UpdateRole
	inReq.Auth0Domain = u.ServerOptions.Auth0Domain
	inReq.Auth0MgmtToken = u.ServerOptions.Auth0MgmtToken
	role, err := common.UpdateRoleResp(ctx, inReq)
	if err != nil {
		log.Error("Error", zap.String("user", inReq.UserEmail), zap.String("reqid", inReq.RequestId), zap.Error(err))
		return nil, err
	}
	roleResponse := partyproto.UpdateRoleResponse{}
	roleResponse.Role = role

	return &roleResponse, nil
}

// GetRoles - used to get Roles
func (u *UserService) GetRoles(ctx context.Context, in *partyproto.GetRolesRequest) (*partyproto.GetRolesResponse, error) {
	inReq := in.GetRoles
	inReq.Auth0Domain = u.ServerOptions.Auth0Domain
	inReq.Auth0MgmtToken = u.ServerOptions.Auth0MgmtToken
	roles, err := common.GetRolesResp(ctx, inReq)
	if err != nil {
		log.Error("Error", zap.String("user", inReq.UserEmail), zap.String("reqid", inReq.RequestId), zap.Error(err))
		return nil, err
	}
	roleResponse := partyproto.GetRolesResponse{}
	roleResponse.Roles = roles
	return &roleResponse, nil
}

func (u *UserService) AddPermisionsToRoles(ctx context.Context, in *partyproto.AddPermisionsToRolesRequest) (*partyproto.AddPermisionsToRolesResponse, error) {
	inReq := in.AddPermisionsToRoles
	inReq.Auth0Domain = u.ServerOptions.Auth0Domain
	inReq.Auth0MgmtToken = u.ServerOptions.Auth0MgmtToken
	err := common.AddPermisionsToRolesResp(ctx, inReq)
	if err != nil {
		log.Error("Error", zap.String("user", inReq.UserEmail), zap.String("reqid", inReq.RequestId), zap.Error(err))
		return nil, err
	}

	return &partyproto.AddPermisionsToRolesResponse{}, nil
}

func (u *UserService) RemoveRolePermission(ctx context.Context, in *partyproto.RemoveRolePermissionRequest) (*partyproto.RemoveRolePermissionResponse, error) {
	inReq := in.RemoveRolePermission
	inReq.Auth0Domain = u.ServerOptions.Auth0Domain
	inReq.Auth0MgmtToken = u.ServerOptions.Auth0MgmtToken
	err := common.RemoveRolePermissionResp(ctx, inReq)
	if err != nil {
		log.Error("Error", zap.String("user", inReq.UserEmail), zap.String("reqid", inReq.RequestId), zap.Error(err))
		return nil, err
	}

	return &partyproto.RemoveRolePermissionResponse{}, nil
}

func (u *UserService) GetRolePermissions(ctx context.Context, in *partyproto.GetRolePermissionsRequest) (*partyproto.GetRolePermissionsResponse, error) {
	inReq := in.GetRolePermissions
	inReq.Auth0Domain = u.ServerOptions.Auth0Domain
	inReq.Auth0MgmtToken = u.ServerOptions.Auth0MgmtToken
	rolePermissions, err := common.GetRolePermissionsResp(ctx, inReq)
	if err != nil {
		log.Error("Error", zap.String("user", inReq.UserEmail), zap.String("reqid", inReq.RequestId), zap.Error(err))
		return nil, err
	}

	rolePermissionsResponse := partyproto.GetRolePermissionsResponse{}
	rolePermissionsResponse.RolePermissions = rolePermissions
	return &rolePermissionsResponse, nil
}

// AssignRolesToUsers - used to assign roles to users
func (u *UserService) AssignRolesToUsers(ctx context.Context, in *partyproto.AssignRolesToUsersRequest) (*partyproto.AssignRolesToUsersResponse, error) {
	inReq := in.AssignRolesToUsers
	inReq.Auth0Domain = u.ServerOptions.Auth0Domain
	inReq.Auth0MgmtToken = u.ServerOptions.Auth0MgmtToken
	err := common.AssignRolesToUsersResp(ctx, inReq)
	if err != nil {
		log.Error("Error", zap.String("user", inReq.UserEmail), zap.String("reqid", inReq.RequestId), zap.Error(err))
		return nil, err
	}
	return &partyproto.AssignRolesToUsersResponse{}, nil
}

// ViewUserRoles - used to View User Roles
func (u *UserService) ViewUserRoles(ctx context.Context, in *partyproto.ViewUserRolesRequest) (*partyproto.ViewUserRolesResponse, error) {
	inReq := in.ViewUserRoles
	inReq.Auth0Domain = u.ServerOptions.Auth0Domain
	inReq.Auth0MgmtToken = u.ServerOptions.Auth0MgmtToken
	roles, err := common.ViewUserRolesResp(ctx, inReq)
	if err != nil {
		log.Error("Error", zap.String("user", inReq.UserEmail), zap.String("reqid", inReq.RequestId), zap.Error(err))
		return nil, err
	}
	roleResponse := partyproto.ViewUserRolesResponse{}
	roleResponse.Roles = roles
	return &roleResponse, nil
}

func (u *UserService) GetRoleUsers(ctx context.Context, in *partyproto.GetRoleUsersRequest) (*partyproto.GetRoleUsersResponse, error) {
	url := "https://" + u.ServerOptions.Auth0Domain + "/api/v2/roles/" + in.RoleId + "/users"

	respBody, err := common.SendRequest("GET", url, nil, "Bearer "+u.ServerOptions.Auth0MgmtToken)
	if err != nil {
		return nil, err
	}

	jsonDataReader := strings.NewReader(string(respBody))
	decoder := json.NewDecoder(jsonDataReader)
	var roleUsers []map[string]interface{}
	err = decoder.Decode(&roleUsers)
	if err != nil {
		return nil, err
	}
	users := []*partyproto.User{}
	for _, usr := range roleUsers {
		user := partyproto.User{}
		user.Id = usr["user_id"].(string)
		user.Email = usr["email"].(string)
		user.Picture = usr["picture"].(string)
		user.Name = usr["name"].(string)
		users = append(users, &user)
	}
	roleUsersResponse := partyproto.GetRoleUsersResponse{}
	roleUsersResponse.Users = users
	return &roleUsersResponse, nil
}

// AddAPIPermission - used to create permission
func (u *UserService) AddAPIPermission(ctx context.Context, in *partyproto.AddAPIPermissionRequest) (*partyproto.AddAPIPermissionResponse, error) {
	inReq := commonproto.AddAPIPermission{}
	inReq.Permissions = in.Permissions
	inReq.Auth0Domain = u.ServerOptions.Auth0Domain
	inReq.Auth0MgmtToken = u.ServerOptions.Auth0MgmtToken
	inReq.Auth0ApiId = u.ServerOptions.Auth0ApiId
	err := common.AddAPIPermissionResp(ctx, &inReq)
	if err != nil {
		log.Error("Error", zap.String("user", inReq.UserEmail), zap.String("reqid", inReq.RequestId), zap.Error(err))
		return nil, err
	}

	return &partyproto.AddAPIPermissionResponse{}, nil
}

// GetUserWithNewContext - GetUserWithNewContext
func GetUserWithNewContext(ctx context.Context, userId string, userEmail string, requestId string, userServiceClient partyproto.UserServiceClient) (*partyproto.User, error) {
	getRequest := commonproto.GetRequest{}
	getRequest.Id = userId
	getRequest.UserEmail = userEmail
	getRequest.RequestId = requestId

	ctxNew, err := common.CreateCtxJWT(ctx)
	if err != nil {
		log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestId), zap.Int("msgnum", 4319), zap.Error(err))
		return nil, err
	}

	form := partyproto.GetUserRequest{}
	form.GetRequest = &getRequest
	userResponse, err := userServiceClient.GetUser(ctxNew, &form)
	if err != nil {
		log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestId), zap.Int("msgnum", 4319), zap.Error(err))
		return nil, err
	}
	user := userResponse.User
	return user, nil
}

func (u *UserService) UserTracer(ctx context.Context) {
	dsn := ""
	if dsn == "" {
		panic("UPTRACE_DSN environment variable is required")
	}

	creds := credentials.NewClientTLSFromCert(nil, "")
	exporter, err := otlptracegrpc.New(
		ctx,
		otlptracegrpc.WithEndpoint("otlp.uptrace.dev:4317"),
		otlptracegrpc.WithTLSCredentials(creds),
		otlptracegrpc.WithHeaders(map[string]string{
			// Set the Uptrace DSN here or use UPTRACE_DSN env var.
			"uptrace-dsn": dsn,
		}),
		otlptracegrpc.WithCompressor(gzip.Name),
	)
	if err != nil {
		panic(err)
	}

	bsp := sdktrace.NewBatchSpanProcessor(exporter,
		sdktrace.WithMaxQueueSize(10_000),
		sdktrace.WithMaxExportBatchSize(10_000))
	// Call shutdown to flush the buffers when program exits.
	defer bsp.Shutdown(ctx)

	resource, err := resource.New(ctx,
		resource.WithFromEnv(),
		resource.WithTelemetrySDK(),
		resource.WithHost(),
		resource.WithAttributes(
			attribute.String("service.name", "myservice"),
			attribute.String("service.version", "1.0.0"),
		))
	if err != nil {
		panic(err)
	}

	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithResource(resource),
		sdktrace.WithIDGenerator(xray.NewIDGenerator()),
	)
	tracerProvider.RegisterSpanProcessor(bsp)

	// Install our tracer provider and we are done.
	otel.SetTracerProvider(tracerProvider)

	tracer := otel.Tracer("myservice")

	ctx, span := tracer.Start(ctx, "UserTracer",
		trace.WithAttributes(attribute.String("extra.key", "extra.value")))
	defer span.End()
}
