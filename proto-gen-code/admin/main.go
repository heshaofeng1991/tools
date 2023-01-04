package main

import (
	"admin/global"
	"admin/router"
	"core/common"
	"core/initialize"
	"log"
)

func main() {
	// 初始化服务配置.
	err := initialize.Config(true, global.Config)
	if err != nil {
		log.Fatalf("加载配置失败 err:%s", err.Error())
	}

	// 初始化服务依赖.
	global.Srv, err = initialize.DefaultServices()
	if err != nil {
		log.Fatalf("服务启动失败 err:%s", err.Error())
	}

	// 启动http服务.
	common.Run(router.Init(), global.Srv)

	//// 创建rpc服务.
	//server := grpc.NewServer(utils.OpenTracingServerInterceptor())
	//
	//// rpc调用方式.
	//admin.RegisterAdminServer(server, &handler.AdminServer{})
	//
	//// 启动rpc服务.
	//common.GrpcRun(server, adminInit.Services)
}
