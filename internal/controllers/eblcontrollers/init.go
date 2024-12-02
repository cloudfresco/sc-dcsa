package eblcontrollers

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	"github.com/cloudfresco/sc-dcsa/internal/config"

	eblproto "github.com/cloudfresco/sc-dcsa/internal/protogen/ebl/v1"
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

// Init the ebl controllers
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

	eblconn, err := grpc.NewClient(grpcServerOpt.GrpcEblServerPort, grpc.WithTransportCredentials(creds), grpc.WithStatsHandler(otelgrpc.NewClientHandler()))
	if err != nil {
		log.Error("Error", zap.Error(err))
		return err
	}

	u := partyproto.NewUserServiceClient(userconn)
	si := eblproto.NewShippingInstructionServiceClient(eblconn)
	sis := eblproto.NewShippingInstructionSummaryServiceClient(eblconn)
	td := eblproto.NewTransportDocumentServiceClient(eblconn)

	sic := NewShippingInstructionController(log, u, si, h, workflowClient)
	sisc := NewShippingInstructionSummaryController(log, u, sis)
	tdc := NewTransportDocumentController(log, u, td, h, workflowClient)

	hrlShippingInstruction := common.GetHTTPRateLimiter(store, rateOpt.EblMaxRate, rateOpt.EblMaxBurst)
	hrlShippingInstructionSummary := common.GetHTTPRateLimiter(store, rateOpt.EblMaxRate, rateOpt.EblMaxBurst)
	hrlTransportDocument := common.GetHTTPRateLimiter(store, rateOpt.EblMaxRate, rateOpt.EblMaxBurst)

	mux.Handle("/v0.1/shipping-instructions", common.AddMiddleware(hrlShippingInstruction.RateLimit(sic),
		common.EnsureValidToken(serverOpt.Auth0Audience, serverOpt.Auth0Domain), common.ValidatePermissions([]string{"shippinginstr:cud", "shippinginstr:read"}, serverOpt.Auth0Audience, serverOpt.Auth0Domain)))

	mux.Handle("/v0.1/shipping-instructions/", common.AddMiddleware(hrlShippingInstruction.RateLimit(sic),
		common.EnsureValidToken(serverOpt.Auth0Audience, serverOpt.Auth0Domain), common.ValidatePermissions([]string{"shippinginstr:cud", "shippinginstr:read"}, serverOpt.Auth0Audience, serverOpt.Auth0Domain)))

	mux.Handle("/v0.1/shipping-instructions-summaries", common.AddMiddleware(hrlShippingInstructionSummary.RateLimit(sisc),
		common.EnsureValidToken(serverOpt.Auth0Audience, serverOpt.Auth0Domain), common.ValidatePermissions([]string{"shippinginstr:cud", "shippinginstr:read"}, serverOpt.Auth0Audience, serverOpt.Auth0Domain)))

	mux.Handle("/v0.1/shipping-instructions-summaries/", common.AddMiddleware(hrlShippingInstructionSummary.RateLimit(sisc),
		common.EnsureValidToken(serverOpt.Auth0Audience, serverOpt.Auth0Domain), common.ValidatePermissions([]string{"shippinginstr:cud", "shippinginstr:read"}, serverOpt.Auth0Audience, serverOpt.Auth0Domain)))

	mux.Handle("/v0.1/transport_documents", common.AddMiddleware(hrlTransportDocument.RateLimit(tdc),
		common.EnsureValidToken(serverOpt.Auth0Audience, serverOpt.Auth0Domain), common.ValidatePermissions([]string{"transportdoc:cud", "transportdoc:read"}, serverOpt.Auth0Audience, serverOpt.Auth0Domain)))

	mux.Handle("/v0.1/transport_documents/", common.AddMiddleware(hrlTransportDocument.RateLimit(tdc),
		common.EnsureValidToken(serverOpt.Auth0Audience, serverOpt.Auth0Domain), common.ValidatePermissions([]string{"transportdoc:cud", "transportdoc:read"}, serverOpt.Auth0Audience, serverOpt.Auth0Domain)))

	mux.Handle("/v0.1/transport-document-summaries", common.AddMiddleware(hrlTransportDocument.RateLimit(tdc),
		common.EnsureValidToken(serverOpt.Auth0Audience, serverOpt.Auth0Domain), common.ValidatePermissions([]string{"transportdocsum:cud", "transportdocsum:read"}, serverOpt.Auth0Audience, serverOpt.Auth0Domain)))

	mux.Handle("/v0.1/transport-document-summaries/", common.AddMiddleware(hrlTransportDocument.RateLimit(tdc),
		common.EnsureValidToken(serverOpt.Auth0Audience, serverOpt.Auth0Domain), common.ValidatePermissions([]string{"transportdocsum:cud", "transportdocsum:read"}, serverOpt.Auth0Audience, serverOpt.Auth0Domain)))

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

	eblconn, err := grpc.NewClient(grpcServerOpt.GrpcEblServerPort, grpc.WithTransportCredentials(creds), grpc.WithStatsHandler(otelgrpc.NewClientHandler()))
	if err != nil {
		log.Error("Error", zap.Error(err))
		return err
	}

	u := partyproto.NewUserServiceClient(userconn)
	si := eblproto.NewShippingInstructionServiceClient(eblconn)
	sis := eblproto.NewShippingInstructionSummaryServiceClient(eblconn)
	td := eblproto.NewTransportDocumentServiceClient(eblconn)

	sic := NewShippingInstructionController(log, u, si, h, workflowClient)
	sisc := NewShippingInstructionSummaryController(log, u, sis)
	tdc := NewTransportDocumentController(log, u, td, h, workflowClient)

	hrlShippingInstruction := common.GetHTTPRateLimiter(store, rateOpt.EblMaxRate, rateOpt.EblMaxBurst)
	hrlShippingInstructionSummary := common.GetHTTPRateLimiter(store, rateOpt.EblMaxRate, rateOpt.EblMaxBurst)
	hrlTransportDocument := common.GetHTTPRateLimiter(store, rateOpt.EblMaxRate, rateOpt.EblMaxBurst)

	mux.Handle("/v0.1/shipping-instructions", common.AddMiddleware(hrlShippingInstruction.RateLimit(sic),
		common.EnsureValidToken(serverOpt.Auth0Audience, serverOpt.Auth0Domain), common.ValidatePermissions([]string{"shippinginstr:cud", "shippinginstr:read"}, serverOpt.Auth0Audience, serverOpt.Auth0Domain)))

	mux.Handle("/v0.1/shipping-instructions/", common.AddMiddleware(hrlShippingInstruction.RateLimit(sic),
		common.EnsureValidToken(serverOpt.Auth0Audience, serverOpt.Auth0Domain), common.ValidatePermissions([]string{"shippinginstr:cud", "shippinginstr:read"}, serverOpt.Auth0Audience, serverOpt.Auth0Domain)))

	mux.Handle("/v0.1/shipping-instructions-summaries", common.AddMiddleware(hrlShippingInstructionSummary.RateLimit(sisc),
		common.EnsureValidToken(serverOpt.Auth0Audience, serverOpt.Auth0Domain), common.ValidatePermissions([]string{"shippinginstr:cud", "shippinginstr:read"}, serverOpt.Auth0Audience, serverOpt.Auth0Domain)))

	mux.Handle("/v0.1/shipping-instructions-summaries/", common.AddMiddleware(hrlShippingInstructionSummary.RateLimit(sisc),
		common.EnsureValidToken(serverOpt.Auth0Audience, serverOpt.Auth0Domain), common.ValidatePermissions([]string{"shippinginstr:cud", "shippinginstr:read"}, serverOpt.Auth0Audience, serverOpt.Auth0Domain)))

	mux.Handle("/v0.1/transport_documents", common.AddMiddleware(hrlTransportDocument.RateLimit(tdc),
		common.EnsureValidToken(serverOpt.Auth0Audience, serverOpt.Auth0Domain), common.ValidatePermissions([]string{"transportdoc:cud", "transportdoc:read"}, serverOpt.Auth0Audience, serverOpt.Auth0Domain)))

	mux.Handle("/v0.1/transport_documents/", common.AddMiddleware(hrlTransportDocument.RateLimit(tdc),
		common.EnsureValidToken(serverOpt.Auth0Audience, serverOpt.Auth0Domain), common.ValidatePermissions([]string{"transportdoc:cud", "transportdoc:read"}, serverOpt.Auth0Audience, serverOpt.Auth0Domain)))

	mux.Handle("/v0.1/transport-document-summaries", common.AddMiddleware(hrlTransportDocument.RateLimit(tdc),
		common.EnsureValidToken(serverOpt.Auth0Audience, serverOpt.Auth0Domain), common.ValidatePermissions([]string{"transportdoc:cud", "transportdoc:read"}, serverOpt.Auth0Audience, serverOpt.Auth0Domain)))

	mux.Handle("/v0.1/transport-document-summaries/", common.AddMiddleware(hrlTransportDocument.RateLimit(tdc),
		common.EnsureValidToken(serverOpt.Auth0Audience, serverOpt.Auth0Domain), common.ValidatePermissions([]string{"transportdoc:cud", "transportdoc:read"}, serverOpt.Auth0Audience, serverOpt.Auth0Domain)))

	return nil
}
