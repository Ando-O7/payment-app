package main

import (
	"os"
	"payment-app/backend/infrastructure"
)

func main() {
	infrastructure.Router.Run(os.Getenv("API_SERVER_PORT"))
}
