### Golang Combo 项目依赖说明(Domain-Driven-Design 领域驱动设计实践)。

#### 这里仅仅只是一个简单的示例，后续会完善更好一些

#### 项目结构
- gen            goa.design生成服务器代码、客户端代码和文档工具。 <br>
- ent            go ent 生成代码orm框架(mysql)  <br>
- reviewdog      项目代码规范（golangci-lint,.golangci.yml)  <br>
- golang版本      go 1.19   <br>
- cmd            combo服务端与客户端 <br>
- design         代码生成依赖文件 <br>
  - model        服务定义出入参 <br>
  - service      具体服务定义 <br>
  - api.go       服务入口定义 <br>
- domain         服务领域层 
- infrastructure 服务仓储层（DB，Redis,Memcached等等）
  -  repository  服务repo实现
- internal       服务内部依赖
  - env          环境变量
  - dao          服务DB
- Makefile       服务makefile

#### goa

###### goa.design说明

Goa 是一个用于编写微服务的 Go 框架，它通过提供从服务器代码、客户端代码和文档派生的单一事实来源来促进最佳实践。 Goa 生成的代码遵循干净的架构模式，其中为传输、端点和业务逻辑层生成可组合的模块。 Goa 包含中间件、插件和其他补充功能，可以与生成的代码一起使用，以有效的方式实现完整的微服务。 通过使用 goa 开发微服务，实现者不必担心文档不同步，因为 Goa 负责为基于 HTTP 的服务生成 OpenAPI 规范，为基于 gRPC 的服务生成 gRPC 协议缓冲区文件（或者如果服务同时支持两者，则两者都支持运输）。

###### goa.design Link

[goa](https://goa.design/)

###### goa 项目代码生成（cmd/oms目录下不建议使用，仅供参考）
cmd/oms目录下不建议使用说明 ： 目前项目的代码结构已趋于稳定，代码也遵循了golangci-lint规范，重新生成需要重新修改改变的文件

goa example github.com/NextSmartShip/openapi/design

#### entgo

###### entgo ORM Link

[ent](https://entgo.io/)

#### 项目cmdline

###### 项目相关命令

[cmd](Makefile)

#### 其他

###### 安利一个好用的工具（数据库SQL转entgo schema生成代码）
```shell
go install github.com/miaogaolin/sql2ent@latest

sql2ent mysql ddl -src mysql.sql -dir "./ent/schema"
```