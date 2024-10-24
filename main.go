package main

import (
    "log"
    "net/http"
    "smart-home-project/api"
)

func main() {
    router := api.Router()
	
    log.Println("Smart Home API server is running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}
