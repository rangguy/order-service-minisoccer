package kafka

import (
	"context"
	"github.com/IBM/sarama"
	"github.com/sirupsen/logrus"
	"order-service/config"
	"time"
)

type (
	TopicName string
	Handler   func(ctx context.Context, message *sarama.ConsumerMessage) error
)

type ConsumerGroup struct {
	handler map[TopicName]Handler
}

func NewConsumerGroup() *ConsumerGroup {
	return &ConsumerGroup{handler: make(map[TopicName]Handler)}
}

func (c *ConsumerGroup) Setup(sarama sarama.ConsumerGroupSession) error {
	logrus.Infof("setup consumer group")
	return nil
}

func (c *ConsumerGroup) Cleanup(sarama sarama.ConsumerGroupSession) error {
	logrus.Infof("cleanup consumer group")
	return nil
}

func (c *ConsumerGroup) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	messages := claim.Messages()
	for message := range messages {
		handler, ok := c.handler[TopicName(message.Topic)]
		if !ok {
			logrus.Errorf("topic %s not exist", message.Topic)
			continue
		}

		var err error
		maxRetry := config.Config.Kafka.MaxRetry
		for attempt := 1; attempt <= maxRetry; attempt++ {
			err = handler(context.Background(), message)
			if err == nil {
				break
			}

			logrus.Errorf("topic %s retry %d times %v", message.Topic, attempt, err)
			if attempt == maxRetry {
				logrus.Errorf("max retry reached, message will be ignored")
			}
		}

		if err != nil {
			logrus.Errorf("error handling message on %s: %v", message.Topic, err)
			break
		}

		session.MarkMessage(message, time.Now().UTC().String())
	}
	return nil
}

func (c *ConsumerGroup) RegisterHandler(topic TopicName, handler Handler) {
	c.handler[topic] = handler
	logrus.Infof("register handler on topic %s", topic)
}
