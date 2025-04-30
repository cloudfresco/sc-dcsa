package tntworkers

import (
	"os"

	"github.com/cloudfresco/sc-dcsa/internal/config"
	"go.uber.org/cadence/worker"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	tntproto "github.com/cloudfresco/sc-dcsa/internal/protogen/tnt/v3"

	tntworkflows "github.com/cloudfresco/sc-dcsa/internal/workflows/tntworkflows"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
)

// This needs to be done as part of a bootstrap step when the process starts.
// The workers are supposed to be long running.
func startWorkers(h *common.WfHelper) {
	// Configure worker options.
	workerOptions := worker.Options{
		MetricsScope: h.WorkerMetricScope,
		Logger:       h.Logger,
	}
	h.StartWorkers(h.Config.DomainName, tntworkflows.ApplicationName, workerOptions)
}

func StartTntWorker(log *zap.Logger, isTest bool, pwd string, grpcServerOpt *config.GrpcServerOptions, configFilePath string) {
	var h common.WfHelper
	h.SetupServiceConfig(configFilePath)

	creds, err := common.GetClientCred(log, isTest, pwd, grpcServerOpt)
	if err != nil {
		log.Error("Error", zap.Error(err))
		os.Exit(1)
	}

	tntconn, err := grpc.NewClient(grpcServerOpt.GrpcTntServerPort, grpc.WithTransportCredentials(creds), grpc.WithStatsHandler(otelgrpc.NewClientHandler()))
	if err != nil {
		log.Error("Error",
			zap.Error(err))
		os.Exit(1)
	}

	eventSubscriptionServiceClient := tntproto.NewEventSubscriptionServiceClient(tntconn)
	eventSubscriptionActivities := &tntworkflows.EventSubscriptionActivities{EventSubscriptionServiceClient: eventSubscriptionServiceClient}

	eventServiceClient := tntproto.NewEventServiceClient(tntconn)
	eventActivities := &tntworkflows.EventActivities{EventServiceClient: eventServiceClient}

	h.RegisterWorkflow(tntworkflows.CreateEventSubscriptionWorkflow)
	h.RegisterWorkflow(tntworkflows.CreateEquipmentEventWorkflow)
	h.RegisterWorkflow(tntworkflows.CreateOperationsEventWorkflow)
	h.RegisterWorkflow(tntworkflows.CreateShipmentEventWorkflow)
	h.RegisterWorkflow(tntworkflows.CreateTransportEventWorkflow)
	h.RegisterActivity(eventSubscriptionActivities)
	h.RegisterActivity(eventActivities)

	startWorkers(&h)

	// The workers are supposed to be long running process that should not exit.
	// Use select{} to block indefinitely for samples, you can quit by CMD+C.
	select {}
}
