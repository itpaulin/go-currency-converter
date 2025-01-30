package api

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	API_URL string
	APP_ID string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	API_URL = os.Getenv("API_URL")
	APP_ID = os.Getenv("APP_ID")
}