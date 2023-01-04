package kafka

import (
	"context"
	conf "core/config"

	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
)

type ProducerConfig struct {
	Broker     []string
	SaslEnable bool
	User       string
	Password   string
}

type Client struct {
	ctx           context.Context
	asyncProducer sarama.AsyncProducer
	syncProducer  sarama.SyncProducer
	consumer      *cluster.Consumer
	conf          *conf.Kafka
	onReceiveMsg  func(message *sarama.ConsumerMessage)
}

func NewClient(ctx context.Context, conf *conf.Kafka, kafkaSyncProducer sarama.SyncProducer,
	kafkaAsyncProducer sarama.AsyncProducer) *Client {
	return &Client{
		ctx:           ctx,
		conf:          conf,
		asyncProducer: kafkaAsyncProducer,
		syncProducer:  kafkaSyncProducer,
	}
}

// InitAsyncProducer 异步生产者.
func InitAsyncProducer(conf *ProducerConfig) (producer sarama.AsyncProducer, err error) {
	configProducer := sarama.NewConfig()
	configProducer.Producer.RequiredAcks = sarama.WaitForAll
	configProducer.Producer.Partitioner = sarama.NewHashPartitioner
	configProducer.Producer.Return.Successes = true
	configProducer.Producer.Return.Errors = true
	configProducer.Version = sarama.V0_11_0_2

	if conf.SaslEnable {
		configProducer.Net.SASL.Enable = true
		configProducer.Net.SASL.User = conf.User
		configProducer.Net.SASL.Password = conf.Password
	}

	producer, err = sarama.NewAsyncProducer(conf.Broker, configProducer)
	if err != nil {
		return
	}

	return
}

// InitSyncProducer 同步生产者.
func InitSyncProducer(conf *ProducerConfig) (producer sarama.SyncProducer, err error) {
	configProducer := sarama.NewConfig()
	configProducer.Producer.RequiredAcks = sarama.WaitForAll
	configProducer.Producer.Partitioner = sarama.NewHashPartitioner
	configProducer.Producer.Return.Successes = true
	configProducer.Producer.Return.Errors = true
	configProducer.Version = sarama.V0_11_0_2

	if conf.SaslEnable {
		configProducer.Net.SASL.Enable = true
		configProducer.Net.SASL.User = conf.User
		configProducer.Net.SASL.Password = conf.Password
	}

	producer, err = sarama.NewSyncProducer(conf.Broker, configProducer)
	if err != nil {
		return
	}

	return
}

type ConsumerGroupConfig struct {
	Topic      []string
	Broker     []string
	SaslEnable bool
	User       string
	Password   string
	Group      string
}

// InitCustomerGroup 消费者组.
func InitCustomerGroup(conf *ConsumerGroupConfig) (consumeGroup sarama.ConsumerGroup, err error) {
	kafkaConf := sarama.NewConfig()
	kafkaConf.Version = sarama.V0_11_0_2
	kafkaConf.Consumer.Return.Errors = true
	kafkaConf.Consumer.Offsets.Retry.Max = 3
	kafkaConf.Consumer.Offsets.Initial = sarama.OffsetNewest
	if conf.SaslEnable {
		kafkaConf.Net.SASL.Enable = true
		kafkaConf.Net.SASL.User = conf.User
		kafkaConf.Net.SASL.Password = conf.Password
	}
	cli, err := sarama.NewClient(conf.Broker, kafkaConf)
	if err != nil {
		return
	}
	consumeGroup, err = sarama.NewConsumerGroupFromClient(conf.Group, cli)
	if err != nil {
		return
	}
	return
}
