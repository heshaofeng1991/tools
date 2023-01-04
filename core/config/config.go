package config

import (
	"github.com/spf13/viper"
)

type Config interface {
	GetConfig(...interface{}) error    //获取配置（要传引用）
	ListenConfig(...interface{}) error //监听配置（要传引用）
	GetListenStatus() bool
	GetFile() string
}

type global struct {
	Mysql         *DB
	Oss           *Oss
	Jwt           *Jwt
	Service       *Service
	Consul        *Consul
	NaCos         *NaCos
	Kafka         *Kafka
	Redis         *Redis
	Logs          *Logs
	Jaeger        *Jaeger
	Sls           *Sls
	ElasticSearch *Elasticsearch
	Sms           *Sms
}

var (
	Global = &global{
		Mysql:         &DB{},
		Oss:           &Oss{},
		Jwt:           &Jwt{},
		Service:       &Service{},
		Consul:        &Consul{},
		NaCos:         &NaCos{},
		Kafka:         &Kafka{},
		Redis:         &Redis{},
		Logs:          &Logs{},
		Jaeger:        &Jaeger{},
		ElasticSearch: &Elasticsearch{},
		Sms:           &Sms{},
		Sls:           &Sls{},
	}
)

func Get(c Config, listen bool, configs ...interface{}) error {
	err := c.GetConfig()
	if err != nil {
		return err
	}

	if !listen {
		return nil
	}
	return c.ListenConfig(configs)
}

func ListenConfig(c Config, dsts ...interface{}) error {
	for _, dst := range dsts {
		err := c.ListenConfig(dst)
		if err != nil {
			return err
		}
	}
	return nil
}

func SetConfig(v *viper.Viper, dsts ...interface{}) error {
	for _, dst := range dsts {
		err := v.Unmarshal(&dst)
		if err != nil {
			return err
		}
	}
	return nil
}
