package kafka

import (
	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
)

// ProducerConfig 生产者配置.
type ProducerConfig struct {
	Broker     []string
	SaslEnable bool
	User       string
	Password   string
}

// ConsumerGroupConfig 消费者配置.
type ConsumerGroupConfig struct {
	Topic      []string
	Broker     []string
	SaslEnable bool
	User       string
	Password   string
	Group      string
}

// Client Kafka客户端.
type Client struct {
	ctx           context.Context
	asyncProducer sarama.AsyncProducer
	syncProducer  sarama.SyncProducer
	consumer      *cluster.Consumer
	conf          *Kafka
	onReceiveMsg  func(message *sarama.ConsumerMessage)
}

type Kafka struct {
	Address    []string `json:"address"`     // 地址列表
	Username   string   `json:"username"`    // 用户名
	Password   string   `json:"password"`    // 密码
	SaslEnable bool     `json:"sasl_enable"` // 是否开启鉴权
	Consume    Consume  `json:"consume"`     // 消费者信息
}

type Consume struct {
	Group  string   `json:"group"`
	Topics []string `json:"topics"`
}
