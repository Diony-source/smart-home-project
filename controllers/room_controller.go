package controllers

import (
    "net/http"
    "encoding/json"
    "smart-home-project/services"
	"smart-home-project/models"
)

type RoomController struct {
    RoomService *services.RoomService
}

func NewRoomController(room *models.Room) *RoomController {
    roomService := services.NewRoomService(room)
    return &RoomController{
        RoomService: roomService,
    }
}

func (r *RoomController) ToggleLight(w http.ResponseWriter, req *http.Request) {
    r.RoomService.ToggleLight()
    json.NewEncoder(w).Encode(r.RoomService.Room)
}

func (r *RoomController) ToggleDevice(w http.ResponseWriter, req *http.Request) {
    deviceName := req.URL.Query().Get("device")
    if deviceName == "" {
        http.Error(w, "Cihaz adı belirtilmedi", http.StatusBadRequest)
        return
    }
    r.RoomService.ToggleDevice(deviceName)
    json.NewEncoder(w).Encode(r.RoomService.Room)
}

func (r *RoomController) SetTemperature(w http.ResponseWriter, req *http.Request) {
    var data struct {
        Temperature int `json:"temperature"`
    }
    err := json.NewDecoder(req.Body).Decode(&data)
    if err != nil {
        http.Error(w, "Geçersiz veri formatı", http.StatusBadRequest)
        return
    }
    r.RoomService.SetTemperature(data.Temperature)
    json.NewEncoder(w).Encode(r.RoomService.Room)
}
