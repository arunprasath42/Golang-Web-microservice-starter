package middleware

import (
	"github.com/gin-gonic/gin"
)

// tracingHeaders - Headers used in tracing
var tracingHeaders = [9]string{
	"x-request-id",
	"x-b3-traceid",
	"x-b3-spanid",
	"x-b3-parentspanid",
	"x-b3-sampled",
	"x-b3-flags",
	"x-ot-span-context",
	"x-cloud-trace-context",
	"traceparent",
}

// TracingMiddleware - middleware
func TracingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		for _, header := range tracingHeaders {
			value := c.GetHeader(header)
			if value != "" {
				c.Writer.Header().Set(header, value)
			}
		}
		c.Next()
	}
}
