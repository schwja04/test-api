package otel

import (
	"context"
	"fmt"
	"regexp"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	Tracer trace.Tracer
)

func InitTraceProvider(ctx context.Context, exportConsole bool) (func(context.Context) error, error) {

	var processor sdktrace.SpanProcessor
	var returnConn *grpc.ClientConn = nil

	if exportConsole {
		exporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
		if err != nil {
			return providerShutdown(nil, nil), err
		}
		processor = sdktrace.NewBatchSpanProcessor(exporter)
	} else {
		endpoint := "jaeger:4317" // os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")
		if endpoint == "" {
			endpoint = "localhost:4317"
		} else {
			// Dialer does not appreciate targets with http:// prefixes
			targetCleaner, err := regexp.Compile(`^https?://`)
			if err != nil {
				return providerShutdown(nil, nil), err
			}
			endpoint = targetCleaner.ReplaceAllString(endpoint, "")
		}

		conn, err := grpc.DialContext(ctx, endpoint,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.FailOnNonTempDialError(true),
			grpc.WithBlock(),
		)
		if err != nil {
			return providerShutdown(nil, nil), err
		}

		returnConn = conn

		exporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(conn))
		if err != nil {
			return providerShutdown(returnConn, nil), err
		}

		processor = sdktrace.NewBatchSpanProcessor(exporter)
	}

	res, err := resource.New(ctx, resource.WithFromEnv())
	if err != nil {
		return providerShutdown(returnConn, nil), err
	}

	provider := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(res),
		sdktrace.WithSpanProcessor(processor),
	)

	otel.SetTracerProvider(provider)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	Tracer = provider.Tracer("todo-app")

	return providerShutdown(returnConn, provider), nil
}

func providerShutdown(conn *grpc.ClientConn, provider *sdktrace.TracerProvider) func(context.Context) error {
	return func(ctx context.Context) error {
		var result error = nil
		if conn != nil {
			if err := conn.Close(); err != nil {
				result = fmt.Errorf("failed to close grpc connection: %w", err)
			}
		}

		if provider != nil {
			if err := provider.Shutdown(ctx); err != nil {
				result = fmt.Errorf("failed calling shutdown on provider: %w; conn close: %w", err, result)
			}
		}

		return result
	}
}
