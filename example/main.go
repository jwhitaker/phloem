package main

import (
	"log"
	"time"

	"github.com/jwhitaker/phloem"
)

func main() {
	log.Println("Hello World")

	eventBus := phloem.NewEventBus()
	consumer := phloem.NewMemoryConsumer(&eventBus)
	producer := phloem.NewMemoryProducer(&eventBus)

	consumer.Subscribe("topic_1", func(event *phloem.Event) {
		log.Printf("Handler #1 : %s\n", event)
	})

	go func() {
		time.Sleep(5 * time.Second)

		producer.Send("topic_1", &phloem.Event { "Hello World" })
	}()

	phloem.Listen(&consumer)
	// eventBus := NewEventBus()

	// eventBus.Subscribe("topic_1", func(event *Event) {
	// 	log.Printf("Handler #2 : %s\n", event.Payload)

	// 	eventBus.Publish("topic_2", &Event{"Firing another event"})
	// })

	// eventBus.Subscribe("topic_2", func(event *Event) {
	// 	log.Printf("Handler #3 : %s\n", event.Payload)

	// 	eventBus.Publish("topic_3", &Event{"This one should get lost"})
	// })

	// eventBus.Publish("topic_1", &Event{"Here is some payload"})

	// for {
	// }
}
