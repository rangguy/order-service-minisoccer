package dto

import "order-service/constants"

type OrderHistoryRequest struct {
	OrderID uint
	Status  constants.OrderStatusString
}
