package phloem

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"encoding/json"
	"log"
)

type KafkaProducer struct {
	producer *kafka.Producer
}

func NewKafkaProducer() KafkaProducer {
	producer, err := kafka.NewProducer(&kafka.ConfigMap {
		"bootstrap.servers": "localhost",
	})

	if err != nil {
		log.Fatal("Failed to create producer: %s\n", err)
	}

	return KafkaProducer{ producer }
}

func (kafkaProducer KafkaProducer) Send(event Event) {
	log.Printf("Sending %+v\n", event)

	deliveryChan := make(chan kafka.Event)
	defer close(deliveryChan)

	rawPayload, err := json.Marshal(event)

	log.Printf("%s", rawPayload)

	if err != nil {
		log.Println("Failed to incode message")
		return
	}

	_ = kafkaProducer.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic: &event.Event.Aggregate,
			Partition: kafka.PartitionAny,
		},
		Value: rawPayload,
		Headers: []kafka.Header{{Key: "myTestHeader", Value: []byte("header values are binary")}},
	}, deliveryChan)

	e := <- deliveryChan
	m := e.(*kafka.Message)

	log.Printf("%+v\n", m)
}
