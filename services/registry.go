package services

import (
	"order-service/clients"
	"order-service/controllers/kafka"
	"order-service/repositories"
	services "order-service/services/order"
)

type Registry struct {
	repository repositories.IRepositoryRegistry
	client     clients.IClientRegistry
	kafka      kafka.IKafkaRegistry
}

type IServiceRegistry interface {
	GetOrder() services.IOrderService
}

func NewServiceRegistry(repository repositories.IRepositoryRegistry, client clients.IClientRegistry) IServiceRegistry {
	return &Registry{
		repository: repository,
		client:     client,
	}
}

func (r *Registry) GetOrder() services.IOrderService {
	return services.NewOrderService(r.repository, r.client)
}
