package main

import (
	"github.com/jwhitaker/phloem/internal/pkg/phloem"
)

type AppConfiguration struct {
	ListenAddress string `default:":8080"`
	phloem.KafkaConfiguration
}
