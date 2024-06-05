package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnveronmentVariabel() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Failed load .env file!")
	}
}
