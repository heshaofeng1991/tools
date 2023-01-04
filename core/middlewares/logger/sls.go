package logger

import (
	"core/config"
	"core/utils"
	"encoding/json"
	"time"

	sls "github.com/aliyun/aliyun-log-go-sdk"
	"github.com/aliyun/aliyun-log-go-sdk/producer"
	"github.com/pkg/errors"
)

/*var (
	instanceSls *Sls
	once        sync.Once
)

func InstanceSls() *Sls {
	once.Do(func() {
		instanceSls = NewSls()
	})
	return instanceSls
}*/

type Sls struct {
	Conf             *config.Sls
	Client           sls.ClientInterface
	producerInstance *producer.Producer
	Ip               string
}

func (s *Sls) Write(p []byte) (n int, err error) {
	result := map[string]string{}
	err = json.Unmarshal(p, &result)
	if err != nil {
		return n, err
	}
	log := producer.GenerateLog(uint32(time.Now().Unix()), result)
	return 1, s.producerInstance.SendLog(s.Conf.Name, s.Conf.StoreName, config.Global.Service.Cluster, s.Ip, log)
}

func (s *Sls) Sync() error {
	return nil
}

func (s *Sls) Close() error {
	s.producerInstance.SafeClose() // 安全关闭
	return nil
}

func NewSls(conf *config.Sls) *Sls {
	s := &Sls{}
	s.Conf = conf
	s.Client = sls.CreateNormalInterface(s.Conf.Host, s.Conf.Id, s.Conf.Secret, "")
	producerConfig := producer.GetDefaultProducerConfig()
	producerConfig.Endpoint = s.Conf.Host
	producerConfig.AccessKeyID = s.Conf.Id
	producerConfig.AccessKeySecret = s.Conf.Secret
	s.producerInstance = producer.InitProducer(producerConfig)
	s.producerInstance.Start()
	s.Ip = utils.GetOutboundIP()
	return s
}

func (s *Sls) Default() (*Sls, error) {
	if err := s.CreateProject(); err != nil {
		return s, errors.Wrap(err, "创建sls Project 失败")
	}
	_, err := s.CreateLogStore(s.Conf.StoreName)
	if err != nil {
		return s, errors.Wrap(err, "创建logStore失败")
	}
	if err = s.CreateIndex(s.Conf.StoreName, map[string]sls.IndexKey{
		"level": {
			Token:         []string{" "},
			CaseSensitive: false,
			Type:          "text",
		},
		"trace_id": {
			Token:         []string{" "},
			CaseSensitive: false,
			Type:          "text",
		},
		"msg": {
			Token:         []string{" "},
			CaseSensitive: false,
			Type:          "text",
		},
		"uri": {
			Token:         []string{" "},
			CaseSensitive: false,
			Type:          "text",
		},
		"ts": {
			Token:         []string{" "},
			CaseSensitive: false,
			Type:          "text",
		},
	}); err != nil {
		return s, errors.Wrap(err, "创建默认sls索引失败")
	}
	return s, nil
}

func (s *Sls) CreateLogStore(storeName string) (*Sls, error) {
	exist, err := s.Client.CheckLogstoreExist(s.Conf.Name, storeName)
	if err != nil {
		return s, err
	}
	if exist {
		return s, nil
	}
	return s, s.Client.CreateLogStore(s.Conf.Name, storeName, 2, 2, true, 64)
}

func (s *Sls) CreateProject() error {
	exist, err := s.Client.CheckProjectExist(s.Conf.Name)
	if err != nil {
		return err
	}
	if exist {
		return nil
	}
	_, err = s.Client.CreateProject(s.Conf.Name, "")
	return nil
}

//	CreateIndex sls.IndexKey{
//			Token:         []string{" "},
//			CaseSensitive: false,
//			Type:          "text", text/long/double/json
//		}
func (s *Sls) CreateIndex(storeName string, fields map[string]sls.IndexKey) error {
	_, err := s.Client.GetIndex(s.Conf.Name, storeName)
	if err == nil {
		return nil
	}
	clientError := sls.NewClientError(err)
	if clientError.Code != "IndexConfigNotExist" {
		return err
	}
	index := sls.Index{
		Keys: fields,
		Line: &sls.IndexLine{
			Token:         []string{",", ":", " "},
			CaseSensitive: false,
			IncludeKeys:   []string{},
			ExcludeKeys:   []string{},
		},
	}
	return s.Client.CreateIndex(s.Conf.Name, storeName, index)
}
