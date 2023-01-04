package log

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

type Log struct {
	traceId   string
	storeName string
	field     []zap.Field
}

func New(ctx context.Context) *Log {
	l := &Log{}
	spanFromContext := trace.SpanFromContext(ctx)
	l.field = make([]zap.Field, 0)
	if spanFromContext.IsRecording() {
		l.traceId = spanFromContext.SpanContext().TraceID().String()
		l.field = append(l.field, zap.String("trace_id", l.traceId))
	}
	if storeName, ok := ctx.Value("storeName").(string); ok {
		l.storeName = storeName
		l.field = append(l.field, zap.String("store_name", l.storeName))
	}
	return l
}

func (l *Log) GetTraceId() string {
	return l.traceId
}

func (l *Log) SetStore(storeName string) *Log {
	l.storeName = storeName
	return l
}

func (l *Log) Field(field ...zap.Field) *Log {
	l.field = append(l.field, field...)
	return l
}

func (l *Log) Info(args ...interface{}) {
	zap.L().Info(fmt.Sprint(args), l.field...)
}

func (l *Log) Infof(template string, args ...interface{}) {
	zap.L().Info(fmt.Sprintf(template, args...), l.field...)
}

func (l *Log) Error(args ...interface{}) {
	zap.L().Error(fmt.Sprint(args), l.field...)
}

func (l *Log) Errorf(template string, args ...interface{}) {
	zap.L().Error(fmt.Sprintf(template, args...), l.field...)
}

func (l *Log) Warn(args ...interface{}) {
	zap.L().Warn(fmt.Sprint(args), l.field...)
}

func (l *Log) Warnf(template string, args ...interface{}) {
	zap.L().Warn(fmt.Sprintf(template, args...), l.field...)
}

func (l *Log) Fatal(args ...interface{}) {
	zap.L().Fatal(fmt.Sprint(args), l.field...)
}

func (l *Log) Fatalf(template string, args ...interface{}) {
	zap.L().Fatal(fmt.Sprintf(template, args...), l.field...)
}

func (l *Log) Panic(args ...interface{}) {
	zap.L().Panic(fmt.Sprint(args), l.field...)
}

func (l *Log) Panicf(template string, args ...interface{}) {
	zap.L().Panic(fmt.Sprintf(template, args...), l.field...)
}

func (l *Log) DPanic(args ...interface{}) {
	zap.L().DPanic(fmt.Sprint(args), l.field...)
}

func (l *Log) DPanicf(template string, args ...interface{}) {
	zap.L().DPanic(fmt.Sprintf(template, args...), l.field...)
}
