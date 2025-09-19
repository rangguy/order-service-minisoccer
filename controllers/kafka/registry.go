package kafka

import (
	kafka "order-service/controllers/kafka/payment"
	"order-service/services"
)

type Registry struct {
	service services.IServiceRegistry
}

type IKafkaRegistry interface {
	GetPayment() kafka.IPaymentKafka
}

func NewKafkaRegistry(service services.IServiceRegistry) *Registry {
	return &Registry{service: service}
}

func (r *Registry) GetPayment() kafka.IPaymentKafka {
	return kafka.NewPaymentKafka(r.service)
}
