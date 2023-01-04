package global

import (
	conf "core/config"
	"core/initialize"
)

type config struct {
	Mysql     conf.DB      `json:"mysql"`
	Redis     conf.Redis   `json:"redis"`
	Service   conf.Service `json:"service"`
	Jaeger    conf.Jaeger  `json:"jaeger"`
	Consul    conf.Consul  `json:"consul"`
	Logs      conf.Logs    `json:"logs"`
	MasterJwt conf.Jwt     `json:"master_jwt"`
}

var (
	Config *config
	Srv    *initialize.Services
)

func init() {
	Config = new(config)
}
