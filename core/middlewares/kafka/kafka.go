package kafka

import (
	"context"
	"core/config"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"
)

type Kfk struct {
	*kafka.Client
	sharedTransport *kafka.Transport
	Mechanism       plain.Mechanism
	conf            *config.Kafka
}

func New(conf *config.Kafka) *Kfk {
	k := new(Kfk)
	k.Mechanism = plain.Mechanism{
		Username: conf.Username,
		Password: conf.Password,
	}
	k.sharedTransport = &kafka.Transport{
		SASL: k.Mechanism,
	}
	k.Client = k.newClient(conf)
	k.conf = conf
	return k
}

// 低等级api用于创建 topic或分区
func (k *Kfk) newClient(conf *config.Kafka) *kafka.Client {
	return &kafka.Client{
		Addr:      kafka.TCP(conf.Address...),
		Timeout:   10 * time.Second,
		Transport: k.sharedTransport,
	}
}

// NewReader 读 close
func (k *Kfk) NewReader(config kafka.ReaderConfig) *kafka.Reader {
	dialer := &kafka.Dialer{
		Timeout:       10 * time.Second,
		DualStack:     true,
		SASLMechanism: k.Mechanism,
	}
	config.Brokers = k.conf.Address
	config.Dialer = dialer
	return kafka.NewReader(config)
}

// NewWriter 写
func (k *Kfk) NewWriter(topic string) *kafka.Writer {
	w := &kafka.Writer{
		Addr:                   kafka.TCP(k.conf.Address...),
		Topic:                  topic,
		Balancer:               &kafka.LeastBytes{},
		Transport:              k.Transport,
		AllowAutoTopicCreation: true,
	}
	return w
}

func client() {
	mechanism := plain.Mechanism{
		Username: "adminplain",
		Password: "admin-secret",
	}

	sharedTransport := &kafka.Transport{
		SASL: mechanism,
	}

	client := &kafka.Client{
		Addr:      kafka.TCP("192.168.8.243:9093"),
		Timeout:   10 * time.Second,
		Transport: sharedTransport,
	}

	//创建新的分区要在topic不存在的情况下
	_, err := client.CreateTopics(context.Background(), &kafka.CreateTopicsRequest{})
	if err != nil {
		panic(err.Error())
	}

	//创建新分区 是要在topic 已存在的情况下
	partitionReq := kafka.CreatePartitionsRequest{
		Topics: []kafka.TopicPartitionsConfig{{Name: "topic-J", Count: 2}},
	}

	res, err := client.CreatePartitions(context.Background(), &partitionReq)
	fmt.Println(err)

	fmt.Println(res.Errors)
}
