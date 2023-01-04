package registry

import (
	"core/config"
)

type (
	// Registry 服务注册发现的接口
	Registry interface {
		// Register 注册服务
		Register(service *config.Service) error
		// Deregister 注销服务
		Deregister(service *config.Service) error
		// GetService 获取服务
		GetService(service config.Service) ([]*config.Service, error)
		// ListServices 服务列表
		ListServices() ([]*config.Service, error)
	}
)
