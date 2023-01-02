# ogen代码生成

[ogen的github仓库](https://github.com/ogen-go/ogen) <br>
[ogen的文档](https://ogen.dev/docs/intro) <br>
[ogen的博客](https://ogen.dev/blog) <br>

## 快速开始
```shell
go install -v github.com/ogen-go/ogen/cmd/ogen@latest

mkdir ogen

go mod init ogen

cd ogen

wget https://raw.githubusercontent.com/ogen-go/web/main/examples/petstore.yml

vi generate.go

package ogen

//go:generate go run github.com/ogen-go/ogen/cmd/ogen --target petstore --clean petstore.yml

vi Makefile

.PHONY: tidy
tidy:
	@echo "go mod tidy..."
	@go mod tidy -compat=1.9

.PHONY: generate
generate:
	@echo "go generate..."
	@go generate ./...

.PHONY: server
server:
	@echo "start server..."
	@go run main.go
	
make tidy       // 拉取项目go mod依赖
make generate   // 生成服务代码
make server     // 启动服务

// 调用服务接口
curl -X "POST" -H "Content-Type: application/json" -d "{\"name\":\"Cat\"}" http://localhost:8088/pet
```

## 整体项目结构
```text
tree .

├── Makefile                                项目Makefile文件
├── README.md                               项目README文件
├── generate.go                             项目代码生成文件generate.go
├── go.mod                                  go mod文件
├── go.sum                                  go sum文件
├── main.go                                 服务入口main.go
├── petstore                                项目生成文件目录
│   ├── oas_cfg_gen.go                服务配置文件
│   ├── oas_client_gen.go             服务client端
│   ├── oas_handlers_gen.go           服务handlers   
│   ├── oas_interfaces_gen.go         服务interfaces 
│   ├── oas_json_gen.go               服务json
│   ├── oas_middleware_gen.go         服务中间件
│   ├── oas_parameters_gen.go         服务出入参结构
│   ├── oas_request_decoders_gen.go   服务请求decode
│   ├── oas_request_encoders_gen.go   服务请求encode
│   ├── oas_response_decoders_gen.go  服务响应decode
│   ├── oas_response_encoders_gen.go  服务响应encode
│   ├── oas_router_gen.go             服务router
│   ├── oas_schemas_gen.go            服务schemas
│   ├── oas_server_gen.go             服务server端
│   ├── oas_unimplemented_gen.go      服务unimplement
│   └── oas_validators_gen.go         服务校验
└── petstore.yml                            服务基于yml文件生成代码
```

## 服务Server Handler(oas_server_gen.go)
```go
// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// AddPet implements addPet operation.
	//
	// Add a new pet to the store.
	//
	// POST /pet
	AddPet(ctx context.Context, req *Pet) (*Pet, error)
	// DeletePet implements deletePet operation.
	//
	// Deletes a pet.
	//
	// DELETE /pet/{petId}
	DeletePet(ctx context.Context, params DeletePetParams) error
	// GetPetById implements getPetById operation.
	//
	// Returns a single pet.
	//
	// GET /pet/{petId}
	GetPetById(ctx context.Context, params GetPetByIdParams) (GetPetByIdRes, error)
	// UpdatePet implements updatePet operation.
	//
	// Updates a pet in the store.
	//
	// POST /pet/{petId}
	UpdatePet(ctx context.Context, params UpdatePetParams) error
}
```

## 服务主程代码
```go
func main() {
	// Create service instance.
	service := &PetsService{
		pets: map[int64]petstore.Pet{},
	}

	// Create generated server.
	srv, err := petstore.NewServer(service)
	if err != nil {
		log.Fatal(err)
	}

	if err := http.ListenAndServe(":8088", srv); err != nil {
		log.Fatal(err)
	}
}
```

## 服务主程business(服务接口业务逻辑)
```go
type petsService struct {
	pets map[int64]petstore.Pet
	id   int64
	mux  sync.Mutex
}

func (p *petsService) AddPet(ctx context.Context, req *petstore.Pet) (*petstore.Pet, error) {
	p.mux.Lock()
	defer p.mux.Unlock()

	p.pets[p.id] = *req
	p.id++

	res := &petstore.Pet{
		ID:        petstore.NewOptInt64(int64(p.id)),
		Name:      req.Name,
		PhotoUrls: []string{"www.petstore.com"},
		Status:    petstore.NewOptPetStatus(petstore.PetStatusAvailable),
	}

	fmt.Println(res)

	return req, nil
}

func (p *petsService) DeletePet(ctx context.Context, params petstore.DeletePetParams) error {
	p.mux.Lock()
	defer p.mux.Unlock()

	delete(p.pets, params.PetId)

	return nil
}

func (p *petsService) GetPetById(ctx context.Context, params petstore.GetPetByIdParams) (petstore.GetPetByIdRes, error) {
	p.mux.Lock()
	defer p.mux.Unlock()

	pet, ok := p.pets[params.PetId]
	if !ok {
		// Return Not Found.
		return &petstore.GetPetByIdNotFound{}, nil
	}

	return &pet, nil
}

func (p *petsService) UpdatePet(ctx context.Context, params petstore.UpdatePetParams) error {
	p.mux.Lock()
	defer p.mux.Unlock()

	pet := p.pets[params.PetId]
	pet.Status = params.Status

	if val, ok := params.Name.Get(); ok {
		pet.Name = val
	}

	p.pets[params.PetId] = pet

	return nil
}
```