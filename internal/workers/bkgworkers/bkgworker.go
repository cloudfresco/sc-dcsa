package bkgworkers

import (
	"os"

	"github.com/cloudfresco/sc-dcsa/internal/config"
	"go.uber.org/cadence/worker"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	bkgproto "github.com/cloudfresco/sc-dcsa/internal/protogen/bkg/v2"

	bkgworkflows "github.com/cloudfresco/sc-dcsa/internal/workflows/bkgworkflows"

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
	h.StartWorkers(h.Config.DomainName, bkgworkflows.ApplicationName, workerOptions)
}

func StartBkgWorker(log *zap.Logger, isTest bool, pwd string, grpcServerOpt *config.GrpcServerOptions, configFilePath string) {
	var h common.WfHelper
	h.SetupServiceConfig(configFilePath)

	creds, err := common.GetClientCred(log, isTest, pwd, grpcServerOpt)
	if err != nil {
		log.Error("Error", zap.Error(err))
		os.Exit(1)
	}

	bkgconn, err := grpc.NewClient(grpcServerOpt.GrpcBkgServerPort, grpc.WithTransportCredentials(creds), grpc.WithStatsHandler(otelgrpc.NewClientHandler()))
	if err != nil {
		log.Error("Error",
			zap.Error(err))
		os.Exit(1)
	}
	bkgServiceClient := bkgproto.NewBkgServiceClient(bkgconn)
	bkgActivities := &bkgworkflows.BkgActivities{BkgServiceClient: bkgServiceClient}

	h.RegisterWorkflow(bkgworkflows.CreateBookingWorkflow)
	h.RegisterWorkflow(bkgworkflows.UpdateBookingWorkflow)
	h.RegisterWorkflow(bkgworkflows.CancelBookingByCarrierBookingReferenceWorkflow)
	h.RegisterActivity(bkgActivities)

	startWorkers(&h)

	// The workers are supposed to be long running process that should not exit.
	// Use select{} to block indefinitely for samples, you can quit by CMD+C.
	select {}
}
