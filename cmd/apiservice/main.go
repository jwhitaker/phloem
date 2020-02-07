package main

import (
	"log"

	"github.com/jwhitaker/phloem"
)

func recipeCreated(event *phloem.Event) {

}

func recipeUpdated(event *phloem.Event) {

}

func main() {
	consumer := phloem.NewKafkaConsumer()

	router := phloem.NewEventRouter()
	router.AddHandler("recipeCreated", "recipe", recipeCreated)
	router.AddHandler("recipeUpdated", "recipe", recipeUpdated)

	phloem.ListenAndRoute(consumer, router)

	log.Println("Starting apiservice")
}
