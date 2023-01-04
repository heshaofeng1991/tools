package config

import (
	"core/utils"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

const (
	DEV      = "dev"     //测试
	PROD     = "prod"    //生产
	SANDBOX  = "sandbox" //沙盒
	PathFile = "./config.yaml"
)

type Local struct {
	Viper  *viper.Viper
	Listen bool
	File   string
}

func defaultFileName() string {
	model := utils.GetEnvInfo("MODEL")
	name := "config-"
	switch model {
	case DEV, PROD, SANDBOX:
		name += model
	default:
		name += PROD
	}
	return "./" + name + ".yaml"
}

func NewLocal(file string) *Local {
	if file == "" {
		file = PathFile
	}
	v := viper.New()
	v.SetConfigFile(file)
	return &Local{File: file, Viper: v}
}

// GetConfig 获取配置
func (l *Local) GetConfig(configs ...interface{}) error {
	err := l.Viper.ReadInConfig()
	if err != nil {
		return err
	}
	for _, c := range configs {
		err = l.Viper.Unmarshal(c)
		if err != nil {
			return err
		}
	}
	return nil
}

func (l *Local) ListenConfig(configs ...interface{}) error {
	l.Viper.WatchConfig()
	l.Viper.OnConfigChange(func(in fsnotify.Event) {
		for _, c := range configs {
			err := l.Viper.Unmarshal(c)
			if err != nil {
				log.Print("配置修改失败")
			} else {
				log.Print("配置修改成功")
			}
		}
	})
	return nil
}

func (l *Local) GetListenStatus() bool {
	return l.Listen
}

func (l *Local) GetFile() string {
	return l.File
}
