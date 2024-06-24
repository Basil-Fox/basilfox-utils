package kafka

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/FiberApps/common-library/logger"
	"github.com/Shopify/sarama"
)

// Consumer
func createConsumer(brokersUrl []string) (sarama.Consumer, error) {
	config := sarama.NewConfig()

	// Additional Config
	config.Consumer.Return.Errors = true

	consumer, err := sarama.NewConsumer(brokersUrl, config)
	if err != nil {
		return nil, err
	}
	return consumer, nil
}

// Producer
func createProducer(brokersUrl []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()

	// Additional Config
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	producer, err := sarama.NewSyncProducer(brokersUrl, config)
	if err != nil {
		return nil, err
	}
	return producer, nil
}

// Publisher
func PublishMessage(brokersUrl []string, topic string, message []byte) error {
	log := logger.New()
	producer, err := createProducer(brokersUrl)
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

	log.Info("Message published on topic(%s)/partition(%d)/offset(%d)", topic, partition, offset)
	return nil
}

// Add worker
func AddWorker(brokersUrl []string, topic string, handler KafkaWorker) {
	log := logger.New()
	logPrefix := "KAFKA_WORKER"

	worker, err := createConsumer(brokersUrl)
	if err != nil {
		log.Error("%s:: Error creating consumer: %v", logPrefix, err)
	}
	// calling ConsumePartition. It will open one connection per broker
	// and share it for all partitions that live on it.
	consumer, err := worker.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Error("%s:: Error while consuming partition: %v", logPrefix, err)
	}

	log.Info("%s:: Consumer started listening on topic(%s)", logPrefix, topic)

	doneCh := make(chan struct{})
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for {
			select {
			// Handle kafka errors
			case err := <-consumer.Errors():
				log.Error("%s:: Consumer error: %v", logPrefix, err)

			// Handle new message from kafka
			case msg := <-consumer.Messages():
				log.Info("%s:: Message received on topic(%s)", logPrefix, string(msg.Topic))
				if err := handler(msg); err != nil {
					log.Error("%s:: Error while consuming message: %v", logPrefix, err)
					continue
				}

			// Handle termination signals
			case <-sigChan:
				log.Info("%s:: Interrupt detected", logPrefix)
				return
			}
		}
	}()

	<-doneCh

	if err := consumer.Close(); err != nil {
		log.Error("%s:: Error while destroying consumer: %v", logPrefix, err)
	}

	if err := worker.Close(); err != nil {
		log.Error("%s:: Error while destroying worker: %v", logPrefix, err)
	}
}
