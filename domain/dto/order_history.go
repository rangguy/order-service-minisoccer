package dto

import "order-service/constants"

type OrderHistory struct {
	OrderID uint
	Status  constants.OrderStatusString
}
