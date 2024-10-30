package main

import (
	"log"
	"net/http"
	"smart-home-project/api"
	"smart-home-project/repositories"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using default environment variables")
	}

	err = repositories.ConnectDatabase()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	router := api.Router()

	log.Println("Smart Home API server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
