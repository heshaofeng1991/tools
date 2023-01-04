package common

import (
	"core/config"
	"core/initialize"
	"core/utils"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func NewGrpc(reg func(server *grpc.Server)) *grpc.Server {
	server := grpc.NewServer(utils.OpenTracingServerInterceptor())
	reg(server)
	return server
}

func GrpcRun(server *grpc.Server, services *initialize.Services) {
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())
	err := services.Registry.Register(config.Global.Service)
	if err != nil {
		zap.S().Fatalf("注册失败:%s", err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", config.Global.Service.Host, config.Global.Service.Port))
	if err != nil {
		panic("failed to listen:" + err.Error())
	}
	go func() {
		err = server.Serve(lis)
		if err != nil {
			panic("failed to start grpc:" + err.Error())
		}
	}()
	zap.S().Info("服务已启动")
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	//服务注销
	err = services.Registry.Deregister(config.Global.Service)
	if err != nil {
		zap.S().Fatalf("注销失败:%s", err)
	}
	zap.S().Info("Server exiting")
}
