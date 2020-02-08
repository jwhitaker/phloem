package main

import (
	"log"
	"encoding/json"
	"time"

	"net/http"

	"github.com/gorilla/mux"
	"github.com/jwhitaker/phloem"
	"github.com/jwhitaker/phloem/internal/recipe"
	"github.com/google/uuid"
)

func getNow() int64 {
	now := time.Now()
	nano := now.UnixNano()

	return nano / 1000
}

type RecipeController struct {
	producer phloem.Producer
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

	w.Header().Add("Location", "/recipe/" + rec.Id)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(rec)

	log.Println("Saved!")
}

func main() {
	log.Println("Starting server")

	producer := phloem.NewKafkaProducer()

	recipeController := RecipeController { producer }

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/recipe", recipeController.SaveRecipe).
		Methods("POST")
	
	log.Fatal(http.ListenAndServe(":8080", router))
}
