package kafka

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/FiberApps/core/logger"
	"github.com/Shopify/sarama"
)

type KafkaWorker func(*sarama.ConsumerMessage) error

func AddWorker(broker string, topic string, handler KafkaWorker) {
	log := logger.New()
	worker, err := createConsumer([]string{broker})
	if err != nil {
		log.Error("KAFKA_WORKER_ERR -> CREATING_CONSUMER::", err)
	}
	// calling ConsumePartition. It will open one connection per broker
	// and share it for all partitions that live on it.
	consumer, err := worker.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Error("KAFKA_WORKER_ERR -> CONSUMING_PARTITION::", err)
	}
	log.Info("KAFKA_WORKER::Consumer started listening on topic:", topic)

	doneCh := make(chan struct{})
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for {
			select {
			// Handle kafka errors
			case err := <-consumer.Errors():
				log.Error("KAFKA_WORKER_ERR -> CONSUMER_ERR::", err)

			// Handle new message from kafka
			case msg := <-consumer.Messages():
				log.Info(fmt.Sprintf("KAFKA_WORKER::Received: | Topic(%s) | Message(%s)", string(msg.Topic), string(msg.Value)))
				if err := handler(msg); err != nil {
					log.Error("KAFKA_WORKER_ERR -> HANDLING_INCOMING_MESSAGE::", err)
					continue
				}

			// Handle termination signals
			case <-sigChan:
				log.Info("KAFKA_WORKER::Interrupt is detected")
				return
			}
		}
	}()

	<-doneCh

	if err := consumer.Close(); err != nil {
		log.Error("KAFKA_WORKER_ERR -> CLOSING_CONSUMER::", err)
	}

	if err := worker.Close(); err != nil {
		log.Error("KAFKA_WORKER_ERR -> CLOSING_WORKER::", err)
	}
}
