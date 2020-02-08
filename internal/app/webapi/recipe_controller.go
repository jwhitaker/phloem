package webapi

import (
	"encoding/json"
	"log"
	"time"

	"net/http"

	"github.com/google/uuid"
	"github.com/jwhitaker/phloem/internal/pkg/phloem"
	"github.com/jwhitaker/phloem/internal/pkg/recipe"
)

func getNow() int64 {
	now := time.Now()
	nano := now.UnixNano()

	return nano / 1000
}

type RecipeController struct {
	producer phloem.Producer
}

func NewRecipeController(producer phloem.Producer) RecipeController {
	return RecipeController{producer}
}

func (recipeController RecipeController) SaveRecipe(w http.ResponseWriter, r *http.Request) {
	var rec recipe.Recipe

	if err := json.NewDecoder(r.Body).Decode(&rec); err != nil {
		log.Printf("Could not parse payload: %s", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	rec.Id = uuid.New().String()
	rec.DateCreated = getNow()
	rec.DateModified = getNow()

	recipeController.producer.Send(recipe.RECIPE_CREATED, rec)

	w.Header().Add("Location", "/recipe/"+rec.Id)
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(rec); err != nil {
		log.Printf("Could not encode response: #{err.Error()}")
	}
}
