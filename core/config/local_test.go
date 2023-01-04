package config

import (
	"airmartCore/middlewares/nacos"
	"core/types"
	"fmt"
	"testing"
)

type Conf struct {
	Mysql   types.DB      `json:"mysql"`
	Redis   types.Redis   `json:"redis"`
	Service types.Service `json:"common"`
	Oss     types.Oss     `json:"oss"`
	Jaeger  types.Jaeger  `json:"jaeger"`
	Logs    types.Logs    `json:"logs"`
	Jwt     types.Jwt     `json:"jwt"`
}

func TestLocal_GetConfig(t *testing.T) {
	local := NewLocal("E:\\zhoushuping\\go\\airmart\\config-dev.yaml")

	NaCosConf := &types.NaCos{}
	err := local.GetConfig(NaCosConf)
	if err != nil {
		t.Error(err)
	}
	config := &Conf{}
	n := nacos.New(NaCosConf)
	err = n.GetConfig(config)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(config)

}
