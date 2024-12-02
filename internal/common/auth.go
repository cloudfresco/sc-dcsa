package common

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/cloudfresco/sc-dcsa/internal/config"
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
	fmt.Println("auth.go GetJWTFromCtx header", header)
	fmt.Println("auth.go GetJWTFromCtx ctx", ctx)
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
	fmt.Println("auth.go CreateCtxJWT")
	fmt.Println("auth.go CreateCtxJWT ctx", ctx)
	auth, err := GetJWTFromCtx(ctx, "authorization")
	fmt.Println("auth.go CreateCtxJWT auth", auth)
	if err != nil {
		log.Error("Error", zap.Error(err))
		return ctx, err
	}
	md := metadata.Pairs("authorization", auth)
	fmt.Println("auth.go CreateCtxJWT md", md)
	newCtx := metadata.NewOutgoingContext(ctx, md)
	fmt.Println("auth.go CreateCtxJWT newCtx", newCtx)
	return newCtx, nil
}
