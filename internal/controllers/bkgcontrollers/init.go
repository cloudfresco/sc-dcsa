package bkgcontrollers

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	"github.com/cloudfresco/sc-dcsa/internal/config"
	bkgproto "github.com/cloudfresco/sc-dcsa/internal/proto-gen/bkg/v2"
	partyproto "github.com/cloudfresco/sc-dcsa/internal/proto-gen/party/v1"
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

// Init the bkg controllers
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

	bkgconn, err := grpc.NewClient(grpcServerOpt.GrpcBkgServerPort, grpc.WithTransportCredentials(creds), grpc.WithStatsHandler(otelgrpc.NewClientHandler()))
	if err != nil {
		log.Error("Error", zap.Error(err))
		return err
	}

	u := partyproto.NewUserServiceClient(userconn)
	b := bkgproto.NewBkgServiceClient(bkgconn)
	bss := bkgproto.NewBkgShipmentSummaryServiceClient(bkgconn)
	bsm := bkgproto.NewBkgSummaryServiceClient(bkgconn)

	bc := NewBkgController(log, u, b, h, workflowClient)
	bssc := NewBkgShipmentSummaryController(log, u, bss)
	bsmc := NewBkgSummaryController(log, u, bsm)

	hrlBkg := common.GetHTTPRateLimiter(store, rateOpt.BkgMaxRate, rateOpt.BkgMaxBurst)
	hrlBkgShipmentSummary := common.GetHTTPRateLimiter(store, rateOpt.BkgMaxRate, rateOpt.BkgMaxBurst)
	hrlBkgSummary := common.GetHTTPRateLimiter(store, rateOpt.BkgMaxRate, rateOpt.BkgMaxBurst)

	mux.Handle("/v2/bookings", common.AddMiddleware(hrlBkg.RateLimit(bc),
		common.EnsureValidToken(serverOpt.Auth0Audience, serverOpt.Auth0Domain), common.ValidatePermissions([]string{"bkg:cud", "bkg:read"}, serverOpt.Auth0Audience, serverOpt.Auth0Domain)))

	mux.Handle("/v2/bookings/", common.AddMiddleware(hrlBkg.RateLimit(bc),
		common.EnsureValidToken(serverOpt.Auth0Audience, serverOpt.Auth0Domain), common.ValidatePermissions([]string{"bkg:cud", "bkg:read"}, serverOpt.Auth0Audience, serverOpt.Auth0Domain)))

	mux.Handle("/v2/bkg-shipment-summaries", common.AddMiddleware(hrlBkgShipmentSummary.RateLimit(bssc),
		common.EnsureValidToken(serverOpt.Auth0Audience, serverOpt.Auth0Domain), common.ValidatePermissions([]string{"bkgss:cud", "bkgss:read"}, serverOpt.Auth0Audience, serverOpt.Auth0Domain)))

	mux.Handle("/v2/bkg-shipment-summaries/", common.AddMiddleware(hrlBkgShipmentSummary.RateLimit(bssc),
		common.EnsureValidToken(serverOpt.Auth0Audience, serverOpt.Auth0Domain), common.ValidatePermissions([]string{"bkgss:cud", "bkgss:read"}, serverOpt.Auth0Audience, serverOpt.Auth0Domain)))

	mux.Handle("/v2/booking-summaries", common.AddMiddleware(hrlBkgSummary.RateLimit(bsmc),
		common.EnsureValidToken(serverOpt.Auth0Audience, serverOpt.Auth0Domain), common.ValidatePermissions([]string{"bkgs:cud", "bkgs:read"}, serverOpt.Auth0Audience, serverOpt.Auth0Domain)))

	mux.Handle("/v2/booking-summaries/", common.AddMiddleware(hrlBkgSummary.RateLimit(bsmc),
		common.EnsureValidToken(serverOpt.Auth0Audience, serverOpt.Auth0Domain), common.ValidatePermissions([]string{"bkgs:cud", "bkgs:read"}, serverOpt.Auth0Audience, serverOpt.Auth0Domain)))

	return nil
}

// InitTest the user controllers
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

	bkgconn, err := grpc.NewClient(grpcServerOpt.GrpcBkgServerPort, grpc.WithTransportCredentials(creds), grpc.WithStatsHandler(otelgrpc.NewClientHandler()))
	if err != nil {
		log.Error("Error", zap.Error(err))
		return err
	}

	u := partyproto.NewUserServiceClient(userconn)
	b := bkgproto.NewBkgServiceClient(bkgconn)
	bss := bkgproto.NewBkgShipmentSummaryServiceClient(bkgconn)
	bsm := bkgproto.NewBkgSummaryServiceClient(bkgconn)

	bc := NewBkgController(log, u, b, h, workflowClient)
	bssc := NewBkgShipmentSummaryController(log, u, bss)
	bsmc := NewBkgSummaryController(log, u, bsm)

	hrlBkg := common.GetHTTPRateLimiter(store, rateOpt.BkgMaxRate, rateOpt.BkgMaxBurst)
	hrlBkgShipmentSummary := common.GetHTTPRateLimiter(store, rateOpt.BkgMaxRate, rateOpt.BkgMaxBurst)
	hrlBkgSummary := common.GetHTTPRateLimiter(store, rateOpt.BkgMaxRate, rateOpt.BkgMaxBurst)

	mux.Handle("/v2/bookings", common.AddMiddleware(hrlBkg.RateLimit(bc),
		common.EnsureValidToken(serverOpt.Auth0Audience, serverOpt.Auth0Domain), common.ValidatePermissions([]string{"bkg:cud", "bkg:read"}, serverOpt.Auth0Audience, serverOpt.Auth0Domain)))

	mux.Handle("/v2/bookings/", common.AddMiddleware(hrlBkg.RateLimit(bc),
		common.EnsureValidToken(serverOpt.Auth0Audience, serverOpt.Auth0Domain), common.ValidatePermissions([]string{"bkg:cud", "bkg:read"}, serverOpt.Auth0Audience, serverOpt.Auth0Domain)))

	mux.Handle("/v2/bkg-shipment-summaries", common.AddMiddleware(hrlBkgShipmentSummary.RateLimit(bssc),
		common.EnsureValidToken(serverOpt.Auth0Audience, serverOpt.Auth0Domain), common.ValidatePermissions([]string{"bkgss:cud", "bkgss:read"}, serverOpt.Auth0Audience, serverOpt.Auth0Domain)))

	mux.Handle("/v2/bkg-shipment-summaries/", common.AddMiddleware(hrlBkgShipmentSummary.RateLimit(bssc),
		common.EnsureValidToken(serverOpt.Auth0Audience, serverOpt.Auth0Domain), common.ValidatePermissions([]string{"bkgss:cud", "bkgss:read"}, serverOpt.Auth0Audience, serverOpt.Auth0Domain)))

	mux.Handle("/v2/booking-summaries", common.AddMiddleware(hrlBkgSummary.RateLimit(bsmc),
		common.EnsureValidToken(serverOpt.Auth0Audience, serverOpt.Auth0Domain), common.ValidatePermissions([]string{"bkgs:cud", "bkgs:read"}, serverOpt.Auth0Audience, serverOpt.Auth0Domain)))

	mux.Handle("/v2/booking-summaries/", common.AddMiddleware(hrlBkgSummary.RateLimit(bsmc),
		common.EnsureValidToken(serverOpt.Auth0Audience, serverOpt.Auth0Domain), common.ValidatePermissions([]string{"bkgs:cud", "bkgs:read"}, serverOpt.Auth0Audience, serverOpt.Auth0Domain)))

	return nil
}
