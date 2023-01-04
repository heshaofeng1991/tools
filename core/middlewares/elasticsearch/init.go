package elasticsearch

import (
	"context"
	"core/config"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/olivere/elastic/v7"
)

var (
	ctx = context.Background()
)

// Client es的连接.
type Client struct {
	Client *elastic.Client
}

// NewESClient 初始化es连接实例..  host := []string{"http://10.0.6.245:9200","http://10.0.6.246:9200","http://10.0.6.247:9200"}
func NewESClient(conf config.Elasticsearch) (*Client, error) {
	length := len(conf.Address)
	if length == 0 {
		return nil, fmt.Errorf("cluster Not Es Node")
	}

	// 创建新的连接.
	for i, ip := range conf.Address {
		// 判断是不是最后一个节点ip.
		if (length - 1) != i {
			es, err := Connect(ip, conf.Username, conf.Password)
			// 如果连接出错，则跳过.
			if err != nil {
				fmt.Println(err)

				continue
			}

			return es, nil
		} else {
			es, err := Connect(ip, conf.Username, conf.Password)
			if err != nil {
				return nil, err
			}

			return es, nil
		}
	}

	return nil, nil
}

/*
	Client连接地址
	Client账号/密码
	监控检查
	失败重试次数
	gzip设置
*/
// Connect 连接.
func Connect(host, name, password string) (*Client, error) {
	client, err := elastic.NewClient(
		// host是可变参数 http://10.0.1.1:9200", "http://10.0.1.2:9200.
		elastic.SetURL(host),
		// 基于http base auth验证机制的账号和密码. 这里跟上面的host应该都是全局配置.
		elastic.SetBasicAuth(name, password),
		// 启用gzip压缩.
		elastic.SetGzip(true),
		// 设置监控检查时间间隔.
		elastic.SetHealthcheckInterval(10*time.Second),
		// 设置错误日志输出.
		elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		// 设置info日志输出.
		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)),
	)
	if err != nil {
		return nil, err
	}

	info, code, err := client.Ping(host).Do(ctx)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Client returned with code %d and version %s\n", code, info.Version.Number)

	esVersion, err := client.ElasticsearchVersion(host)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Client version %s\n", esVersion)

	return &Client{
		Client: client,
	}, nil
}
