package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadDotEnvConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}
	log.Println(".env file loaded")
}
