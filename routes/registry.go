package routes

import (
	"github.com/gin-gonic/gin"
	"payment-service/clients"
	controllers "payment-service/controllers/http"
	routes "payment-service/routes/payment"
)

type Registry struct {
	controller controllers.IControllerRegistry
	group      *gin.RouterGroup
	client     clients.IClientRegistry
}

type IRouteRegistry interface {
	Serve()
}

func NewRouteRegistry(controller controllers.IControllerRegistry, group *gin.RouterGroup, client clients.IClientRegistry) IRouteRegistry {
	return &Registry{
		controller: controller,
		group:      group,
		client:     client,
	}
}

func (r *Registry) Serve() {
	r.paymentRoute().Run()
}

func (r *Registry) paymentRoute() routes.IPaymentRoute {
	return routes.NewPaymentRoute(r.group, r.controller, r.client)
}
