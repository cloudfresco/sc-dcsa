package tntcontrollers

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	"github.com/cloudfresco/sc-dcsa/internal/config"

	partyproto "github.com/cloudfresco/sc-dcsa/internal/proto-gen/party/v1"
	tntproto "github.com/cloudfresco/sc-dcsa/internal/proto-gen/tnt/v3"
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

// Init the tnt controllers
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

	tntconn, err := grpc.NewClient(grpcServerOpt.GrpcTntServerPort, grpc.WithTransportCredentials(creds), grpc.WithStatsHandler(otelgrpc.NewClientHandler()))
	if err != nil {
		log.Error("Error", zap.Error(err))
		return err
	}

	u := partyproto.NewUserServiceClient(userconn)
	e := tntproto.NewEventServiceClient(tntconn)
	es := tntproto.NewEventSubscriptionServiceClient(tntconn)

	ec := NewEventController(log, u, e)
	esc := NewEventSubscriptionController(log, u, es, h, workflowClient)

	hrlEvent := common.GetHTTPRateLimiter(store, rateOpt.TntMaxRate, rateOpt.TntMaxBurst)
	hrlEventSubscription := common.GetHTTPRateLimiter(store, rateOpt.TntMaxRate, rateOpt.TntMaxBurst)

	mux.Handle("/v3/events", common.AddMiddleware(hrlEvent.RateLimit(ec),
		common.EnsureValidToken(serverOpt.Auth0Audience, serverOpt.Auth0Domain), common.ValidatePermissions([]string{"event:cud", "event:read"}, serverOpt.Auth0Audience, serverOpt.Auth0Domain)))

	mux.Handle("/v3/events/", common.AddMiddleware(hrlEvent.RateLimit(ec),
		common.EnsureValidToken(serverOpt.Auth0Audience, serverOpt.Auth0Domain), common.ValidatePermissions([]string{"event:cud", "event:read"}, serverOpt.Auth0Audience, serverOpt.Auth0Domain)))

	mux.Handle("/v3/event-subscriptions", common.AddMiddleware(hrlEventSubscription.RateLimit(esc),
		common.EnsureValidToken(serverOpt.Auth0Audience, serverOpt.Auth0Domain), common.ValidatePermissions([]string{"eventsub:cud", "eventsub:read"}, serverOpt.Auth0Audience, serverOpt.Auth0Domain)))

	mux.Handle("/v3/event-subscriptions/", common.AddMiddleware(hrlEventSubscription.RateLimit(esc),
		common.EnsureValidToken(serverOpt.Auth0Audience, serverOpt.Auth0Domain), common.ValidatePermissions([]string{"eventsub:cud", "eventsub:read"}, serverOpt.Auth0Audience, serverOpt.Auth0Domain)))

	return nil
}

// InitTest the tnt controllers
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

	tntconn, err := grpc.NewClient(grpcServerOpt.GrpcTntServerPort, grpc.WithTransportCredentials(creds), grpc.WithStatsHandler(otelgrpc.NewClientHandler()))
	if err != nil {
		log.Error("Error", zap.Error(err))
		return err
	}

	u := partyproto.NewUserServiceClient(userconn)
	e := tntproto.NewEventServiceClient(tntconn)
	es := tntproto.NewEventSubscriptionServiceClient(tntconn)

	ec := NewEventController(log, u, e)
	esc := NewEventSubscriptionController(log, u, es, h, workflowClient)

	hrlEvent := common.GetHTTPRateLimiter(store, rateOpt.TntMaxRate, rateOpt.TntMaxBurst)
	hrlEventSubscription := common.GetHTTPRateLimiter(store, rateOpt.TntMaxRate, rateOpt.TntMaxBurst)

	mux.Handle("/v3/events", common.AddMiddleware(hrlEvent.RateLimit(ec),
		common.EnsureValidToken(serverOpt.Auth0Audience, serverOpt.Auth0Domain), common.ValidatePermissions([]string{"event:cud", "event:read"}, serverOpt.Auth0Audience, serverOpt.Auth0Domain)))

	mux.Handle("/v3/events/", common.AddMiddleware(hrlEvent.RateLimit(ec),
		common.EnsureValidToken(serverOpt.Auth0Audience, serverOpt.Auth0Domain), common.ValidatePermissions([]string{"event:cud", "event:read"}, serverOpt.Auth0Audience, serverOpt.Auth0Domain)))

	mux.Handle("/v3/event-subscriptions", common.AddMiddleware(hrlEventSubscription.RateLimit(esc),
		common.EnsureValidToken(serverOpt.Auth0Audience, serverOpt.Auth0Domain), common.ValidatePermissions([]string{"eventsub:cud", "eventsub:read"}, serverOpt.Auth0Audience, serverOpt.Auth0Domain)))

	mux.Handle("/v3/event-subscriptions/", common.AddMiddleware(hrlEventSubscription.RateLimit(esc),
		common.EnsureValidToken(serverOpt.Auth0Audience, serverOpt.Auth0Domain), common.ValidatePermissions([]string{"eventsub:cud", "eventsub:read"}, serverOpt.Auth0Audience, serverOpt.Auth0Domain)))

	return nil
}
