package phloem

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"os"
)

// KafkaConsumer is a struct for a KafkaConsumer
type KafkaConsumer struct {
	consumer *kafka.Consumer
}

// NewKafkaConsumer returns a new kafka consumer
func NewKafkaConsumer() KafkaConsumer {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"broker.address.family": "v4",
		"group.id": "a_group",
		"session.timeout.ms": 6000,
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		log.Fatalf("Failed to create consumer: %s\n", err)
	}

	return KafkaConsumer{ consumer }
}

// Subscribe to the following events
func (kafkaConsumer KafkaConsumer) Subscribe(events []string) {
	kafkaConsumer.consumer.SubscribeTopics(events, nil)

	log.Printf("Consuming %+v", events)
}

// Poll retrieve the next available event from Kafka
func (kafkaConsumer KafkaConsumer) Poll() *Event {
	timeoutMs := 1000

	ev := kafkaConsumer.consumer.Poll(timeoutMs)

	if ev == nil {
		return nil
	}

	switch e := ev.(type) {
	case *kafka.Message:
		event := Event{
			Event: *e.TopicPartition.Topic,
			Payload: e.Value,
		}

		return &event

	case kafka.Error:
		fmt.Fprintf(os.Stderr, "%% Error: %v: %v\n", e.Code(), e)

	default:
		log.Printf("Ignored %v\n", e)
	}

	return nil
}

// Close close the consumer connection to kafka.
func (kafkaConsumer KafkaConsumer) Close() {
	kafkaConsumer.consumer.Close()
}