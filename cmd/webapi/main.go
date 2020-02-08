package main

import (
	"log"

	"net/http"
	"github.com/jwhitaker/phloem/internal/pkg/commons"
	

	"github.com/gorilla/mux"
	"github.com/jwhitaker/phloem/internal/app/webapi"
	"github.com/jwhitaker/phloem/internal/pkg/phloem"
)

func main() {
	log.Println("Starting server")

	var appConfiguration AppConfiguration

	commons.LoadConfiguration(&appConfiguration)

	producer := phloem.NewKafkaProducer(appConfiguration.KafkaConfiguration)

	recipeController := webapi.NewRecipeController(producer)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/recipe", recipeController.SaveRecipe).
		Methods("POST")

	log.Println("Server started")
	log.Fatal(http.ListenAndServe(appConfiguration.ListenAddress, router))
}
