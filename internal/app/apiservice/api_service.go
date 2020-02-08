package apiservice

import (
	"log"

	"github.com/jwhitaker/phloem/internal/pkg/phloem"
	"github.com/jwhitaker/phloem/internal/pkg/recipe"
)

// APIService defines methods for the api service
type APIService struct {
}

// NewApiService constructs a new api service instance
func NewApiService() APIService {
	return APIService{}
}

// RecipeCreated handles when a recipe has been created
func (apiService APIService) RecipeCreated(event *phloem.Event) {
	log.Printf("save recipe created")

	var rec recipe.Recipe

	err := event.GetPayload(&rec)

	if err != nil {
		log.Printf("Failed to get payload %s", err.Error())
	}

	log.Printf("%+v", rec)
}

// RecipeUpdated handles when a recipe has been updated
func (apiService APIService) RecipeUpdated(event *phloem.Event) {
	log.Printf("Updating recipe: %+v\n", event)
}
