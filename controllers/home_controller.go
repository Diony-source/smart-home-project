package controllers

import (
    "net/http"
    "encoding/json"
    "smart-home-project/services"
    "smart-home-project/models"
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
        Temperature int `json:"temperature"`
    }
    err := json.NewDecoder(req.Body).Decode(&data)
    if err != nil {
        http.Error(w, "Geçersiz veri formatı", http.StatusBadRequest)
        return
    }
    h.HouseService.SetTotalTemperature(data.Temperature)
    json.NewEncoder(w).Encode(h.HouseService.House)
}

func (h *HomeController) ToggleCorridorLight(w http.ResponseWriter, req *http.Request) {
    h.HouseService.ToggleCorridorLight()
    json.NewEncoder(w).Encode(h.HouseService.House)
}
