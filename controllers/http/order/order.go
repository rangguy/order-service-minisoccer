package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	errValidation "order-service/common/error"
	"order-service/common/response"
	"order-service/domain/dto"
	"order-service/services"
)

type OrderController struct {
	service services.IServiceRegistry
}

type IOrderController interface {
	GetAllWithPagination(*gin.Context)
	GetByUUID(*gin.Context)
	GetOrderByUserID(*gin.Context)
	Create(*gin.Context)
}

func NewOrderController(service services.IServiceRegistry) IOrderController {
	return &OrderController{service: service}
}

func (o *OrderController) GetAllWithPagination(ctx *gin.Context) {
	var params dto.OrderRequestParam
	err := ctx.ShouldBindQuery(&params)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  ctx,
		})
		return
	}

	validate := validator.New()
	if err = validate.Struct(params); err != nil {
		errMessage := http.StatusText(http.StatusUnprocessableEntity)
		errorResponse := errValidation.ErrValidationResponse(err)
		response.HttpResponse(response.ParamHTTPResp{
			Err:     err,
			Code:    http.StatusUnprocessableEntity,
			Message: &errMessage,
			Data:    errorResponse,
			Gin:     ctx,
		})
		return
	}

	result, err := o.service.GetOrder().GetAllWithPagination(ctx, &params)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  ctx,
		})
		return
	}

	response.HttpResponse(response.ParamHTTPResp{
		Code: http.StatusOK,
		Data: result,
		Gin:  ctx,
	})
}

func (o *OrderController) GetByUUID(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	result, err := o.service.GetOrder().GetByUUID(ctx, uuid)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  ctx,
		})
		return
	}

	response.HttpResponse(response.ParamHTTPResp{
		Code: http.StatusOK,
		Data: result,
		Gin:  ctx,
	})
}

func (o *OrderController) GetOrderByUserID(ctx *gin.Context) {
	result, err := o.service.GetOrder().GetOrderByUserID(ctx.Request.Context())
	if err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  ctx,
		})
		return
	}

	response.HttpResponse(response.ParamHTTPResp{
		Code: http.StatusOK,
		Data: result,
		Gin:  ctx,
	})
}

func (o *OrderController) Create(ctx *gin.Context) {
	var (
		request dto.OrderRequest
		c       = ctx.Request.Context()
	)

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  ctx,
		})
		return
	}

	validate := validator.New()
	if err = validate.Struct(request); err != nil {
		errMessage := http.StatusText(http.StatusUnprocessableEntity)
		errorResponse := errValidation.ErrValidationResponse(err)
		response.HttpResponse(response.ParamHTTPResp{
			Err:     err,
			Code:    http.StatusUnprocessableEntity,
			Message: &errMessage,
			Data:    errorResponse,
			Gin:     ctx,
		})
		return
	}

	result, err := o.service.GetOrder().Create(c, &request)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  ctx,
		})
		return
	}

	response.HttpResponse(response.ParamHTTPResp{
		Code: http.StatusCreated,
		Data: result,
		Gin:  ctx,
	})
}
