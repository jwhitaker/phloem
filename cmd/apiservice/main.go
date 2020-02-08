package main

import (
	"github.com/jwhitaker/phloem/internal/app/apiservice"
	"log"

	"github.com/jwhitaker/phloem/internal/pkg/phloem"
	"github.com/jwhitaker/phloem/internal/pkg/recipe"
)

func main() {
	consumer := phloem.NewKafkaConsumer()

	service := apiservice.NewApiService()

	router := phloem.NewEventRouter()
	router.AddHandler(recipe.RECIPE_CREATED, service.RecipeCreated)
	router.AddHandler(recipe.RECIPE_UPDATED, service.RecipeUpdated)

	phloem.ListenAndRoute(consumer, router)

	log.Println("Starting apiservice")
}
