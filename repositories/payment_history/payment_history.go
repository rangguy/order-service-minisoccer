package repositories

import (
	"context"
	"gorm.io/gorm"
	errWrap "payment-service/common/error"
	errConstant "payment-service/constants/error"
	"payment-service/domain/dto"
	"payment-service/domain/models"
)

type PaymentHistoryRepository struct {
	db *gorm.DB
}

type IPaymentHistoryRepository interface {
	Create(context.Context, *gorm.DB, *dto.PaymentHistoryRequest) error
}

func NewPaymentHistoryRepository(db *gorm.DB) IPaymentHistoryRepository {
	return &PaymentHistoryRepository{db: db}
}

func (p *PaymentHistoryRepository) Create(ctx context.Context, tx *gorm.DB, request *dto.PaymentHistoryRequest) error {
	paymentHistory := models.PaymentHistory{
		PaymentID: request.PaymentID,
		Status:    request.Status,
	}

	err := tx.
		WithContext(ctx).
		Create(&paymentHistory).
		Error
	if err != nil {
		return errWrap.WrapError(errConstant.ErrSQLError)
	}

	return nil
}
