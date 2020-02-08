package recipe

import "github.com/jwhitaker/phloem/internal/pkg/phloem"

const (
	RECIPE_CREATED = "recipeCreated"
	RECIPE_UPDATED = "recipeUpdated"
)

func NewRecipeCreatedEvent(recipe Recipe) (*phloem.Event, error) {
	return phloem.NewEvent(RECIPE_CREATED, recipe)
}
