package commons

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

func LoadConfiguration(obj interface{}) {
	err := envconfig.Process("APP", obj)

	if err != nil {
		log.Fatalf("Failed to load configuration %+v", err)
	}

	log.Printf("Loading configuration : %+v", obj)
}
