package kafka

import (
	"context"
	"core/common/log"
	"sync"

	"github.com/Shopify/sarama"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// AsyncSend 异步发送
func (kc *Client) AsyncSend(topic string, msg string, keys ...string) {
	producer := kc.asyncProducer
	var key string
	if len(keys) > 0 {
		key = keys[0]
	}
	producer.Input() <- &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(msg),
		Key:   sarama.StringEncoder(key),
	}

	//select {
	//case res := <-producer.Successes():
	//	return res.Partition, res.Offset, nil
	//case err := <-producer.Errors():
	//	return 0, 0, errors.Wrap(err, "")
	//}
	return
}

func (kc *Client) SyncSend(topic string, msg string, keys ...string) (int32, int64, error) {
	var key string
	if len(keys) > 0 {
		key = keys[0]
	}
	partition, offset, err := kc.syncProducer.SendMessage(&sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(msg),
		Key:   sarama.StringEncoder(key),
	})

	return partition, offset, errors.Wrap(err, "")
}

// Consumer 重写订阅者，并重写订阅者的所有方法
type Consumer struct {
	callback func(message *sarama.ConsumerMessage) (err error)
	log      *log.Log
	w        *sync.WaitGroup
}

// Setup 方法在新会话开始时运行的，然后才使用声明
func (consumer *Consumer) Setup(sess sarama.ConsumerGroupSession) error {
	consumer.w = &sync.WaitGroup{}
	return nil
}

// Cleanup 一旦所有的订阅者协程都退出，Cleaup方法将在会话结束时运行
func (consumer *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	zap.S().Info("Cleanup")
	consumer.w.Wait()
	zap.S().Info("Cleanup end")
	return nil
}

// ConsumeClaim 订阅者在会话中消费消息，并标记当前消息已经被消费。
func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	consumer.w.Add(1)
	defer consumer.w.Done()
	zap.S().Infof("ConsumeClaim 开始消费, topic: %v, partition: %v ", claim.Topic(), claim.Partition())
	for message := range claim.Messages() {
		zap.S().Infof("Kafka receive msg, topic: %v, partion: %v, offset: %v, key: %v, value: %v", message.Topic, message.Partition, message.Offset, string(message.Key), string(message.Value))
		if consumer.callback == nil {
			session.MarkMessage(message, "")
			continue
		}

		err := consumer.callback(message)
		//有错误，不确认消息
		if err != nil {
			continue
		}
		session.MarkMessage(message, "")
	}

	zap.S().Infof("ConsumeClaim 退出, topic: %v, partition: %v ", claim.Topic(), claim.Partition())
	return nil
}

func (Consumer *Consumer) SetCallback(callback func(message *sarama.ConsumerMessage) (err error)) {
	Consumer.callback = callback
}

func (kc *Client) ReceiveMsg(ctx context.Context, consumerMsg func(message *sarama.ConsumerMessage) (err error)) {
	consumerGroup, err := InitCustomerGroup(&ConsumerGroupConfig{
		Broker:     kc.conf.Address,
		SaslEnable: kc.conf.SaslEnable,
		User:       kc.conf.Username,
		Password:   kc.conf.Password,
		Topic:      kc.conf.Consume.Topics,
		Group:      kc.conf.Consume.Group,
	})
	if err != nil {
		zap.S().Errorf("Kafka InitCustomerGroup fail, err: %v", err)
		return
	}
	defer func() {
		if err := recover(); err != nil {
			zap.S().Errorf("Kafka ReceiveMsg 意外退出，err: %v", err)
			return
		}
		zap.S().Infof("Kafka ReceiveMsg 退出")
	}()
	consume := new(Consumer)
	consume.SetCallback(consumerMsg)
	zap.S().Infof("Kafka ReceiveMsg 开始消费")
	err = consumerGroup.Consume(ctx, kc.conf.Consume.Topics, consume)
	if err != nil {
		zap.S().Errorf("Kafka consumerGroup.Consume fail, err: %v", err)
		return
	}
	return
}
