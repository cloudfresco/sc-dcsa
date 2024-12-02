package jitcontrollers

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	"github.com/cloudfresco/sc-dcsa/internal/config"

	jitproto "github.com/cloudfresco/sc-dcsa/internal/protogen/jit/v1"
	partyproto "github.com/cloudfresco/sc-dcsa/internal/protogen/party/v1"
	"github.com/throttled/throttled/v2/store/goredisstore"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.uber.org/cadence/client"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	h              common.WfHelper
	workflowClient client.Client
)

// Init the jit controllers
func Init(log *zap.Logger, rateOpt *config.RateOptions, jwtOpt *config.JWTOptions, mux *http.ServeMux, store *goredisstore.GoRedisStore, serverOpt *config.ServerOptions, grpcServerOpt *config.GrpcServerOptions, uptraceOpt *config.UptraceOptions, configFilePath string) error {
	h.SetupServiceConfig(configFilePath)
	var err error
	workflowClient, err = h.Builder.BuildCadenceClient()
	if err != nil {
		panic(err)
	}

	pwd, _ := os.Getwd()
	keyPath := pwd + filepath.FromSlash(grpcServerOpt.GrpcCaCertPath)
	creds, err := credentials.NewClientTLSFromFile(keyPath, "localhost")
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	userconn, err := grpc.NewClient(grpcServerOpt.GrpcUserServerPort, grpc.WithTransportCredentials(creds), grpc.WithStatsHandler(otelgrpc.NewClientHandler()))
	if err != nil {
		log.Error("Error", zap.Error(err))
		return err
	}

	jitconn, err := grpc.NewClient(grpcServerOpt.GrpcJitServerPort, grpc.WithTransportCredentials(creds), grpc.WithStatsHandler(otelgrpc.NewClientHandler()))
	if err != nil {
		log.Error("Error", zap.Error(err))
		return err
	}

	u := partyproto.NewUserServiceClient(userconn)
	t := jitproto.NewTimestampServiceClient(jitconn)

	tc := NewTimestampController(log, u, t, h, workflowClient)

	hrlTimestamp := common.GetHTTPRateLimiter(store, rateOpt.JitMaxRate, rateOpt.JitMaxBurst)

	mux.Handle("/v0.1/timestamps", common.AddMiddleware(hrlTimestamp.RateLimit(tc),
		common.EnsureValidToken(serverOpt.Auth0Audience, serverOpt.Auth0Domain), common.ValidatePermissions([]string{"ts:cud", "ts:read"}, serverOpt.Auth0Audience, serverOpt.Auth0Domain)))

	mux.Handle("/v0.1/timestamps/", common.AddMiddleware(hrlTimestamp.RateLimit(tc),
		common.EnsureValidToken(serverOpt.Auth0Audience, serverOpt.Auth0Domain), common.ValidatePermissions([]string{"ts:cud", "ts:read"}, serverOpt.Auth0Audience, serverOpt.Auth0Domain)))

	return nil
}

// InitTest the jit controllers
func InitTest(log *zap.Logger, rateOpt *config.RateOptions, jwtOpt *config.JWTOptions, mux *http.ServeMux, store *goredisstore.GoRedisStore, serverOpt *config.ServerOptions, grpcServerOpt *config.GrpcServerOptions, uptraceOpt *config.UptraceOptions, configFilePath string) error {
	h.SetupServiceConfig(configFilePath)
	var err error
	workflowClient, err = h.Builder.BuildCadenceClient()
	if err != nil {
		panic(err)
	}

	pwd, _ := os.Getwd()
	keyPath := filepath.Join(pwd, filepath.FromSlash("/../../../")+filepath.FromSlash(grpcServerOpt.GrpcCaCertPath))
	creds, err := credentials.NewClientTLSFromFile(keyPath, "localhost")
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	userconn, err := grpc.NewClient(grpcServerOpt.GrpcUserServerPort, grpc.WithTransportCredentials(creds), grpc.WithStatsHandler(otelgrpc.NewClientHandler()))
	if err != nil {
		log.Error("Error", zap.Error(err))
		return err
	}

	jitconn, err := grpc.NewClient(grpcServerOpt.GrpcJitServerPort, grpc.WithTransportCredentials(creds), grpc.WithStatsHandler(otelgrpc.NewClientHandler()))
	if err != nil {
		log.Error("Error", zap.Error(err))
		return err
	}

	u := partyproto.NewUserServiceClient(userconn)
	t := jitproto.NewTimestampServiceClient(jitconn)

	tc := NewTimestampController(log, u, t, h, workflowClient)

	hrlTimestamp := common.GetHTTPRateLimiter(store, rateOpt.JitMaxRate, rateOpt.JitMaxBurst)

	mux.Handle("/v0.1/timestamps", common.AddMiddleware(hrlTimestamp.RateLimit(tc),
		common.EnsureValidToken(serverOpt.Auth0Audience, serverOpt.Auth0Domain), common.ValidatePermissions([]string{"ts:cud", "ts:read"}, serverOpt.Auth0Audience, serverOpt.Auth0Domain)))

	mux.Handle("/v0.1/timestamps/", common.AddMiddleware(hrlTimestamp.RateLimit(tc),
		common.EnsureValidToken(serverOpt.Auth0Audience, serverOpt.Auth0Domain), common.ValidatePermissions([]string{"ts:cud", "ts:read"}, serverOpt.Auth0Audience, serverOpt.Auth0Domain)))

	return nil
}
