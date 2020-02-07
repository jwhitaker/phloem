package phloem

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"os"
)

type KafkaConsumer struct {
	consumer *kafka.Consumer
}

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

func (kafkaConsumer KafkaConsumer) Poll() *Event {
	timeoutMs := 1000

	ev := kafkaConsumer.consumer.Poll(timeoutMs)

	if ev == nil {
		return nil
	}

	switch e := ev.(type) {
	case *kafka.Message:
		e.String()
		log.Println("Here is a message")

		return &Event {}

	case kafka.Error:
		fmt.Fprintf(os.Stderr, "%% Error: %v: %v\n", e.Code(), e)

	default:
		fmt.Printf("Ignored %v\n", e)
	}

	return nil
}

func (kafkaConsumer KafkaConsumer) Close() {
	kafkaConsumer.consumer.Close()
}