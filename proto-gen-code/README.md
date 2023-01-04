# proto-gen-code项目说明
---

## 项目结构
- admin（web 同理）
  - proto-gen-code 管理端服务代码
  - admin服务结构
    - config 配置定义及文件
    - global 全局配置依赖定义
    - handler 服务handler具体业务逻辑实现
    - initialize 服务依赖初始化
    - router 服务路由初始化
    - service 服务持久化层具体实现
    - main.go 服务入口
    - Makefile 服务命令
    - model 服务持久化层model定义（所属服务自己的model）
- web服务
  - proto-gen-code H5服务代码
- proto代码生成
  - admin （web 同理）
    - admin.proto文件 proto定义
    - admin.pb.go 出入参生成文件
    - admin.validator.go 参数校验生成文件
    - admin_grpc.pb.go rpc服务生成文件
    - admin_http.pb.go http服务生成文件
  - web
- model 服务持久化层model定义(数据库表结构定义，这里应该是公共model)
- Makefile 项目依赖及代码生成命令

## 项目依赖设置 go mod && git config
```shell
go env -w GO111MODULE=on
```

## 项目代码生成说明
- 项目生成依赖
```shell
go get -u github.com/gogo/protobuf@v1.3.2 // go install也可
go get -u github.com/mwitkow/go-proto-validators@v0.3.2 // go install也可
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install github.com/mwitkow/go-proto-validators/protoc-gen-govalidators@latest
go install github.com/heshaofeng1991/protoc-gen-go-http@latest (个人封装的http方式，感兴趣可以研究学习，相互进步)
go install github.com/gogo/protobuf/protoc-gen-gogofaster@latest
```
- 项目代码生成命令
```shell
make proto 
or 
protoc -I=. -I=${GOPATH}/pkg/mod \
    --gogofaster_out=. --gogofaster_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    --govalidators_out=. --govalidators_opt=paths=source_relative \
    --go-http_out=. --go-http_opt=paths=source_relative \
    proto/admin/admin.proto
```
- notes
  - Windows make设置说明
  ```toml
  1. 这里是针对Windows的说明，Mac,Linux可以不用理会，可能需要自己安装下make命令,根据自己的系统版本去下载
  https://sourceforge.net/projects/mingw-w64/files/Toolchains%20targetting%20Win64/
  2. Windows系统在bin目录下修改mingw64_make.exe为make.exe
  3. 设置一下bin目录系统环境变量
  4. Goland工具设置make运行make.exe的目录
  ```
  - Windows代码生成设置说明，主要是这个库没有release
  ```toml
  1. 在自己的${GOPATH}/pkg/mod新建依赖文件夹google/api google/protobuf
  2. 在自己依赖的mod目录下git clone git@github.com:googleapis/googleapis.git
  3. 把依赖的googleapis/google目录下api都拷贝过去
  4. 把依赖的github.com/gogo/protobuf@v1.3.2/protobuf/google目录下protobuf都拷贝过去
  ```
  - Goland设置
  ```toml
  1. 主要是为了解决proto文件的包import问题
  2. 打开Goland设置Protocol Buffers 引入${GOPATH}/pkg/mod
  ```

## 快速开始
```shell
git clone git@github.com:heshaofeng1991/proto-gen-code.git
cd proto-gen-code
make tidy 
or 
go mod tidy -compat=1.9
make server 
or 
go run main.go -confFile ./config/config-dev.yaml
```

## 其它组件
```text
Go1.19 Golang最新版本
Consul 服务注册，服务发现
Jaeger Trace链路追踪，监控
Mysql  数据库
Redis  Redis
Nacos  配置中心
JWT    鉴权
Logger zap
Kafka  消息队列
OSS    对象存储
ElasticSearch
```