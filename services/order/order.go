package services

import (
	"context"
	"order-service/clients"
	"order-service/common/util"
	"order-service/domain/dto"
	"order-service/repositories"
)

type OrderService struct {
	repository repositories.IRepositoryRegistry
	client     clients.IClientRegistry
}

type IOrderService interface {
	GetAllWithPagination(context.Context, *dto.OrderRequestParam) (*util.PaginationResult, error)
	GetByID(context.Context, string) (*dto.OrderResponse, error)
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

func (o *OrderService) GetByID(ctx context.Context, s string) (*dto.OrderResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (o *OrderService) GetOrderByUserID(ctx context.Context) ([]dto.OrderByUserIDResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (o *OrderService) Create(ctx context.Context, request *dto.OrderRequest) (*dto.OrderResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (o *OrderService) HandlePayment(ctx context.Context, data *dto.PaymentData) error {
	//TODO implement me
	panic("implement me")
}
