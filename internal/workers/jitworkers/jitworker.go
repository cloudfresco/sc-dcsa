package jitworkers

import (
	"os"

	"github.com/cloudfresco/sc-dcsa/internal/config"
	"go.uber.org/cadence/worker"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	jitproto "github.com/cloudfresco/sc-dcsa/internal/protogen/jit/v1"

	jitworkflows "github.com/cloudfresco/sc-dcsa/internal/workflows/jitworkflows"

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
	h.StartWorkers(h.Config.DomainName, jitworkflows.ApplicationName, workerOptions)
}

func StartJitWorker(log *zap.Logger, isTest bool, pwd string, grpcServerOpt *config.GrpcServerOptions, configFilePath string) {
	var h common.WfHelper
	h.SetupServiceConfig(configFilePath)

	creds, err := common.GetClientCred(log, isTest, pwd, grpcServerOpt)
	if err != nil {
		log.Error("Error", zap.Error(err))
		os.Exit(1)
	}

	jitconn, err := grpc.NewClient(grpcServerOpt.GrpcJitServerPort, grpc.WithTransportCredentials(creds), grpc.WithStatsHandler(otelgrpc.NewClientHandler()))
	if err != nil {
		log.Error("Error",
			zap.Error(err))
		os.Exit(1)
	}
	timestampServiceClient := jitproto.NewTimestampServiceClient(jitconn)
	timestampActivities := &jitworkflows.TimestampActivities{TimestampServiceClient: timestampServiceClient}

	h.RegisterWorkflow(jitworkflows.CreateTimestampWorkflow)
	h.RegisterActivity(timestampActivities)

	startWorkers(&h)

	// The workers are supposed to be long running process that should not exit.
	// Use select{} to block indefinitely for samples, you can quit by CMD+C.
	select {}
}
