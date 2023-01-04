### core 使用手册

### 拉取依赖需要的设置的参数

```shell
go env -w GO111MODULE=on
```
***

#### 目录说明
```text
common --  服务启动的入口   
config --  配置 加载配置文件 nacos 加载配置  
middlewares 中间件  
initialize 初始化服务   
registry   服务发现  
response   统一接口响应返回格式  
utils      工具
```

***
### 配置文件
```shell
go run  main.go -confFile ./config-dev.yaml  默认格式 ./config.yaml
```

***
### 服务启动
```text
默认服务启动 mysql redis log jaeger
```

```
func main() {
    //加载配置
	if err := initialize.Config(true, global.Config); err != nil {
		log.Fatalf("加载配置失败 err:%s", err.Error())
	}
	
	//启动默认服务 mysql redis zaplog jaeger
	global.Srv, err = initialize.DefaultServices()
	if err != nil {
		log.Fatalf("服务启动失败 err:%s", err.Error())
	}
	
	common.Run(router.Init(), global.Srv)
}

```

***
#### 自定义服务启动 
```
    //加载配置
	if err := initialize.Config(true, global.Config); err != nil {
		log.Fatalf("加载配置失败 err:%s", err.Error())
	}
	
	global.Srv, err = initialize.Start(initialize.Mysql()) 
	if err != nil {
		log.Fatalf("服务启动失败 err:%s", err.Error())
	}
	
	common.Run(router.Init(), global.Srv)
```