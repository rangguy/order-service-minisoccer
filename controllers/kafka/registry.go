package kafka

import (
	"order-service/controllers/kafka/payment"
	"order-service/services"
)

type Registry struct {
	service services.IServiceRegistry
}

type IKafkaRegistry interface {
	GetPaymentKafka() kafka.IPaymentKafka
}

func NewKafkaRegistry(service services.IServiceRegistry) *Registry {
	return &Registry{service: service}
}

func (r *Registry) GetPaymentKafka() kafka.IPaymentKafka {
	return kafka.NewPaymentKafka(r.service)
}
