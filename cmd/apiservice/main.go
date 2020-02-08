package main

import (
	"log"

	"github.com/jwhitaker/phloem/internal/pkg/phloem"
	"github.com/jwhitaker/phloem/internal/pkg/recipe"
)

func recipeCreated(event *phloem.Event) {
	log.Printf("save recipe created")

	var rec recipe.Recipe

	err := event.GetPayload(&rec)

	if err != nil {
		log.Printf("Failed to get payload", err)
	}

	log.Printf("%s", rec)
}

func recipeUpdated(event *phloem.Event) {

}

func main() {
	consumer := phloem.NewKafkaConsumer()

	router := phloem.NewEventRouter()
	router.AddHandler(recipe.RECIPE_CREATED, recipeCreated)
	router.AddHandler(recipe.RECIPE_UPDATED, recipeUpdated)

	phloem.ListenAndRoute(consumer, router)

	log.Println("Starting apiservice")
}
