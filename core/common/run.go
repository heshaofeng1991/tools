package common

import (
	"context"
	"core/config"
	"core/initialize"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"go.uber.org/zap"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func Run(router http.Handler, services *initialize.Services) {
	//服务注册
	err := services.Registry.Register(config.Global.Service)
	if err != nil {
		zap.S().Fatalf("注册失败:%s", err)
	}

	h2Handler := h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 &&
			strings.HasPrefix(r.Header.Get("Content-Type"), "application/grpc") && services.GrpcSrv != nil {
			services.GrpcSrv.ServeHTTP(w, r)
		} else {
			router.ServeHTTP(w, r)
		}
	}), &http2.Server{})

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Global.Service.Port),
		Handler: h2Handler,
	}
	zap.S().Infof("服务启动-端口：%d", config.Global.Service.Port)
	go func() {
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			zap.S().Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	//服务注销
	err = services.Registry.Deregister(config.Global.Service)
	if err != nil {
		zap.S().Fatalf("注销失败:%s", err)
	}

	zap.S().Infof("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		zap.S().Info("Server forced to shutdown:", err)
	}

	zap.S().Info("Server exiting")
}
