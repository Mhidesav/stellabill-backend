package postgres

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

var planTracer = otel.Tracer("repository/plans")

// StartPlanSpan creates an OpenTelemetry span for plan repository operations.
func StartPlanSpan(ctx context.Context, name string, planID string) (context.Context, trace.Span) {
	opts := []trace.SpanStartOption{}
	if planID != "" {
		opts = append(opts, trace.WithAttributes(attribute.String("plan.id", planID)))
	}
	return planTracer.Start(ctx, name, opts...)
}
