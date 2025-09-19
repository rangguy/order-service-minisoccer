package kafka

import (
	"golang.org/x/exp/slices"
	"order-service/config"
	"order-service/controllers/kafka"
	kafka2 "order-service/controllers/kafka/payment"
)

type Kafka struct {
	consumer *ConsumerGroup
	kafka    kafka.IKafkaRegistry
}

type IKafka interface {
	Register()
}

func NewKafkaConsumer(consumer *ConsumerGroup, kafka kafka.IKafkaRegistry) IKafka {
	return &Kafka{
		consumer: consumer,
		kafka:    kafka,
	}
}

func (k *Kafka) Register() {
	k.paymentHandler()
}

func (k *Kafka) paymentHandler() {
	if slices.Contains(config.Config.Kafka.Topics, kafka2.PaymentTopic) {
		k.consumer.RegisterHandler(kafka2.PaymentTopic, k.kafka.GetPayment().HandlePayment)
	}
}
