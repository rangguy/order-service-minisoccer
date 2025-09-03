package repositories

import (
	"gorm.io/gorm"
	paymentRepository "payment-service/repositories/payment"
	paymentHistoryRepository "payment-service/repositories/payment_history"
)

type Registry struct {
	db *gorm.DB
}

type IRepositoryRegistry interface {
	GetPayment() paymentRepository.IPaymentRepository
	GetPaymentHistory() paymentHistoryRepository.IPaymentHistoryRepository
	GetTx() *gorm.DB
}

func NewRepositoryRegistry(db *gorm.DB) IRepositoryRegistry {
	return &Registry{
		db: db,
	}
}

func (r *Registry) GetPayment() paymentRepository.IPaymentRepository {
	return paymentRepository.NewPaymentRepository(r.db)
}

func (r *Registry) GetPaymentHistory() paymentHistoryRepository.IPaymentHistoryRepository {
	return paymentHistoryRepository.NewPaymentHistoryRepository(r.db)
}

func (r *Registry) GetTx() *gorm.DB {
	return r.db
}
