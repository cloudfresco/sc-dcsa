package tntcontrollers

import (
	"context"
	"net/http"
	"os"
	"path/filepath"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	"github.com/cloudfresco/sc-dcsa/internal/config"

	partyproto "github.com/cloudfresco/sc-dcsa/internal/protogen/party/v1"
	tntproto "github.com/cloudfresco/sc-dcsa/internal/protogen/tnt/v3"
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

// InitTest the user controllers
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

	tntconn, err := grpc.NewClient(grpcServerOpt.GrpcTntServerPort, grpc.WithTransportCredentials(creds), grpc.WithStatsHandler(otelgrpc.NewClientHandler()))
	if err != nil {
		log.Error("Error", zap.Error(err))
		return err
	}

	e := tntproto.NewEventServiceClient(tntconn)
	es := tntproto.NewEventSubscriptionServiceClient(tntconn)

	initEvent(mux, serverOpt, log, u, e, h, workflowClient)
	initEventSubscription(mux, serverOpt, log, u, es, h, workflowClient)

	return nil
}

func initEvent(mux *http.ServeMux, serverOpt *config.ServerOptions, log *zap.Logger, u partyproto.UserServiceClient, e tntproto.EventServiceClient, wfHelper common.WfHelper, workflowClient client.Client) {
	ec := NewEventController(log, u, e, h, workflowClient, serverOpt)

	mux.Handle("POST /v3/events/equipment-event", http.HandlerFunc(ec.CreateEquipmentEvent))
	mux.Handle("POST /v3/events/operations-event", http.HandlerFunc(ec.CreateOperationsEvent))
	mux.Handle("POST /v3/events/shipment-event", http.HandlerFunc(ec.CreateShipmentEvent))
	mux.Handle("POST /v3/events/transport-event", http.HandlerFunc(ec.CreateTransportEvent))
}

func initEventSubscription(mux *http.ServeMux, serverOpt *config.ServerOptions, log *zap.Logger, u partyproto.UserServiceClient, es tntproto.EventSubscriptionServiceClient, wfHelper common.WfHelper, workflowClient client.Client) {
	esc := NewEventSubscriptionController(log, u, es, h, workflowClient, serverOpt)

	mux.Handle("GET /v3/event-subscriptions", http.HandlerFunc(esc.GetEventSubscriptions))
	mux.Handle("GET /v3/event-subscriptions/{id}", http.HandlerFunc(esc.GetEventSubscriptionByCarrierEventSubscriptionRequestReference))
	mux.Handle("POST /v3/event-subscriptions", http.HandlerFunc(esc.CreateEventSubscription))
	mux.Handle("DELETE /v3/event-subscriptions/{id}", http.HandlerFunc(esc.DeleteEventSubscription))
}
