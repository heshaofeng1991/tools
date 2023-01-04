package jaeger

import (
	"core/config"
	"core/utils"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
)

func New(conf *config.Jaeger, service *config.Service) (*tracesdk.TracerProvider, error) {
	var endpointOption jaeger.EndpointOption
	if conf.Port == "" {
		endpointOption = jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(conf.Host))
	} else {
		endpointOption = jaeger.WithAgentEndpoint(jaeger.WithAgentHost(conf.Host), jaeger.WithAgentPort(conf.Port))
	}
	exp, err := jaeger.New(endpointOption)
	if err != nil {
		return nil, err
	}

	tp := tracesdk.NewTracerProvider(
		tracesdk.WithBatcher(exp),
		tracesdk.WithSampler(tracesdk.AlwaysSample()),
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			attribute.String("hostname", utils.GetOutboundIP()),
			semconv.ServiceNameKey.String(service.Name),
		)),
	)

	otel.SetTracerProvider(tp)
	return tp, nil
}
