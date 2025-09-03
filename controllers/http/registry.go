package controllers

import (
	controllers "payment-service/controllers/http/payment"
	"payment-service/services"
)

type Registry struct {
	service services.IServiceRegistry
}

type IControllerRegistry interface {
	GetPayment() controllers.IPaymentController
}

func NewControllerRegistry(service services.IServiceRegistry) IControllerRegistry {
	return &Registry{service: service}
}

func (r *Registry) GetPayment() controllers.IPaymentController {
	return controllers.NewPaymentController(r.service)
}
