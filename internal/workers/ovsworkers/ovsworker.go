package ovsworkers

import (
	"os"

	"github.com/cloudfresco/sc-dcsa/internal/config"
	"go.uber.org/cadence/worker"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	ovsproto "github.com/cloudfresco/sc-dcsa/internal/protogen/ovs/v3"

	ovsworkflows "github.com/cloudfresco/sc-dcsa/internal/workflows/ovsworkflows"

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
	h.StartWorkers(h.Config.DomainName, ovsworkflows.ApplicationName, workerOptions)
}

func StartOvsWorker(log *zap.Logger, isTest bool, pwd string, grpcServerOpt *config.GrpcServerOptions, configFilePath string) {
	var h common.WfHelper
	h.SetupServiceConfig(configFilePath)

	creds, err := common.GetClientCred(log, isTest, pwd, grpcServerOpt)
	if err != nil {
		log.Error("Error", zap.Error(err))
		os.Exit(1)
	}

	ovsconn, err := grpc.NewClient(grpcServerOpt.GrpcOvsServerPort, grpc.WithTransportCredentials(creds), grpc.WithStatsHandler(otelgrpc.NewClientHandler()))
	if err != nil {
		log.Error("Error",
			zap.Error(err))
		os.Exit(1)
	}
	serviceScheduleServiceClient := ovsproto.NewServiceScheduleServiceClient(ovsconn)
	serviceScheduleActivities := &ovsworkflows.ServiceScheduleActivities{ServiceScheduleServiceClient: serviceScheduleServiceClient}

	h.RegisterWorkflow(ovsworkflows.CreateServiceScheduleWorkflow)
	h.RegisterWorkflow(ovsworkflows.UpdateServiceScheduleWorkflow)
	h.RegisterActivity(serviceScheduleActivities)

	startWorkers(&h)

	// The workers are supposed to be long running process that should not exit.
	// Use select{} to block indefinitely for samples, you can quit by CMD+C.
	select {}
}
