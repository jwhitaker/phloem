package apiservice

import (
	"github.com/jwhitaker/phloem/internal/pkg/phloem"
	"github.com/jwhitaker/phloem/internal/pkg/recipe"
	"log"
)

// ApiService defines methods for the api service
type ApiService struct {
}

// NewApiService constructs a new api service instance
func NewApiService() ApiService {
	return ApiService{}
}

// RecipeCreated handles when a recipe has been created
func (apiService ApiService) RecipeCreated(event *phloem.Event) {
	log.Printf("save recipe created")

	var rec recipe.Recipe

	err := event.GetPayload(&rec)

	if err != nil {
		log.Printf("Failed to get payload", err)
	}

	log.Printf("%s", rec)
}

// RecipeUpdated handles when a recipe has been updated
func (apiService ApiService) RecipeUpdated(event *phloem.Event) {
	log.Printf("Updating recipe: %+v\n", event)
}
