package api

import (
	"bytes"
	"core/common"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

func Trace() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bodyLogWriter := &BodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}
		ctx.Writer = bodyLogWriter

		//心跳检测不记录
		if ctx.Request.URL.Path == common.HeartbeatPath {
			ctx.Next()
			return
		}

		spanCtx, span := otel.Tracer("middlewares").Start(ctx, ctx.Request.URL.Path)
		ctx.Set("traceId", span.SpanContext().TraceID().String())
		ctx.Set("parentCtx", spanCtx)
		ctx.Next()

		responseBody := bodyLogWriter.body.String()
		span.SetAttributes(attribute.String("request", ctx.Request.PostForm.Encode()))
		span.SetAttributes(attribute.String("response", responseBody))

		span.AddEvent("")

		span.End()
	}
}
