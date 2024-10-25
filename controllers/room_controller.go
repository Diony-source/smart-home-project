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

func NewRoomController(room *models.Room, house *models.House) *RoomController {
    roomService := services.NewRoomService(room, house)
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
        http.Error(w, "Device name not specified", http.StatusBadRequest)
        return
    }
    r.RoomService.ToggleDevice(deviceName)
    json.NewEncoder(w).Encode(r.RoomService.Room)
}

func (r *RoomController) SetTemperature(w http.ResponseWriter, req *http.Request) {
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

    err = r.RoomService.SetTemperature(data.Temperature)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    json.NewEncoder(w).Encode(r.RoomService.Room)
}

func (r *RoomController) GetRoomStatus(w http.ResponseWriter, req *http.Request) {
    json.NewEncoder(w).Encode(r.RoomService.Room)
}

