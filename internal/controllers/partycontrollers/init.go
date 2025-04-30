package partycontrollers

import (
	"context"
	"net/http"
	"os"
	"path/filepath"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	"github.com/cloudfresco/sc-dcsa/internal/config"

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

// Init the party controllers
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

// InitTest the party controllers
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

	partyconn, err := grpc.NewClient(grpcServerOpt.GrpcPartyServerPort, grpc.WithTransportCredentials(creds), grpc.WithStatsHandler(otelgrpc.NewClientHandler()))
	if err != nil {
		log.Error("Error", zap.Int("msgnum", 110), zap.Error(err))
		return err
	}

	p := partyproto.NewPartyServiceClient(partyconn)

	initUsers(mux, serverOpt, log, u, h, workflowClient)
	initParties(mux, serverOpt, log, u, p, h, workflowClient)

	return nil
}

func initParties(mux *http.ServeMux, serverOpt *config.ServerOptions, log *zap.Logger, u partyproto.UserServiceClient, p partyproto.PartyServiceClient, wfHelper common.WfHelper, workflowClient client.Client) {
	pp := NewPartyController(log, p, u, serverOpt)

	mux.Handle("GET /v0.1/parties", http.HandlerFunc(pp.GetParties))
	mux.Handle("GET /v0.1/parties/{id}", http.HandlerFunc(pp.GetParty))
	mux.Handle("POST /v0.1/create", http.HandlerFunc(pp.CreateParty))
	mux.Handle("POST /v0.1/parties/{id}/partyContactDetailcreate", http.HandlerFunc(pp.CreatePartyContactDetail))
	mux.Handle("PUT /v0.1/parties/{id}", http.HandlerFunc(pp.UpdateParty))
	mux.Handle("DELETE /v0.1/parties/{id}", http.HandlerFunc(pp.DeleteParty))
}

func initUsers(mux *http.ServeMux, serverOpt *config.ServerOptions, log *zap.Logger, u partyproto.UserServiceClient, wfHelper common.WfHelper, workflowClient client.Client) {
	usc := NewUserController(log, u, h, workflowClient, serverOpt)

	mux.Handle("GET /v0.1/users", http.HandlerFunc(usc.GetUsers))
	mux.Handle("GET /v0.1/users/{id}", http.HandlerFunc(usc.GetUser))
	mux.Handle("POST /v0.1/users/getuserbyemail", http.HandlerFunc(usc.GetUserByEmail))
	mux.Handle("POST /v0.1/users/change-password", http.HandlerFunc(usc.ChangePassword))
	mux.Handle("PUT /v0.1/users/{id}", http.HandlerFunc(usc.UpdateUser))
	mux.Handle("DELETE /v0.1/users/{id}", http.HandlerFunc(usc.DeleteUser))
}
