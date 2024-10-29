package main

import (
	"log"
	"net/http"
	"smart-home-project/api"
	"smart-home-project/repositories"
)

func main() {
	err := repositories.ConnectDatabase()
	if err != nil {
		log.Fatal("Database connection failed: ", err)
	}

	router := api.Router()

	log.Println("Smart Home API server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
