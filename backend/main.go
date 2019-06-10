package main

import (
	"log"
	"os"

	"payment-app/backend/infrastructure"

	"github.com/joho/godotenv"
)

func main() {
	// load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	infrastructure.Router.Run(os.Getenv("API_SERVER_PORT"))
}
