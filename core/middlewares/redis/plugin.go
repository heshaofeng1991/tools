package redis

import (
	"context"
	"core/common/log"
	"fmt"

	"github.com/go-redis/redis/v8"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type jaegerHook struct{}

const (
	optRedis  = "Redis-"
	CmdName   = "command"
	CmdArgs   = "args"
	CmdResult = "result"
)

func (j *jaegerHook) BeforeProcess(ctx context.Context, cmd redis.Cmder) (context.Context, error) {
	return trace.ContextWithSpan(otel.Tracer(cmd.Name()).Start(ctx, optRedis+cmd.Name())), nil
}

func (j *jaegerHook) AfterProcess(ctx context.Context, cmd redis.Cmder) error {
	span := trace.SpanFromContext(ctx)
	if span.IsRecording() {
		str := fmt.Sprintf("%+v", cmd)
		log.New(ctx).Info(str)
		span.SetAttributes(attribute.String(CmdResult, str))
		span.End()
	}
	return nil
}

func (j *jaegerHook) BeforeProcessPipeline(ctx context.Context, cmds []redis.Cmder) (context.Context, error) {
	return trace.ContextWithSpan(otel.Tracer("pipeline").Start(ctx, optRedis+"pipeline")), nil
}

func (j *jaegerHook) AfterProcessPipeline(ctx context.Context, cmds []redis.Cmder) error {
	span := trace.SpanFromContext(ctx)
	if span.IsRecording() {
		defer span.End()
		str := fmt.Sprintf("%+v", cmds)
		log.New(ctx).Info(str)
		span.SetAttributes(attribute.String(CmdResult, str))
		attribute.String(CmdResult, str)
	}
	return nil
}

func newJaegerHook() redis.Hook {
	return &jaegerHook{}
}
