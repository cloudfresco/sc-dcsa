package bkgcontrollers

import (
	"context"
	"net/http"
	"os"
	"path/filepath"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	"github.com/cloudfresco/sc-dcsa/internal/config"
	bkgproto "github.com/cloudfresco/sc-dcsa/internal/protogen/bkg/v2"
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

// Init the bkg controllers
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

// InitTest the bkg controllers
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

	bkgconn, err := grpc.NewClient(grpcServerOpt.GrpcBkgServerPort, grpc.WithTransportCredentials(creds), grpc.WithStatsHandler(otelgrpc.NewClientHandler()))
	if err != nil {
		log.Error("Error", zap.Error(err))
		return err
	}

	b := bkgproto.NewBkgServiceClient(bkgconn)
	bss := bkgproto.NewBkgShipmentSummaryServiceClient(bkgconn)
	bs := bkgproto.NewBkgSummaryServiceClient(bkgconn)

	initBkg(mux, serverOpt, log, u, b, h, workflowClient)
	initBkgShipmentSummary(mux, serverOpt, log, u, bss, h, workflowClient)
	initBkgSummary(mux, serverOpt, log, u, bs, h, workflowClient)

	return nil
}

func initBkg(mux *http.ServeMux, serverOpt *config.ServerOptions, log *zap.Logger, u partyproto.UserServiceClient, b bkgproto.BkgServiceClient, wfHelper common.WfHelper, workflowClient client.Client) {
	bc := NewBkgController(log, u, b, h, workflowClient, serverOpt)

	mux.Handle("GET /v2/bookings/{carrierBookingRequestReference}", http.HandlerFunc(bc.GetBookingByCarrierBookingRequestReference))

	mux.Handle("POST /v2/bookings", http.HandlerFunc(bc.CreateBooking))

	mux.Handle("PUT /v2/bookings/{carrierBookingRequestReference}", http.HandlerFunc(bc.UpdateBooking))

	mux.Handle("PATCH /v2/bookings/{carrierBookingRequestReference}", http.HandlerFunc(bc.CancelBookingByCarrierBookingRequestReference))
}

func initBkgShipmentSummary(mux *http.ServeMux, serverOpt *config.ServerOptions, log *zap.Logger, u partyproto.UserServiceClient, bss bkgproto.BkgShipmentSummaryServiceClient, wfHelper common.WfHelper, workflowClient client.Client) {
	bssc := NewBkgShipmentSummaryController(log, u, bss, serverOpt)

	mux.Handle("GET /v2/shipment-summaries", http.HandlerFunc(bssc.GetBkgShipmentSummaries))
}

func initBkgSummary(mux *http.ServeMux, serverOpt *config.ServerOptions, log *zap.Logger, u partyproto.UserServiceClient, bs bkgproto.BkgSummaryServiceClient, wfHelper common.WfHelper, workflowClient client.Client) {
	bsc := NewBkgSummaryController(log, u, bs, serverOpt)

	mux.Handle("GET /v2/booking-summaries", http.HandlerFunc(bsc.GetBookingSummaries))
}
