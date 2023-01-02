# ORM
- [ent](https://github.com/heshaofeng1991/entgo)<br>
  - 
  - 快速开始
  ```shell
  // 安装
  go get -d entgo.io/ent/cmd/ent
  or 
  go install entgo.io/ent/cmd/ent@latest
  
  // generate.go
  //go:generate go run -mod=mod entgo.io/ent/cmd/ent generate --target ./gen --feature privacy,entql,schema/snapshot,sql/schemaconfig,sql/modifier,sql/upsert ./schema

  // generate
  go run -mod=mod entgo.io/ent/cmd/ent generate --target ./gen --feature privacy,entql,schema/snapshot,sql/schemaconfig,sql/modifier,sql/upsert ./schema
  
  // run
  go generate ./ent
  ```
  - Feature Flags
    - privacy            隐私层允许为数据库中实体的查询和变更配置隐私策略 
    - entql              运行时为不同的查询构建器提供通用和动态过滤功能
    - schema/snapshot    告诉entc(ent codegen -- 对应internal) 将最新模式的快照存储在内部包中，并在无法构建用户模式时使用它自动解决合并冲突
    - sql/schemaconfig   允许您将备用SQL数据库名称传递给模型
    - sql/modifier       允许将自定义SQL修饰符添加到构建器并在语句执行之前对其进行修改
    - sql/upsert         配置更新插入和批量更新插入逻辑
    - sql/lock           配置行级锁定
    - sql/execquery      允许使用底层驱动程序的ExecContext/方法执行语句

  - 项目结构
  ```text
  tree .
  
  ├── ent
  │   ├── gen           ent数据库操作代码生成目录
  │   ├── internal      ent生成内部依赖(最新模式的快照存储在内部包)   
  │   ├── rule          ent生成应遵循的规则        
  │   ├── schema        ent代码生成依赖数据库schema      
  │   └── viewer        ent代码生成viewer(查看器描述查询/变更,查看器上下文)   
  │   ├── generate.go   ent代码生成generate文件    
  ├── Makefile                项目Makefile
  ```
  - Makefile
  ```shell
  GO_FILES=$(shell find . -type f -name '*.go')

  .PHONY: ent
  ent:
  @echo "Generating ent code..."
  @rm -rf ./ent/gen
  @mkdir -p ./ent/gen
  @cp -r ./ent/internal ./ent/gen
  go generate ./ent
  @cp -r ./ent/gen/internal ./ent
    
  .PHONY: ent
  tidy:
  @echo "TIDYING CODE..."
  @go mod tidy -compat=1.19
    
  .PHONY: all
  all: ent tidy
  ```
- gorm
- xorm
- sqlx
- ...