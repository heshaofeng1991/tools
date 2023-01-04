package config

import (
	"errors"
	"strings"

	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	logDir   = "./tmp/nacos/log"
	CacheDir = "./tmp/nacos/cache"
)

type NaCosObj struct {
	Listen      bool
	File        string
	clientParam vo.NacosClientParam
	conf        *NaCos
	Viper       *viper.Viper
}

func NewNaCos(conf *NaCos) *NaCosObj {
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
	return &NaCosObj{
		conf:        conf,
		clientParam: vo.NacosClientParam{ClientConfig: &clientConfig, ServerConfigs: serverConfigs},
	}
}

type data struct {
	Consul *Consul
}

func (n *NaCosObj) GetConfig(configs ...interface{}) (err error) {
	configClient, err := clients.NewConfigClient(n.clientParam)
	if err != nil {
		return
	}

	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: n.conf.DataId,
		Group:  n.conf.Group},
	)

	if err != nil {
		return
	}

	if content == "" {
		return errors.New("nacos获取配置为空！请检查配置")
	}

	n.Viper = viper.New()
	n.Viper.SetConfigType("yaml")
	err = n.Viper.ReadConfig(strings.NewReader(content))
	if err != nil {
		return
	}
	for _, c := range configs {
		err = n.Viper.Unmarshal(c)
		if err != nil {
			return err
		}
	}

	return nil
}

func (n *NaCosObj) ListenConfig(configs ...interface{}) error {
	configClient, err := clients.NewConfigClient(n.clientParam)
	if err != nil {
		return err
	}
	err = configClient.ListenConfig(vo.ConfigParam{
		DataId: n.conf.DataId,
		Group:  n.conf.Group,
		OnChange: func(namespace, group, dataId, data string) {
			for _, c := range configs {
				err = n.Viper.Unmarshal(c)
				if err != nil {
					zap.S().Errorf("namespace:%s group:%s dataId:%s data %s \n 动态修改配置失败 err:%s",
						namespace, group, dataId, data, err)
				} else {
					zap.S().Info("配置中心修改成功")
				}
			}
		},
	})
	return err
}

func (n *NaCosObj) GetListenStatus() bool {
	return n.Listen
}

func (n *NaCosObj) GetFile() string {
	return n.File
}

func (n *NaCosObj) SetConfig(conf *NaCos) *NaCosObj {
	n.conf = conf
	return n
}
