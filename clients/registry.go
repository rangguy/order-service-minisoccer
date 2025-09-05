package clients

import (
	"order-service/clients/config"
	fieldClient "order-service/clients/field"
	paymentClient "order-service/clients/payment"
	userClient "order-service/clients/user"
	configApp "order-service/config"
)

type ClientRegistry struct{}

type IClientRegistry interface {
	GetUser() userClient.IUserClient
	GetPayment() paymentClient.IPaymentClient
	GetField() fieldClient.IFieldClient
}

func NewClientRegistry() IClientRegistry {
	return &ClientRegistry{}
}

func (c *ClientRegistry) GetUser() userClient.IUserClient {
	return userClient.NewUserClient(
		config.NewClientConfig(
			config.WithBaseURL(configApp.Config.InternalService.User.Host),
			config.WithSignatureKey(configApp.Config.InternalService.User.SignatureKey),
		))
}

func (c *ClientRegistry) GetPayment() paymentClient.IPaymentClient {
	return paymentClient.NewPaymentClient(
		config.NewClientConfig(
			config.WithBaseURL(configApp.Config.InternalService.Payment.Host),
			config.WithSignatureKey(configApp.Config.InternalService.Payment.SignatureKey),
		))
}

func (c *ClientRegistry) GetField() fieldClient.IFieldClient {
	return fieldClient.NewFieldClient(
		config.NewClientConfig(
			config.WithBaseURL(configApp.Config.InternalService.Field.Host),
			config.WithSignatureKey(configApp.Config.InternalService.Field.SignatureKey),
		))
}
