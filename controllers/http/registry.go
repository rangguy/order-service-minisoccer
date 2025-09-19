package controllers

import (
	controllers "order-service/controllers/http/order"
	"order-service/services"
)

type Registry struct {
	service services.IServiceRegistry
}

type IControllerRegistry interface {
	GetOrder() controllers.IOrderController
}

func NewControllerRegistry(service services.IServiceRegistry) IControllerRegistry {
	return &Registry{service: service}
}

func (r *Registry) GetOrder() controllers.IOrderController {
	return controllers.NewOrderController(r.service)
}
