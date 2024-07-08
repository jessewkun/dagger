package sys

import (
	"context"
	"dagger/lib/logger"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

// TODO
// 未完成功能，请勿使用
// go resty support ot https://github.com/dubonzi/otelresty

func getTracer() (*sdktrace.TracerProvider, error) {
	exporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		return nil, err
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("dagger-service"),
		)),
	)

	otel.SetTracerProvider(tp)
	return tp, nil
}

// 初始化 OpenTelemetry
func InitTracer() {
	ctx := context.Background()
	tp, err := getTracer()
	if err != nil {
		logger.Fatal(ctx, "trace", "failed to initialize tracer: %v", err)
	}
	defer func() {
		if err := tp.Shutdown(ctx); err != nil {
			logger.Fatal(ctx, "trace", "failed to shutdown tracer: %v", err)
		}
	}()
}
