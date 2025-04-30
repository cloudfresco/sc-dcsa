package eblcontrollers

import (
	"context"
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
func Init(log *zap.Logger, mux *http.ServeMux, store *goredisstore.GoRedisStore, serverOpt *config.ServerOptions, grpcServerOpt *config.GrpcServerOptions, uptraceOpt *config.UptraceOptions, configFilePath string) error {
	pwd, _ := os.Getwd()
	keyPath := pwd + filepath.FromSlash(grpcServerOpt.GrpcCaCertPath)

	err := initSetup(log, mux, keyPath, configFilePath, serverOpt, grpcServerOpt)
	if err != nil {
		log.Error("Error", zap.Int("msgnum", 110), zap.Error(err))
		return err
	}

	return nil
}

// InitTest the ebl controllers
func InitTest(log *zap.Logger, mux *http.ServeMux, store *goredisstore.GoRedisStore, serverOpt *config.ServerOptions, grpcServerOpt *config.GrpcServerOptions, uptraceOpt *config.UptraceOptions, configFilePath string) error {
	pwd, _ := os.Getwd()
	keyPath := filepath.Join(pwd, filepath.FromSlash("/../../../")+filepath.FromSlash(grpcServerOpt.GrpcCaCertPath))

	err := initSetup(log, mux, keyPath, configFilePath, serverOpt, grpcServerOpt)
	if err != nil {
		log.Error("Error", zap.Int("msgnum", 110), zap.Error(err))
		return err
	}

	return nil
}

func initSetup(log *zap.Logger, mux *http.ServeMux, keyPath string, configFilePath string, serverOpt *config.ServerOptions, grpcServerOpt *config.GrpcServerOptions) error {
	creds, err := credentials.NewClientTLSFromFile(keyPath, "localhost")
	if err != nil {
		log.Error("Error", zap.Int("msgnum", 110), zap.Error(err))
	}

	tp, err := config.InitTracerProvider()
	if err != nil {
		log.Error("Error", zap.Int("msgnum", 9108), zap.Error(err))
	}
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Error("Error", zap.Int("msgnum", 9108), zap.Error(err))
		}
	}()

	h.SetupServiceConfig(configFilePath)
	workflowClient, err = h.Builder.BuildCadenceClient()
	if err != nil {
		panic(err)
	}

	userconn, err := grpc.NewClient(grpcServerOpt.GrpcUserServerPort, grpc.WithTransportCredentials(creds), grpc.WithStatsHandler(otelgrpc.NewClientHandler()))
	if err != nil {
		log.Error("Error", zap.Int("msgnum", 113), zap.Error(err))
		return err
	}

	u := partyproto.NewUserServiceClient(userconn)

	eblconn, err := grpc.NewClient(grpcServerOpt.GrpcEblServerPort, grpc.WithTransportCredentials(creds), grpc.WithStatsHandler(otelgrpc.NewClientHandler()))
	if err != nil {
		log.Error("Error", zap.Error(err))
		return err
	}

	si := eblproto.NewShippingInstructionServiceClient(eblconn)
	sis := eblproto.NewShippingInstructionSummaryServiceClient(eblconn)
	td := eblproto.NewTransportDocumentServiceClient(eblconn)

	initShippingInstruction(mux, serverOpt, log, u, si, h, workflowClient)
	initShippingInstructionSummary(mux, serverOpt, log, u, sis, h, workflowClient)
	initTransportDocument(mux, serverOpt, log, u, td, h, workflowClient)
	initTransportDocumentSummary(mux, serverOpt, log, u, td, h, workflowClient)

	return nil
}

func initShippingInstruction(mux *http.ServeMux, serverOpt *config.ServerOptions, log *zap.Logger, u partyproto.UserServiceClient, si eblproto.ShippingInstructionServiceClient, wfHelper common.WfHelper, workflowClient client.Client) {
	sic := NewShippingInstructionController(log, u, si, h, workflowClient, serverOpt)

	mux.Handle("GET /v2/shipping-instructions/{shippingInstructionRequestReference}", http.HandlerFunc(sic.GetShippingInstructionByShippingInstructionReference))

	mux.Handle("POST /v2/shipping-instructions", http.HandlerFunc(sic.CreateShippingInstruction))

	mux.Handle("PUT /v2/shipping-instructions/{shippingInstructionRequestReference}", http.HandlerFunc(sic.UpdateShippingInstruction))
}

func initShippingInstructionSummary(mux *http.ServeMux, serverOpt *config.ServerOptions, log *zap.Logger, u partyproto.UserServiceClient, sis eblproto.ShippingInstructionSummaryServiceClient, wfHelper common.WfHelper, workflowClient client.Client) {
	sisc := NewShippingInstructionSummaryController(log, u, sis, serverOpt)

	mux.Handle("GET /v2/shipping-instructions-summaries}", http.HandlerFunc(sisc.GetShippingInstructionSummaries))
}

func initTransportDocument(mux *http.ServeMux, serverOpt *config.ServerOptions, log *zap.Logger, u partyproto.UserServiceClient, td eblproto.TransportDocumentServiceClient, wfHelper common.WfHelper, workflowClient client.Client) {
	tdc := NewTransportDocumentController(log, u, td, h, workflowClient, serverOpt)

	mux.Handle("GET /v2/transport-documents/{transportDocumentRequestReference}", http.HandlerFunc(tdc.GetTransportDocumentByTransportDocumentReference))

	mux.Handle("POST /v2/transport-documents", http.HandlerFunc(tdc.CreateTransportDocument))

	mux.Handle("PUT /v2/transport-documents/{transportDocumentRequestReference}", http.HandlerFunc(tdc.ApproveTransportDocument))
}

func initTransportDocumentSummary(mux *http.ServeMux, serverOpt *config.ServerOptions, log *zap.Logger, u partyproto.UserServiceClient, td eblproto.TransportDocumentServiceClient, wfHelper common.WfHelper, workflowClient client.Client) {
	tdsc := NewTransportDocumentSummaryController(log, u, td, serverOpt)

	mux.Handle("GET /v2/transport-document-summaries}", http.HandlerFunc(tdsc.GetTransportDocumentSummaries))
}
