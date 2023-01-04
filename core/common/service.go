package common

import (
	"context"
	"core/common/log"
	"core/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type Service struct {
	GinCtx    *gin.Context
	parentCtx context.Context
	ctx       context.Context
	span      trace.Span
	Log       *log.Log
}

func New(ginCtx *gin.Context) *Service {
	s := &Service{}
	s.SetParentCtx(ginCtx)
	s.GinCtx = ginCtx
	s.Log = log.New(s.parentCtx)
	return s
}

// SetStore 设置阿里日志logStore
func (s *Service) SetStore(storeName string) *Service {
	s.parentCtx = context.WithValue(s.parentCtx, "storeName", storeName)
	s.GinCtx.Set("storeName", storeName)
	s.Log.SetStore(storeName)
	return s
}

// Tracer 创建一个链路span
func (s *Service) Tracer(f func(public *Service, spanName string) error, spanName string) (err error) {
	ctx, span := otel.Tracer("Tracer").Start(s.GetParentCtx(), spanName)
	s.span = span
	s.ctx = ctx
	defer span.End()
	return f(s, spanName)
}

// SetParentCtx 设置父span
func (s *Service) SetParentCtx(ginCtx *gin.Context) *Service {
	if value, ok := ginCtx.Get("parentCtx"); ok {
		s.parentCtx = value.(context.Context)
	} else {
		s.parentCtx = context.Background()
	}
	return s
}

// GetParentCtx 获取父span
func (s *Service) GetParentCtx() context.Context {
	if s.parentCtx == nil {
		s.parentCtx = context.Background()
	}
	return s.parentCtx
}

// GetCtx 获取Ctx
func (s *Service) GetCtx() context.Context {
	if s.ctx == nil {
		s.ctx = s.parentCtx
	}
	return s.ctx
}

// GetSpan 获取当前span
func (s *Service) GetSpan() trace.Span {
	return s.span
}

// CreateSpan 创建一个span
func (s *Service) CreateSpan(name string) trace.Span {
	s.ctx, s.span = otel.Tracer("CreateSpan").Start(s.parentCtx, name)
	return s.span
}

func (s *Service) Failed(err error) {
	out := response.Fail
	if e, ok := err.(*response.Codes); ok {
		out = e
	} else {
		out.Msg = err.Error()
	}
	s.Log.Error("response out:", out)
	s.output(http.StatusOK, out, response.Empty{})
}

func (s *Service) Success(data interface{}) {
	s.output(http.StatusOK, response.Succ, data)
}

func (s *Service) output(httpCode int, codes *response.Codes, data interface{}) {
	out := response.Data{
		Code:  codes.Code,
		Desc:  codes.DescCh,
		Data:  data,
		Trace: s.Log.GetTraceId(),
	}
	language := s.GinCtx.GetString("language")
	if language != "" {
		out.Desc = codes.DescEn
	}
	if codes.Msg != "" {
		out.Desc += "(" + codes.Msg + ")"
	}
	s.GinCtx.JSON(httpCode, out)
}

func (s *Service) CustomOut(httpCode int, codes *response.Codes) {
	s.output(http.StatusUnauthorized, codes, "")
}
