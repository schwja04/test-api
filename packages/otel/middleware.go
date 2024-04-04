package otel

import (
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func DatadogFriendlyOtelMapping() gin.HandlerFunc {
	return func(c *gin.Context) {
		span := trace.SpanFromContext(c.Request.Context())
		span.SetAttributes(attribute.String("operation.name", "gin.request"))

		c.Next()
	}
}
