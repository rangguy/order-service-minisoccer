package config

import "github.com/parnurzeal/gorequest"

type ClientConfig struct {
	client       *gorequest.SuperAgent
	baseURL      string
	signatureKey string
}

type IClientConfig interface {
	Client() *gorequest.SuperAgent
	BaseURL() string
	SignatureKey() string
}

type Option func(*ClientConfig)

func NewClientConfig(options ...Option) IClientConfig {
	clientConfig := &ClientConfig{
		client: gorequest.New().
			Set("Content-Type", "application/json").
			Set("Accept", "application/json"),
	}

	for _, option := range options {
		option(clientConfig)
	}

	return clientConfig
}

func (c *ClientConfig) Client() *gorequest.SuperAgent {
	return c.client
}

func (c *ClientConfig) BaseURL() string {
	return c.baseURL
}

func (c *ClientConfig) SignatureKey() string {
	return c.signatureKey
}

func WithBaseURL(baseURL string) Option {
	return func(c *ClientConfig) {
		c.baseURL = baseURL
	}
}

func WithSignatureKey(signatureKey string) Option {
	return func(c *ClientConfig) {
		c.signatureKey = signatureKey
	}
}
