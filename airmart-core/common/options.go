package common

/*import (
	"airmart-core/registry"
	"airmart-core/types"
	"airmart-core/utils"
	"google.golang.org/grpc"
)

type Option func(*Options)

type Options struct {
	registry registry.Registry //服务注册 默认consul
	server   *grpc.Server      //grpc服务
}

func defaultOptions(config interface{}) (*Options, error) {
	consul, err := utils.ReflectRestore(config, types.Consul{})
	if err != nil {
		return nil, err
	}
	return &Options{
		registry: registry.NewConsul(consul.(types.Consul)),
		server:   &grpc.Server{},
	}, nil
}

// WithRegistry 服务发现
func WithRegistry(registry registry.Registry) Option {
	return func(options *Options) {
		options.registry = registry
	}
}

// WithGrpcServer grpc服务
func WithGrpcServer(server *grpc.Server) Option {
	return func(options *Options) {
		options.server = server
	}
}
*/
