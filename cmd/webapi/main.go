package main

import (
	"log"
	"encoding/json"
	"fmt"
	"time"

	"net/http"

	"github.com/gorilla/mux"
	"github.com/jwhitaker/phloem"
	"github.com/google/uuid"
)

func getNow() int64 {
	now := time.Now()
	nano := now.UnixNano()

	return nano / 1000
}

type RecipeRepository struct {
}

func (RecipeRepository) SaveRecipe(recipe Recipe) {
	log.Println("Saving recipe...")
}

type Recipe struct {
	Id string
	Name string
	DateCreated int64
	DateModified int64
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Done!")
}

type RecipeController struct {
	producer phloem.Producer
}

func (recipeController RecipeController) SaveRecipe(w http.ResponseWriter, r *http.Request) {
	var recipe Recipe

	if err := json.NewDecoder(r.Body).Decode(&recipe); err != nil {
		log.Printf("Could not parse payload: %s", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	recipe.Id = uuid.New().String()
	recipe.DateCreated = getNow()
	recipe.DateModified = getNow()

	recipeController.producer.Send(phloem.Event{ phloem.EventIdentifier { "recipeCreated", "recipe" }, recipe })

	w.Header().Add("Location", "/recipe/" + recipe.Id)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(recipe)

	log.Println("Saved!")
}

func main() {
	log.Println("Starting server")

	producer := phloem.NewKafkaProducer()

	recipeController := RecipeController { producer }

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/recipe", recipeController.SaveRecipe).
		Methods("POST")
	
	log.Fatal(http.ListenAndServe(":8080", router))
}
