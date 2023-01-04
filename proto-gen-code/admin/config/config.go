package config

import (
	types "core/config"
)

type Config struct {
	Mysql   *types.DB      `json:"mysql"`
	Redis   *types.Redis   `json:"redis"`
	Service *types.Service `json:"service"`
	Jaeger  *types.Jaeger  `json:"jaeger"`
	Logs    *types.Logs    `json:"logs"`
	Jwt     *types.Jwt     `json:"jwt"`
	Consul  *types.Consul  `json:"consul"`
}
