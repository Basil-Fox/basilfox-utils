package kafka

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/FiberApps/common-library/logger"
	"github.com/Shopify/sarama"
	"github.com/gofiber/fiber/v2"
)

type Config struct {
	BrokerUrls []string
}

var (
	kConfig  *Config
	producer sarama.SyncProducer // Global producer instance
)

// Setup Kafka Client and Producer
func SetupClient(config Config) error {
	if len(config.BrokerUrls) == 0 {
		return fmt.Errorf("kafka broker URLs must be provided")
	}

	// Initialize Kafka client config
	kConfig = &config

	// Initialize Kafka Producer
	clientConfig := sarama.NewConfig()
	clientConfig.Producer.Return.Successes = true
	clientConfig.Producer.RequiredAcks = sarama.WaitForAll
	clientConfig.Producer.Retry.Max = 5

	var err error
	producer, err = sarama.NewSyncProducer(kConfig.BrokerUrls, clientConfig)
	if err != nil {
		return fmt.Errorf("failed to initialize Kafka producer: %w", err)
	}

	return nil
}

// Publish Message to Kafka
func PublishMessage(ctx *fiber.Ctx, topic string, message []byte) error {
	log := logger.GetLogger(ctx).With().Str("kafka", "producer").Str("topic", topic).Logger()

	if producer == nil {
		return fmt.Errorf("kafka producer is not initialized")
	}

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		return err
	}

	log.Info().Int32("partition", partition).Int64("offset", offset).Msg("message published")
	return nil
}

// Create Kafka Consumer
func createConsumer() (sarama.Consumer, error) {
	if kConfig == nil {
		return nil, fmt.Errorf("kafka client isn't initialized yet")
	}

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	consumer, err := sarama.NewConsumer(kConfig.BrokerUrls, config)
	if err != nil {
		return nil, err
	}
	return consumer, nil
}

// Add Kafka Worker to Process Messages from All Partitions
func AddWorker(topic string, handler KafkaWorker) error {
	log := logger.GetLogger(nil).With().Str("kafka", "consumer").Str("topic", topic).Logger()

	consumer, err := createConsumer()
	if err != nil {
		return err
	}

	partitions, err := consumer.Partitions(topic)
	if err != nil {
		return err
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	doneCh := make(chan struct{})

	for _, partition := range partitions {
		go func(partition int32) {
			log.Info().Int32("partition", partition).Msg("starting consumer")

			partitionConsumer, err := consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
			if err != nil {
				log.Err(err).Int32("partition", partition).Msg("failed to consume from partition")
				return
			}
			defer partitionConsumer.Close()

			for {
				select {
				case err := <-partitionConsumer.Errors():
					log.Err(err).Int32("partition", partition).Msg("kafka consumer error")

				case msg := <-partitionConsumer.Messages():
					log.Info().Int32("partition", partition).Int64("offset", msg.Offset).Msg("message received")

					if err := handler(msg, &log); err != nil {
						log.Err(err).Msg("failed to process message")
					}

				case <-sigChan:
					log.Info().Int32("partition", partition).Msg("stopping consumer")
					close(doneCh)
					return
				}
			}
		}(partition)
	}

	<-doneCh // Wait for shutdown signal

	log.Info().Msg("closing Kafka consumer")
	if err := consumer.Close(); err != nil {
		log.Err(err).Msg("error closing Kafka consumer")
	}

	return nil
}
