package nacos

import (
	"core/config"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"go.uber.org/zap"
)

var (
	logDir   = "./tmp/nacos/log"
	CacheDir = "./tmp/nacos/cache"
)

type NaCos struct {
	clientParam vo.NacosClientParam
	conf        *config.NaCos
}

func New(conf *config.NaCos) *NaCos {
	clientConfig := constant.ClientConfig{
		NamespaceId:         conf.Namespace,
		TimeoutMs:           50000,
		Username:            conf.User,
		Password:            conf.Password,
		NotLoadCacheAtStart: true,
		LogDir:              logDir,
		CacheDir:            CacheDir,
		LogLevel:            "debug",
	}
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: conf.Host,
			Port:   conf.Port,
		},
	}
	return &NaCos{
		conf:        conf,
		clientParam: vo.NacosClientParam{ClientConfig: &clientConfig, ServerConfigs: serverConfigs},
	}
}

func (n *NaCos) Register(service *config.Service) error {
	namingClient, err := clients.NewNamingClient(n.clientParam)
	if err != nil {
		return err
	}
	success, err := namingClient.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          service.Host,
		Port:        service.Port,
		ServiceName: service.Name,
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Metadata:    service.Metadata,
		ClusterName: service.Cluster,
		GroupName:   service.Group,
	})

	if err != nil {
		return err
	}
	if !success {
		return errors.New("服务注册失败")
	}

	return nil
}

func (n *NaCos) Deregister(service *config.Service) error {
	namingClient, err := clients.NewNamingClient(n.clientParam)
	if err != nil {
		return err
	}
	success, err := namingClient.DeregisterInstance(vo.DeregisterInstanceParam{
		Ip:          service.Host,
		Port:        service.Port,
		ServiceName: service.Name,
		Ephemeral:   true,
		Cluster:     service.Cluster,
		GroupName:   service.Group,
	})
	if err != nil || !success {
		return fmt.Errorf("服务注销失败!err:%v", err)
	}
	return nil
}

func (n *NaCos) GetService(service config.Service) (services []*config.Service, err error) {
	namingClient, err := clients.NewNamingClient(n.clientParam)
	if err != nil {
		return
	}

	param := vo.GetServiceParam{
		Clusters:    []string{service.Cluster},
		ServiceName: service.Name,
		GroupName:   service.Group,
	}
	defer namingClient.CloseClient()
	result, err := namingClient.GetService(param)
	if err != nil {
		return
	}

	for _, s := range result.Hosts {
		if s.Healthy == true {
			services = append(
				services,
				&config.Service{
					Name:     s.ServiceName,
					Metadata: nil,
					Host:     s.Ip,
					Port:     s.Port,
					Group:    result.GroupName,
					Cluster:  s.ClusterName,
				},
			)
		}
	}

	return services, nil
}

func (n *NaCos) ListServices() ([]*config.Service, error) {
	return nil, nil
}

func (n *NaCos) GetConfig(config interface{}) error {
	configClient, err := clients.NewConfigClient(n.clientParam)
	if err != nil {
		return err
	}

	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: n.conf.DataId,
		Group:  n.conf.Group},
	)

	if err != nil {
		return err
	}

	if content == "" {
		return errors.New("nacos获取配置为空！请检查配置")
	}

	return json.Unmarshal([]byte(content), &config)
}

func (n *NaCos) ListenConfig(config interface{}) error {
	configClient, err := clients.NewConfigClient(n.clientParam)
	if err != nil {
		return err
	}
	err = configClient.ListenConfig(vo.ConfigParam{
		DataId: n.conf.DataId,
		Group:  n.conf.Group,
		OnChange: func(namespace, group, dataId, data string) {
			//TODO 这里可能要判断data为空或空json
			err = json.Unmarshal([]byte(data), &config)
			if err != nil {
				zap.S().Errorf("namespace:%s group:%s dataId:%s data %s \n 动态修改配置失败 err:%s",
					namespace, group, dataId, data, err)

			} else {
				zap.S().Info("配置中心修改成功")
			}
		},
	})
	return err
}
