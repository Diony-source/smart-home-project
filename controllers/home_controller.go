package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"smart-home-project/repositories"
)

type HomeController struct{}

func NewHomeController() *HomeController {
	return &HomeController{}
}

func (h *HomeController) ToggleMainDoorLock(w http.ResponseWriter, req *http.Request) {
	err := repositories.ToggleMainDoorLock()
	if err != nil {
		http.Error(w, "Failed to toggle main door lock", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode("Main door lock toggled successfully")
}

func (h *HomeController) SetTotalTemperature(w http.ResponseWriter, req *http.Request) {
    err := repositories.UpdateHouseTemperature()
    if err != nil {
        http.Error(w, "Failed to update house temperature", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode("House temperature updated based on room average successfully")
}

func (h *HomeController) GetHouseStatus(w http.ResponseWriter, req *http.Request) {
	house, err := repositories.GetHouseStatus()
	if err != nil {
		log.Println("Error getting house status:", err)
		http.Error(w, "Failed to get house status", http.StatusInternalServerError)
		return
	}

	rooms, err := repositories.GetRoomsStatus()
	if err != nil {
		log.Println("Error getting rooms status:", err)
		http.Error(w, "Failed to get rooms status", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"house": house,
		"rooms": rooms,
	}

	json.NewEncoder(w).Encode(response)
}

func (h *HomeController) ToggleCorridorLight(w http.ResponseWriter, req *http.Request) {
    err := repositories.ToggleCorridorLight()
    if err != nil {
        log.Println("Error toggling corridor light:", err)
        http.Error(w, "Failed to toggle corridor light", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode("Corridor light toggled successfully")
}
