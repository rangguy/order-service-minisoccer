package kafka

type Registry struct {
	brokers []string
}

type IKafkaRegistry interface {
	GetKafkaProducer() IKafka
}

func NewKafkaRegistry(brokers []string) *Registry {
	return &Registry{brokers}
}

func (r *Registry) GetKafkaProducer() IKafka {
	return NewKafkaProducer(r.brokers)
}
