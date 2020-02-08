package phloem

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
)

// KafkaProducer defines a Kafka version of a producer
type KafkaProducer struct {
	producer *kafka.Producer
}

// NewKafkaProducer creates a new KafkaProducer instance
func NewKafkaProducer(configuration KafkaConfiguration) KafkaProducer {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": configuration.BootstrapServers,
	})

	if err != nil {
		log.Fatalf("Failed to create producer: %s\n", err)
	}

	return KafkaProducer{producer}
}

// Send will send a event with the payload
//func (kafkaProducer KafkaProducer) Send(event string, payload interface{}) {
func (kafkaProducer KafkaProducer) Send(event Event) {
	log.Printf("Sending %+v\n", event)

	deliveryChan := make(chan kafka.Event)
	defer close(deliveryChan)

	//rawPayload, err := json.Marshal(event.Payload)

	//log.Printf("%s", rawPayload)

	//if err != nil {
	//	log.Println("Failed to encode message")
	//	return
	//}

	_ = kafkaProducer.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &event.Event,
			Partition: kafka.PartitionAny,
		},
		Value: event.Payload,
	}, deliveryChan)

	e := <-deliveryChan
	m := e.(*kafka.Message)

	log.Printf("%+v\n", m)
}
