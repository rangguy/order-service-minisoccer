package services

import (
	"context"
	"fmt"
	"order-service/clients"
	clientUser "order-service/clients/user"
	"order-service/common/util"
	"order-service/constants"
	"order-service/domain/dto"
	"order-service/domain/models"
	"order-service/repositories"
)

type OrderService struct {
	repository repositories.IRepositoryRegistry
	client     clients.IClientRegistry
}

type IOrderService interface {
	GetAllWithPagination(context.Context, *dto.OrderRequestParam) (*util.PaginationResult, error)
	GetByUUID(context.Context, string) (*dto.OrderResponse, error)
	GetOrderByUserID(context.Context) ([]dto.OrderByUserIDResponse, error)
	Create(context.Context, *dto.OrderRequest) (*dto.OrderResponse, error)
	HandlePayment(context.Context, *dto.PaymentData) error
}

func NewOrderService(repository repositories.IRepositoryRegistry, client clients.IClientRegistry) *OrderService {
	return &OrderService{repository: repository, client: client}
}

func (o *OrderService) GetAllWithPagination(ctx context.Context, param *dto.OrderRequestParam) (*util.PaginationResult, error) {
	orders, total, err := o.repository.GetOrder().FindAllWithPagination(ctx, param)
	if err != nil {
		return nil, err
	}

	orderResults := make([]*dto.OrderResponse, 0, len(orders))
	for _, order := range orders {
		user, err := o.client.GetUser().GetUserByUUID(ctx, order.UserID)
		if err != nil {
			return nil, err
		}
		orderResults = append(orderResults, &dto.OrderResponse{
			UUID:      order.UUID,
			Code:      order.Code,
			Username:  user.Name,
			Amount:    order.Amount,
			Status:    order.Status.GetStatusString(),
			OrderDate: order.Date,
			CreatedAt: *order.CreatedAt,
			UpdatedAt: *order.UpdatedAt,
		})
	}

	paginationParam := util.PaginationParam{
		Page:  param.Page,
		Limit: param.Limit,
		Count: total,
		Data:  orderResults,
	}

	response := util.GeneratePagination(paginationParam)
	return &response, nil
}

func (o *OrderService) GetByUUID(ctx context.Context, uuid string) (*dto.OrderResponse, error) {
	var (
		order *models.Order
		user  *clientUser.UserData
		err   error
	)

	order, err = o.repository.GetOrder().FindByUUID(ctx, uuid)
	if err != nil {
		return nil, err
	}

	user, err = o.client.GetUser().GetUserByUUID(ctx, order.UserID)
	if err != nil {
		return nil, err
	}

	response := dto.OrderResponse{
		UUID:      order.UUID,
		Code:      order.Code,
		Username:  user.Name,
		Amount:    order.Amount,
		Status:    order.Status.GetStatusString(),
		OrderDate: order.Date,
		CreatedAt: *order.CreatedAt,
		UpdatedAt: *order.UpdatedAt,
	}

	return &response, nil
}

func (o *OrderService) GetOrderByUserID(ctx context.Context) ([]dto.OrderByUserIDResponse, error) {
	var (
		order []models.Order
		user  = ctx.Value(constants.User).(*clientUser.UserData)
		err   error
	)

	order, err = o.repository.GetOrder().FindByUserID(ctx, user.UUID.String())
	if err != nil {
		return nil, err
	}

	orderList := make([]dto.OrderByUserIDResponse, 0, len(order))
	for _, item := range order {
		payment, err := o.client.GetPayment().GetPaymentByUUID(ctx, item.PaymentID)
		if err != nil {
			return nil, err
		}
		orderList = append(orderList, dto.OrderByUserIDResponse{
			Code:        item.Code,
			Amount:      fmt.Sprintf("%s", util.RupiahFormat(&item.Amount)),
			Status:      item.Status.GetStatusString(),
			OrderDate:   item.Date.String(),
			PaymentLink: payment.PaymentLink,
			InvoiceLink: payment.InvoiceLink,
		})
	}

	return orderList, nil
}

func (o *OrderService) Create(ctx context.Context, request *dto.OrderRequest) (*dto.OrderResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (o *OrderService) HandlePayment(ctx context.Context, data *dto.PaymentData) error {
	//TODO implement me
	panic("implement me")
}
