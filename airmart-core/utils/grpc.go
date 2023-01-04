package utils

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type grpcServerOptions struct {
}

// OpenTracingClientInterceptor grpc客户端拦截器
func OpenTracingClientInterceptor() grpc.DialOption {
	return grpc.WithUnaryInterceptor(func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		span := trace.SpanFromContext(ctx)
		if span.IsRecording() {
			md, ok := metadata.FromOutgoingContext(ctx)
			if ok {
				kvs := []string{
					"span-id",
					span.SpanContext().SpanID().String(),
					"trace-id",
					span.SpanContext().TraceID().String(),
				}
				ctx = metadata.AppendToOutgoingContext(ctx, kvs...)
			} else {
				md = metadata.New(map[string]string{
					"span-id":  span.SpanContext().SpanID().String(),
					"trace-id": span.SpanContext().TraceID().String(),
				})
				ctx = metadata.NewOutgoingContext(ctx, md)
			}
		}
		return invoker(ctx, method, req, reply, cc)
	})
}

// OpenTracingServerInterceptor grpc服务端拦截器
func OpenTracingServerInterceptor() grpc.ServerOption {
	return grpc.UnaryInterceptor(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		if info.FullMethod == "/grpc.health.v1.Health/Check" {
			return handler(ctx, req)
		}
		md, ok := metadata.FromIncomingContext(ctx)
		if ok {
			traceId, traceOk := md["trace-id"]
			spanId, spanOk := md["span-id"]
			if spanOk && traceOk {
				traceid, _ := trace.TraceIDFromHex(traceId[0])
				spanid, _ := trace.SpanIDFromHex(spanId[0])
				spanCtx := trace.NewSpanContext(trace.SpanContextConfig{
					TraceID:    traceid,
					SpanID:     spanid,
					TraceFlags: trace.FlagsSampled,
					TraceState: trace.TraceState{},
					Remote:     true,
				})
				ctx = trace.ContextWithRemoteSpanContext(ctx, spanCtx)
			}
		}
		spanCtx, span := otel.Tracer("OpenTracingServerInterceptor").Start(ctx, "grpc-"+info.FullMethod)
		res, err := handler(spanCtx, req)
		if err != nil {
			span.SetAttributes(attribute.String("err", err.Error()))
		}
		span.End()
		return res, err
	})
}
