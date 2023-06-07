package kafka

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Shopify/sarama"
)

type KafkaWorker func(*sarama.ConsumerMessage) error

func AddWorker(broker string, topic string, handler KafkaWorker) {
	worker, err := createConsumer([]string{broker})
	if err != nil {
		log.Printf("[KafkaWorkerErr]::%s", err)
	}
	// calling ConsumePartition. It will open one connection per broker
	// and share it for all partitions that live on it.
	consumer, err := worker.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Printf("[KafkaWorkerErr]::%s", err)
	}
	log.Printf("[KafkaWorker]::Consumer started listening on topic:%s", topic)

	doneCh := make(chan struct{})
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for {
			select {
			// Handle kafka errors
			case err := <-consumer.Errors():
				log.Printf("[KafkaWorkerErr]::%s", err)

			// Handle new message from kafka
			case msg := <-consumer.Messages():
				log.Printf("[KafkaWorker]::Received: | Topic(%s) | Message(%s) \n", string(msg.Topic), string(msg.Value))
				if err := handler(msg); err != nil {
					log.Printf("[KafkaWorkerErr]::%s", err)
					continue
				}

			// Handle termination signals
			case <-sigChan:
				log.Println("[KafkaWorker]::Interrupt is detected")
				return
			}
		}
	}()

	<-doneCh

	if err := consumer.Close(); err != nil {
		log.Printf("[KafkaWorkerErr]::%s", err)
	}

	if err := worker.Close(); err != nil {
		log.Printf("[KafkaWorkerErr]::%s", err)
	}
}
