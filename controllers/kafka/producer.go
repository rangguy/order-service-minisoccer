package kafka

import (
	"github.com/IBM/sarama"
	"github.com/sirupsen/logrus"
	configApp "payment-service/config"
)

type Kafka struct {
	brokers []string
}

type IKafka interface {
	ProducerMessage(string, []byte) error
}

func NewKafkaProducer(brokers []string) *Kafka {
	return &Kafka{
		brokers: brokers,
	}
}

func (k *Kafka) ProducerMessage(topic string, data []byte) error {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = configApp.Config.Kafka.MaxRetry
	producer, err := sarama.NewSyncProducer(k.brokers, config)
	if err != nil {
		logrus.Errorf("failed create producer: %v", err)
		return err
	}
	defer func(producer sarama.SyncProducer) {
		err = producer.Close()
		if err != nil {
			logrus.Errorf("failed close producer: %v", err)
			return
		}
	}(producer)

	message := &sarama.ProducerMessage{
		Topic:   topic,
		Headers: nil,
		Value:   sarama.ByteEncoder(data),
	}

	partition, offset, err := producer.SendMessage(message)
	if err != nil {
		logrus.Errorf("failed send message: %v", err)
		return err
	}

	logrus.Infof("message is stored in topic(%s)/partition(%d)/offset(%d)\n", topic, partition, offset)
	return nil
}
