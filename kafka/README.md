# Kafka

## 生产者&消费者
- 新建Kafka客户端
- 同步生产者
- 异步生产者
- 消费者组
```go
// NewClient 客户端.
func NewClient(ctx context.Context, conf *Kafka, kafkaSyncProducer sarama.SyncProducer,
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
```

## Kafka Service
- 同步发送
- 异步发送
- 消费消息
- 订阅者重写
```go
// AsyncSend 异步发送.
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

	return
}

// SyncSend 同步发送.
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

// Consumer 重写订阅者，并重写订阅者的所有方法.
type Consumer struct {
	callback func(message *sarama.ConsumerMessage) (err error)
	log      *log.Log
	w        *sync.WaitGroup
}

// Setup 方法在新会话开始时运行的，然后才使用声明.
func (consumer *Consumer) Setup(sess sarama.ConsumerGroupSession) error {
	consumer.w = &sync.WaitGroup{}

	return nil
}

// Cleanup 一旦所有的订阅者协程都退出，Cleaup方法将在会话结束时运行.
func (consumer *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	zap.S().Info("Cleanup")
	consumer.w.Wait()
	zap.S().Info("Cleanup end")

	return nil
}

// ConsumeClaim 订阅者在会话中消费消息，并标记当前消息已经被消费.
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

		//有错误，不确认消息
		if err := consumer.callback(message); err != nil {
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

	if err = consumerGroup.Consume(ctx, kc.conf.Consume.Topics, consume); err != nil {
		zap.S().Errorf("Kafka consumerGroup.Consume fail, err: %v", err)

		return
	}

	return
}
```