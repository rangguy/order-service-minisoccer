package routes

import (
	"github.com/gin-gonic/gin"
	"order-service/clients"
	controllers "order-service/controllers/http"
	routes "order-service/routes/order"
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
	r.orderRoute().Run()
}

func (r *Registry) orderRoute() routes.IOrderRoute {
	return routes.NewOrderRoute(r.group, r.controller, r.client)
}
