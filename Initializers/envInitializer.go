package Initializers

import (
	"github.com/joho/godotenv"
	"log"
)

func EnvInitializer() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
