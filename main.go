package main

import (
	"net/http"
	"smart-home-project/api"
	"smart-home-project/repositories"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	// Log yapılandırması
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.InfoLevel)

	logrus.Info("Starting Smart Home API server...")

	err := godotenv.Load()
	if err != nil {
		logrus.Warn("No .env file found, using default environment variables")
	}

	err = repositories.ConnectDatabase()
	if err != nil {
		logrus.Fatal("Failed to connect to database:", err)
	}

	router := api.Router()

	logrus.Info("Smart Home API server is running on port 8080")
	http.ListenAndServe(":8080", router)
}
