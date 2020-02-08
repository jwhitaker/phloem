package main

import (
	"log"

	"net/http"

	"github.com/gorilla/mux"
	"github.com/jwhitaker/phloem/internal/app/webapi"
	"github.com/jwhitaker/phloem/internal/pkg/phloem"
)

func main() {
	log.Println("Starting server")

	producer := phloem.NewKafkaProducer()

	recipeController := webapi.NewRecipeController(producer)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/recipe", recipeController.SaveRecipe).
		Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}
