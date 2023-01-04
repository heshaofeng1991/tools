package initialize

import (
	"core/config"
	"core/middlewares/elasticsearch"
	"core/middlewares/jaeger"
	"core/middlewares/kafka"
	logger "core/middlewares/logger"
	"core/middlewares/mysql"
	"core/middlewares/redis"
	"core/registry"
	"core/utils"
	"encoding/json"
	"errors"
	"flag"
	"log"

	"github.com/hashicorp/consul/api"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type Service func(*Services) error

type Services struct {
	Ctx           context.Context
	DB            *gorm.DB
	Redis         *redis.Client
	NaCosConfCli  config_client.IConfigClient
	Registry      registry.Registry
	ConsulCli     *api.Client
	GrpcSrv       *grpc.Server
	Sls           *logger.Sls
	Kafka         *kafka.Client
	ElasticSearch *elasticsearch.Client
	Sms           *logger.Sms
}

var (
	confFile = flag.String("confFile", "", "加载配置文件路径")
	envNaCos = flag.Bool("envNaCos", false, "从环境变量读出naCos配置")
	naCos    = flag.String("naCos", "", "naCos配置")
)

func new() *Services {
	return &Services{
		Ctx:      context.Background(),
		Registry: &registry.Empty{},
	}
}

// TODO 后面是使用变量还是json
func getEnvNaCosConfig() (*config.NaCos, error) {
	info := utils.GetEnvInfo("NACOS_INFO")
	if info == "" {
		return nil, errors.New("没有配置nacos环境变量信息")
	}
	c := &config.NaCos{}
	err := json.Unmarshal([]byte(info), c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func Config(listen bool, conf interface{}) error {
	var (
		err       error
		naCosConf *config.NaCos
		local     *config.Local
	)
	flag.Parse()

	if *envNaCos {
		naCosConf, err = getEnvNaCosConfig()
		if err != nil {
			return err
		}
		goto NACOS
	}

	//从本地读取配置
	local = config.NewLocal(*confFile)
	if err != nil {
		return err
	}

	err = local.GetConfig(config.Global, conf)
	if err != nil {
		return err
	}

	naCosConf = config.Global.NaCos
	//如果本地配置开启了nacos，还是从nacos读
	if naCosConf.Open {
		goto NACOS
	}

	log.Println("从本地获取配置")
	if listen {
		return local.ListenConfig(config.Global, conf)
	}

NACOS:
	log.Println("从naCos获取配置")
	cos := config.NewNaCos(naCosConf)
	err = cos.GetConfig(conf, config.Global)
	if err != nil {
		return err
	}
	if listen {
		return cos.ListenConfig(conf, config.Global)
	}
	return nil
}

// Logger zap日志 要使用sls先要启动sls服务
func Logger(opt ...logger.Option) Service {
	return func(s *Services) error {
		if s.Sls != nil {
			opt = append(opt, logger.WithAddSync(s.Sls))
		}
		return logger.InitLogger(config.Global.Logs, opt...)
	}
}

// Start 服务启动
func Start(opts ...Service) (s *Services, err error) {
	return new().start(opts...)
}

// DefaultServices 默认服务
func DefaultServices(opts ...Service) (s *Services, err error) {
	sls, err := logger.NewSls(config.Global.Sls).Default()
	if err != nil {
		return
	}
	err = logger.InitLogger(config.Global.Logs, logger.WithAddSync(sls))
	if err != nil {
		return
	}

	db, err := mysql.New(config.Global.Mysql)
	if err != nil {
		return
	}

	redisCli, err := redis.New(config.Global.Redis)
	if err != nil {
		return
	}

	_, err = jaeger.New(config.Global.Jaeger, config.Global.Service)
	if err != nil {
		return
	}

	srv := new()
	srv.DB = db
	srv.Redis = redisCli
	srv.Sls = sls
	return srv.start(opts...)
}

func Mysql(opts ...mysql.Option) Service {
	return func(s *Services) error {
		db, err := mysql.New(config.Global.Mysql, opts...)
		s.DB = db
		return err
	}
}

func Redis() Service {
	return func(s *Services) error {
		db, err := redis.New(config.Global.Redis)
		s.Redis = db
		return err
	}
}

func Consul() Service {
	return func(s *Services) error {
		conf := api.DefaultConfig()
		conf.Address = config.Global.Consul.Host
		client, err := api.NewClient(conf)
		if err != nil {
			return err
		}

		s.Registry = registry.NewConsul(client, config.Global.Consul)
		s.ConsulCli = client
		return nil
	}
}

func Jaeger() Service {
	return func(services *Services) error {
		_, err := jaeger.New(config.Global.Jaeger, config.Global.Service)
		return err
	}
}

// GrpcSrv grpc服务
func GrpcSrv(server *grpc.Server) Service {
	return func(s *Services) error {
		s.GrpcSrv = server
		return nil
	}
}

func Sls() Service {
	return func(s *Services) error {
		var err error
		s.Sls, err = logger.NewSls(config.Global.Sls).Default()
		return err
	}
}

// Start 服务启动
func (srv *Services) start(opts ...Service) (s *Services, err error) {
	for _, o := range opts {
		err = o(srv)
		if err != nil {
			return
		}
	}
	return srv, nil
}

func Kafka() Service {
	return func(s *Services) error {
		syncProducer, err := kafka.InitSyncProducer(&kafka.ProducerConfig{
			Broker:     config.Global.Kafka.Address,
			SaslEnable: config.Global.Kafka.SaslEnable,
			User:       config.Global.Kafka.Username,
			Password:   config.Global.Kafka.Password,
		})
		if err != nil {
			return err
		}
		asyncProducer, err := kafka.InitAsyncProducer(&kafka.ProducerConfig{
			Broker:     config.Global.Kafka.Address,
			SaslEnable: config.Global.Kafka.SaslEnable,
			User:       config.Global.Kafka.Username,
			Password:   config.Global.Kafka.Password,
		})
		if err != nil {
			return err
		}

		s.Kafka = kafka.NewClient(s.Ctx, config.Global.Kafka, syncProducer, asyncProducer)
		return nil
	}
}

func ElasticSearch() Service {
	return func(s *Services) error {
		client, err := elasticsearch.NewESClient(config.Elasticsearch{
			Address:  config.Global.ElasticSearch.Address,
			Username: config.Global.ElasticSearch.Username,
			Password: config.Global.ElasticSearch.Password,
		})
		if err != nil {
			return err
		}

		s.ElasticSearch = client

		return nil
	}
}

func Sms() Service {
	return func(s *Services) error {
		var err error
		s.Sms, err = logger.NewSms(config.Global.Sms)
		return err
	}
}
