package common

import (
	"errors"

	"github.com/opentracing/opentracing-go"
	"github.com/uber-go/tally"
	apiv1 "github.com/uber/cadence-idl/go/proto/api/v1"
	"go.uber.org/cadence/.gen/go/cadence/workflowserviceclient"
	"go.uber.org/cadence/client"
	"go.uber.org/cadence/compatibility"
	"go.uber.org/cadence/encoded"
	"go.uber.org/cadence/workflow"
	"go.uber.org/yarpc"
	"go.uber.org/yarpc/transport/grpc"
	"go.uber.org/zap"
)

const (
	_cadenceClientName      = "cadence-client"
	_cadenceFrontendService = "cadence-frontend"
)

// WorkflowClientBuilder build client to cadence service
type WorkflowClientBuilder struct {
	hostPort       string
	dispatcher     *yarpc.Dispatcher
	domain         string
	clientIdentity string
	metricsScope   tally.Scope
	Logger         *zap.Logger
	ctxProps       []workflow.ContextPropagator
	dataConverter  encoded.DataConverter
	tracer         opentracing.Tracer
}

// NewBuilder creates a new WorkflowClientBuilder
func NewBuilder(logger *zap.Logger) *WorkflowClientBuilder {
	return &WorkflowClientBuilder{
		Logger: logger,
	}
}

// SetHostPort sets the hostport for the builder
func (wcb *WorkflowClientBuilder) SetHostPort(hostport string) *WorkflowClientBuilder {
	wcb.hostPort = hostport
	return wcb
}

// SetDomain sets the domain for the builder
func (wcb *WorkflowClientBuilder) SetDomain(domain string) *WorkflowClientBuilder {
	wcb.domain = domain
	return wcb
}

// SetClientIdentity sets the identity for the builder
func (wcb *WorkflowClientBuilder) SetClientIdentity(identity string) *WorkflowClientBuilder {
	wcb.clientIdentity = identity
	return wcb
}

// SetMetricsScope sets the metrics scope for the builder
func (wcb *WorkflowClientBuilder) SetMetricsScope(metricsScope tally.Scope) *WorkflowClientBuilder {
	wcb.metricsScope = metricsScope
	return wcb
}

// SetDispatcher sets the dispatcher for the builder
func (wcb *WorkflowClientBuilder) SetDispatcher(dispatcher *yarpc.Dispatcher) *WorkflowClientBuilder {
	wcb.dispatcher = dispatcher
	return wcb
}

// SetContextPropagators sets the context propagators for the builder
func (wcb *WorkflowClientBuilder) SetContextPropagators(ctxProps []workflow.ContextPropagator) *WorkflowClientBuilder {
	wcb.ctxProps = ctxProps
	return wcb
}

// SetDataConverter sets the data converter for the builder
func (wcb *WorkflowClientBuilder) SetDataConverter(dataConverter encoded.DataConverter) *WorkflowClientBuilder {
	wcb.dataConverter = dataConverter
	return wcb
}

// SetTracer sets the tracer for the builder
func (wcb *WorkflowClientBuilder) SetTracer(tracer opentracing.Tracer) *WorkflowClientBuilder {
	wcb.tracer = tracer
	return wcb
}

// BuildCadenceClient builds a client to cadence service
func (wcb *WorkflowClientBuilder) BuildCadenceClient() (client.Client, error) {
	service, err := wcb.BuildServiceClient()
	if err != nil {
		return nil, err
	}

	return client.NewClient(
		service,
		wcb.domain,
		&client.Options{
			Identity:           wcb.clientIdentity,
			MetricsScope:       wcb.metricsScope,
			DataConverter:      wcb.dataConverter,
			ContextPropagators: wcb.ctxProps,
			Tracer:             wcb.tracer,
			FeatureFlags: client.FeatureFlags{
				WorkflowExecutionAlreadyCompletedErrorEnabled: true,
			},
		}), nil
}

// BuildCadenceDomainClient builds a domain client to cadence service
func (wcb *WorkflowClientBuilder) BuildCadenceDomainClient() (client.DomainClient, error) {
	service, err := wcb.BuildServiceClient()
	if err != nil {
		return nil, err
	}

	return client.NewDomainClient(
		service,
		&client.Options{
			Identity:           wcb.clientIdentity,
			MetricsScope:       wcb.metricsScope,
			ContextPropagators: wcb.ctxProps,
			FeatureFlags: client.FeatureFlags{
				WorkflowExecutionAlreadyCompletedErrorEnabled: true,
			},
		},
	), nil
}

// BuildServiceClient builds a rpc service client to cadence service
func (wcb *WorkflowClientBuilder) BuildServiceClient() (workflowserviceclient.Interface, error) {
	if err := wcb.build(); err != nil {
		return nil, err
	}

	if wcb.dispatcher == nil {
		wcb.Logger.Fatal("No RPC dispatcher provided to create a connection to Cadence Service")
	}

	clientConfig := wcb.dispatcher.ClientConfig(_cadenceFrontendService)
	return compatibility.NewThrift2ProtoAdapter(
		apiv1.NewDomainAPIYARPCClient(clientConfig),
		apiv1.NewWorkflowAPIYARPCClient(clientConfig),
		apiv1.NewWorkerAPIYARPCClient(clientConfig),
		apiv1.NewVisibilityAPIYARPCClient(clientConfig),
	), nil
}

func (wcb *WorkflowClientBuilder) build() error {
	if wcb.dispatcher != nil {
		return nil
	}

	if len(wcb.hostPort) == 0 {
		return errors.New("HostPort is empty")
	}

	wcb.Logger.Debug("Creating RPC dispatcher outbound",
		zap.String("ServiceName", _cadenceFrontendService),
		zap.String("HostPort", wcb.hostPort))

	wcb.dispatcher = yarpc.NewDispatcher(yarpc.Config{
		Name: _cadenceClientName,
		Outbounds: yarpc.Outbounds{
			_cadenceFrontendService: {Unary: grpc.NewTransport().NewSingleOutbound(wcb.hostPort)},
		},
	})

	if wcb.dispatcher != nil {
		if err := wcb.dispatcher.Start(); err != nil {
			wcb.Logger.Fatal("Failed to create outbound transport channel: %v", zap.Error(err))
		}
	}

	return nil
}
