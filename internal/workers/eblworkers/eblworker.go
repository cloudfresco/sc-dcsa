package eblworkers

import (
	"os"

	"github.com/cloudfresco/sc-dcsa/internal/config"
	"go.uber.org/cadence/worker"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	eblproto "github.com/cloudfresco/sc-dcsa/internal/protogen/ebl/v1"

	eblworkflows "github.com/cloudfresco/sc-dcsa/internal/workflows/eblworkflows"

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
	h.StartWorkers(h.Config.DomainName, eblworkflows.ApplicationName, workerOptions)
}

func StartEblWorker(log *zap.Logger, isTest bool, pwd string, grpcServerOpt *config.GrpcServerOptions, configFilePath string) {
	var h common.WfHelper
	h.SetupServiceConfig(configFilePath)

	creds, err := common.GetClientCred(log, isTest, pwd, grpcServerOpt)
	if err != nil {
		log.Error("Error", zap.Error(err))
		os.Exit(1)
	}

	eblconn, err := grpc.NewClient(grpcServerOpt.GrpcEblServerPort, grpc.WithTransportCredentials(creds), grpc.WithStatsHandler(otelgrpc.NewClientHandler()))
	if err != nil {
		log.Error("Error",
			zap.Error(err))
		os.Exit(1)
	}
	shippingInstructionServiceClient := eblproto.NewShippingInstructionServiceClient(eblconn)
	shippingInstructionActivities := &eblworkflows.ShippingInstructionActivities{ShippingInstructionServiceClient: shippingInstructionServiceClient}
	transportDocumentServiceClient := eblproto.NewTransportDocumentServiceClient(eblconn)
	transportDocumentActivities := &eblworkflows.TransportDocumentActivities{TransportDocumentServiceClient: transportDocumentServiceClient}
	surrenderRequestServiceClient := eblproto.NewSurrenderRequestServiceClient(eblconn)
	surrenderRequestAnswerServiceClient := eblproto.NewSurrenderRequestAnswerServiceClient(eblconn)
	surrenderRequestActivities := &eblworkflows.SurrenderRequestActivities{SurrenderRequestServiceClient: surrenderRequestServiceClient}
	surrenderRequestAnswerActivities := &eblworkflows.SurrenderRequestAnswerActivities{SurrenderRequestAnswerServiceClient: surrenderRequestAnswerServiceClient}
	issueRequestServiceClient := eblproto.NewIssueRequestServiceClient(eblconn)
	issueRequestResponseServiceClient := eblproto.NewIssueRequestResponseServiceClient(eblconn)
	issueRequestActivities := &eblworkflows.IssueRequestActivities{IssueRequestServiceClient: issueRequestServiceClient}
	issueRequestResponseActivities := &eblworkflows.IssueRequestResponseActivities{IssueRequestResponseServiceClient: issueRequestResponseServiceClient}

	h.RegisterWorkflow(eblworkflows.CreateShippingInstructionWorkflow)
	h.RegisterWorkflow(eblworkflows.UpdateShippingInstructionWorkflow)
	h.RegisterWorkflow(eblworkflows.CreateTransportDocumentWorkflow)
	h.RegisterWorkflow(eblworkflows.ApproveTransportDocumentWorkflow)
	h.RegisterWorkflow(eblworkflows.CreateTransactionPartyWorkflow)
	h.RegisterWorkflow(eblworkflows.UpdateTransactionPartyWorkflow)
	h.RegisterWorkflow(eblworkflows.CreateTransactionPartySupportingCodeWorkflow)
	h.RegisterWorkflow(eblworkflows.UpdateTransactionPartySupportingCodeWorkflow)
	h.RegisterWorkflow(eblworkflows.CreateSurrenderRequestWorkflow)
	h.RegisterWorkflow(eblworkflows.UpdateSurrenderRequestWorkflow)
	h.RegisterWorkflow(eblworkflows.CreateSurrenderRequestAnswerWorkflow)
	h.RegisterWorkflow(eblworkflows.UpdateSurrenderRequestAnswerWorkflow)
	h.RegisterWorkflow(eblworkflows.CreateEndorsementChainLinkWorkflow)
	h.RegisterWorkflow(eblworkflows.UpdateEndorsementChainLinkWorkflow)
	h.RegisterWorkflow(eblworkflows.CreateIssuePartyWorkflow)
	h.RegisterWorkflow(eblworkflows.UpdateIssuePartyWorkflow)
	h.RegisterWorkflow(eblworkflows.CreateIssuePartySupportingCodeWorkflow)
	h.RegisterWorkflow(eblworkflows.UpdateIssuePartySupportingCodeWorkflow)
	h.RegisterWorkflow(eblworkflows.CreateIssuanceRequestWorkflow)
	h.RegisterWorkflow(eblworkflows.UpdateIssuanceRequestWorkflow)
	h.RegisterWorkflow(eblworkflows.CreateIssuanceRequestResponseWorkflow)
	h.RegisterWorkflow(eblworkflows.UpdateIssuanceRequestResponseWorkflow)
	h.RegisterWorkflow(eblworkflows.CreateEblVisualizationWorkflow)
	h.RegisterWorkflow(eblworkflows.UpdateEblVisualizationWorkflow)
	h.RegisterActivity(shippingInstructionActivities)
	h.RegisterActivity(transportDocumentActivities)
	h.RegisterActivity(surrenderRequestActivities)
	h.RegisterActivity(surrenderRequestAnswerActivities)
	h.RegisterActivity(issueRequestActivities)
	h.RegisterActivity(issueRequestResponseActivities)

	startWorkers(&h)

	// The workers are supposed to be long running process that should not exit.
	// Use select{} to block indefinitely for samples, you can quit by CMD+C.
	select {}
}
