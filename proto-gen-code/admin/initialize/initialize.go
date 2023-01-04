package initialize

import (
	"flag"
	"log"
	"os"
	"runtime"
	"strings"

	"admin/global"
	"airmart-core/initialize"
)

var Services *initialize.Services

// 初始化操作
func init() {
	var (
		conf string
		err  error
	)

	baseDir, _ := os.Getwd()
	if runtime.GOOS == "windows" {
		baseDir = strings.ReplaceAll(baseDir, "\\", "/")
	}

	// 这里预留，可能需要自定义需要的启动项.
	flag.StringVar(&conf, "conf", baseDir+"/config/config-dev.yaml", "-conf config/config-dev.yaml")
	flag.Parse()

	if err = initialize.Config(true, global.Config); err != nil {
		log.Fatalf("初始化服务配置文件失败 err:%s", err.Error())
	}

	// 初始化服务资源配置.
	services, err := initialize.DefaultServices()
	if err != nil {
		log.Fatalf("初始化服务资源失败 err:%s", err.Error())
	}

	Services = services
	{
		global.Srv.DB = Services.DB
		// global.Srv.Redis = Services.Redis
	}
}
