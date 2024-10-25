package controllers

import (
	"encoding/json"
	"net/http"
	"smart-home-project/models"
	"smart-home-project/services"
)

type HomeController struct {
	HouseService *services.HouseService
}

func NewHomeController(house *models.House) *HomeController {
	houseService := services.NewHouseService(house)
	return &HomeController{
		HouseService: houseService,
	}
}

func (h *HomeController) ToggleMainDoorLock(w http.ResponseWriter, req *http.Request) {
	h.HouseService.ToggleMainDoorLock()
	json.NewEncoder(w).Encode(h.HouseService.House)
}

func (h *HomeController) SetTotalTemperature(w http.ResponseWriter, req *http.Request) {
    var data struct {
        Temperature float64 `json:"temperature"`
    }

    err := json.NewDecoder(req.Body).Decode(&data)
    if err != nil {
        http.Error(w, "Invalid data format", http.StatusBadRequest)
        return
    }

    if data.Temperature < 10.0 || data.Temperature > 30.0 {
        http.Error(w, "Temperature must be between 10.0 and 30.0 degrees.", http.StatusBadRequest)
        return
    }

    err = h.HouseService.SetTotalTemperature(data.Temperature)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    json.NewEncoder(w).Encode(h.HouseService.House)
}

func (h *HomeController) ToggleCorridorLight(w http.ResponseWriter, req *http.Request) {
	h.HouseService.ToggleCorridorLight()
	json.NewEncoder(w).Encode(h.HouseService.House)
}

func (h *HomeController) GetHomeStatus(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(h.HouseService.House)
}
