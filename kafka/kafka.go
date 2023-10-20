package kafka

import (
	"fmt"
	"os"
	"strings"

	"github.com/FiberApps/core/logger"
	"github.com/Shopify/sarama"
)

func getBaseConfig() ([]string, *sarama.Config) {
	brokersList := os.Getenv("KAFKA_BROKERS")
	brokersUrl := strings.Split(brokersList, ",")

	// Base config
	config := sarama.NewConfig()

	return brokersUrl, config
}

// Consumer
func createConsumer() (sarama.Consumer, error) {
	brokers, config := getBaseConfig()

	// Additional Config
	config.Consumer.Return.Errors = true

	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		return nil, err
	}
	return consumer, nil
}

// Producer
func createProducer() (sarama.SyncProducer, error) {
	brokers, config := getBaseConfig()

	// Additional Config
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}
	return producer, nil
}

// Publisher
func PublishMessage(topic string, message []byte) error {
	log := logger.New()
	producer, err := createProducer()
	if err != nil {
		return err
	}
	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		return err
	}
	log.Info(fmt.Sprintf("Message is stored in topic(%s)/partition(%d)/offset(%d)", topic, partition, offset))
	return nil
}
