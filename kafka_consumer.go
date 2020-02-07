package phloem

import (
	"fmt"
	"encoding/json"
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
func (kafkaConsumer KafkaConsumer) Subscribe(eventIds []EventIdentifier) {
	// Get the event ids
	// create a map with all the values as key
	uniqMap := make(map[string]struct{})

	for _, v := range eventIds {
		uniqMap[v.Aggregate] = struct{}{}
	}

	topics := make([]string, 0, len(uniqMap))

	for v := range uniqMap {
		topics = append(topics, v)
	}

	kafkaConsumer.consumer.SubscribeTopics(topics, nil)

	log.Printf("Consuming %+v", topics)
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
		log.Printf("%+s\n", e.Value)
		log.Println("Here is a message")

		var event Event

		json.Unmarshal(e.Value, &event)

		log.Printf("%s\n", event)

		return &event

	case kafka.Error:
		fmt.Fprintf(os.Stderr, "%% Error: %v: %v\n", e.Code(), e)

	default:
		fmt.Printf("Ignored %v\n", e)
	}

	return nil
}

// Close close the consumer connection to kafka.
func (kafkaConsumer KafkaConsumer) Close() {
	kafkaConsumer.consumer.Close()
}