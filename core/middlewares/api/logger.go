package api

import (
	"bytes"
	"core/common"
	"fmt"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		bodyLogWriter := &BodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}
		ctx.Writer = bodyLogWriter
		ctx.Next()
		end := time.Now()
		requestBody, _ := ctx.GetRawData()
		raw, _ := url.QueryUnescape(ctx.Request.URL.RawQuery)
		srv := common.New(ctx)
		srv.Log.Field(
			zap.String("trace_id", ctx.GetString("traceId")),
			zap.String("store_mame", ctx.GetString("storeName")),
			zap.String("uri", ctx.Request.URL.Path),
			zap.String("raw", raw),
			zap.String("response_time", end.Sub(start).String()),
			zap.String("method", ctx.Request.Method),
			zap.String("status", fmt.Sprint(ctx.Writer.Status())),
			zap.String("remote_addr", ctx.ClientIP()),
			zap.String("content_type", ctx.Request.Header.Get("Content-Type")),
			zap.String("user_agent", ctx.Request.UserAgent()),
			zap.String("post_form", ctx.Request.PostForm.Encode()),
			zap.String("request_body", string(requestBody)),
			zap.String("response_body", bodyLogWriter.body.String()),
			zap.Error(ctx.Errors.Last())).Info("响应")
	}
}
