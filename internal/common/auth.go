package common

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"strings"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/cloudfresco/sc-dcsa/internal/config"
	partyproto "github.com/cloudfresco/sc-dcsa/internal/protogen/party/v1"
	"github.com/rs/cors"
	"github.com/unrolled/secure"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

type Auth0Config struct {
	Port          string
	SecureOptions secure.Options
	CorsOptions   cors.Options
	Audience      string
	Domain        string
}

var log *zap.Logger

// DBMysql for DbType is mysql
const DBMysql string = "mysql"

// DBPgsql for DbType is pgsql
const DBPgsql string = "pgsql"

var jwtOpt *config.JWTOptions

// SetJWTOpt set JWT opt used in auth middleware
func SetJWTOpt(jwt *config.JWTOptions) {
	jwtOpt = jwt
}

// GetJWTOpt get JWT opt used in auth middleware
func GetJWTOpt() *config.JWTOptions {
	return jwtOpt
}

// GetAuthUserDetailsResponse - details of a user stored in the Redis cache
type GetAuthUserDetailsResponse struct {
	Email  string
	UserID string
	Roles  []string
}

// Key - type of the key used in the request context
type Key string

// KeyEmailToken - used for the request context key
const KeyEmailToken Key = "emailtoken"

// ContextStruct - stored in the request context
// set in AuthMiddleware
type ContextStruct struct {
	Email       string
	TokenString string
}

// GetAuthBearerToken - extract the BEARER token from the auth header
func GetAuthBearerToken(r *http.Request) (string, error) {
	var APIkey string
	bearer := r.Header.Get("authorization")
	if len(bearer) > 7 && strings.ToUpper(bearer[0:6]) == "BEARER" {
		APIkey = bearer[7:]
	} else {
		log.Error("Error",
			zap.Int("msgnum", 252),
			zap.Error(errors.New("APIkey Not Found")))
		return "", errors.New("APIkey Not Found ")
	}
	return APIkey, nil
}

// GetAuthData - used to get auth details
func GetAuthData(r *http.Request) ContextStruct {
	data := r.Context().Value(KeyEmailToken).(ContextStruct)
	return data
}

// GetJWTFromCtx - used to get jwt from context
func GetJWTFromCtx(ctx context.Context, header string) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Error("Error", zap.Error(errors.New(`no headers in request`)))
		return "", errors.New("no headers in request")
	}

	authHeaders, ok := md[header]
	if !ok {
		log.Error("Error", zap.Error(errors.New(`no headers in request`)))
		return "", errors.New("no header in request")
	}

	if len(authHeaders) != 1 {
		log.Error("Error", zap.Error(errors.New(`more than 1 header in request`)))
		return "", errors.New("more than 1 header in request")
	}
	return authHeaders[0], nil
}

// CreateCtxJWT - used to get context
func CreateCtxJWT(ctx context.Context) (context.Context, error) {
	auth, err := GetJWTFromCtx(ctx, "authorization")
	if err != nil {
		log.Error("Error", zap.Error(err))
		return ctx, err
	}
	md := metadata.Pairs("authorization", auth)
	newCtx := metadata.NewOutgoingContext(ctx, md)
	return newCtx, nil
}

// GetProtoMd - used to get auth details and context
func GetProtoMd(r *http.Request, email, tokenString string) (context.Context, partyproto.GetAuthUserDetailsRequest) {
	cdata := partyproto.GetAuthUserDetailsRequest{}
	cdata.TokenString = tokenString
	cdata.Email = email
	cdata.RequestUrlPath = r.URL.Path
	cdata.RequestMethod = r.Method
	md := metadata.Pairs("authorization", "Bearer "+cdata.TokenString)

	ctx := metadata.NewOutgoingContext(r.Context(), md)
	return ctx, cdata
}

func ValidatePermissions(w http.ResponseWriter, r *http.Request, expectedClaims []string, audience string, domain string) error {
	tokenString, err := getToken(r)
	if err != nil {
		http.Error(w, "Error parsing token", http.StatusUnauthorized)
		return errors.New("Error parsing token")
	}
	_, claims, err := getClaims(audience, domain, tokenString, w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return errors.New("Error parsing token")
	}

	if len(claims.Permissions) == 0 {
		return errors.New("Permission Denied")
	}

	if !claims.HasPermissions(expectedClaims) {
		return errors.New("Permission Denied")
	}
	return nil
}

func getClaims(audience string, domain string, tokenString string, w http.ResponseWriter, r *http.Request) (*jwtmiddleware.JWTMiddleware, *CustomClaims, error) {
	issuerURL, err := url.Parse("https://" + domain + "/")
	if err != nil {
		return nil, nil, errors.New("Failed to parse the issuer url")
	}

	provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

	jwtValidator, err := validator.New(
		provider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		[]string{audience},
		validator.WithCustomClaims(func() validator.CustomClaims {
			return new(CustomClaims)
		}),
	)
	if err != nil {
		return nil, nil, errors.New("Failed to set up the jwt validator")
	}

	errorHandler := func(w http.ResponseWriter, r *http.Request, err error) {
		if errors.Is(err, jwtmiddleware.ErrJWTMissing) {
			http.Error(w, "Error parsing token", http.StatusUnauthorized)
			return
		}
		if errors.Is(err, jwtmiddleware.ErrJWTInvalid) {
			http.Error(w, "Error parsing token", http.StatusUnauthorized)
			return
		}
	}

	middleware := jwtmiddleware.New(
		jwtValidator.ValidateToken,
		jwtmiddleware.WithErrorHandler(errorHandler),
	)

	tokenClaims, err := jwtValidator.ValidateToken(r.Context(), tokenString)
	if err != nil {
		return nil, nil, err
	}
	m := tokenClaims.(*validator.ValidatedClaims)

	claims := m.CustomClaims.(*CustomClaims)

	return middleware, claims, nil
}

func SetEmailToken(req *http.Request, tokenString string, email string) *http.Request {
	req.Header.Set("Authorization", "Bearer "+tokenString)
	req.Header.Set("X-User-Email", email)
	req.Header.Set("X-Auth-Token", tokenString)
	return req
}

func GetEmailToken(req *http.Request) (string, string) {
	email := req.Header.Get("X-User-Email")
	token := req.Header.Get("X-Auth-Token")
	return email, token
}

func GetContextAuthUser(w http.ResponseWriter, r *http.Request, permissions []string, audience string, domain string, userServiceClient partyproto.UserServiceClient) (context.Context, *partyproto.GetAuthUserDetailsResponse, string, error) {
	err := ValidatePermissions(w, r, permissions, audience, domain)
	if err != nil {
		return nil, nil, "", err
	}

	email, token := GetEmailToken(r)

	ctx, cdata := GetProtoMd(r, email, token)

	user, err := userServiceClient.GetAuthUserDetails(ctx, &cdata)
	if err != nil {
		return nil, nil, "", err
	}

	return ctx, user, token, nil
}
