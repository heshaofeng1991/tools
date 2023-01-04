package elasticsearch

import (
	"core/config"
	"log"
	"os"

	elastic "github.com/olivere/elastic/v7"
)

func New(conf config.Elasticsearch) (client *elastic.Client, err error) {
	client, err = elastic.NewSimpleClient(
		elastic.SetURL(conf.Address...),                            // 服务地址
		elastic.SetBasicAuth(conf.Username, conf.Password),         // 账号密码
		elastic.SetErrorLog(log.New(os.Stderr, "", log.LstdFlags)), // 设置错误日志输出
		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)),  // 设置info日志输出
	)
	return
}
