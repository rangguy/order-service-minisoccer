package routes

import (
	"github.com/gin-gonic/gin"
	"order-service/clients"
	"order-service/constants"
	"order-service/controllers/http"
	"order-service/middlewares"
)

type OrderRoute struct {
	controller controllers.IControllerRegistry
	client     clients.IClientRegistry
	group      *gin.RouterGroup
}

type IOrderRoute interface {
	Run()
}

func NewOrderRoute(group *gin.RouterGroup, controller controllers.IControllerRegistry, client clients.IClientRegistry) IOrderRoute {
	return &OrderRoute{
		group:      group,
		controller: controller,
		client:     client,
	}
}

func (o *OrderRoute) Run() {
	group := o.group.Group("/order")
	group.Use(middlewares.Authenticate())
	group.GET("/", middlewares.CheckRole([]string{
		constants.Admin,
		constants.Customer,
	}, o.client), o.controller.GetOrder().GetAllWithPagination)
	group.GET("/:uuid", middlewares.CheckRole([]string{
		constants.Admin,
		constants.Customer,
	}, o.client), o.controller.GetOrder().GetByUUID)
	group.GET("/user", middlewares.CheckRole([]string{
		constants.Customer,
	}, o.client), o.controller.GetOrder().GetOrderByUserID)
	group.POST("/user", middlewares.CheckRole([]string{
		constants.Customer,
	}, o.client), o.controller.GetOrder().Create)
}
