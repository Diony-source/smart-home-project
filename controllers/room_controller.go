package controllers

import (
	"encoding/json"
	"net/http"
	"smart-home-project/repositories"
	"strconv"
)

type RoomController struct{}

func NewRoomController() *RoomController {
	return &RoomController{}
}

func (r *RoomController) ToggleLight(w http.ResponseWriter, req *http.Request) {
	roomID, err := strconv.Atoi(req.URL.Query().Get("room_id"))
	if err != nil {
		http.Error(w, "Invalid room ID", http.StatusBadRequest)
		return
	}

	err = repositories.ToggleRoomLight(roomID)
	if err != nil {
		http.Error(w, "Failed to toggle room light", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("Room light toggled successfully")
}

func (r *RoomController) ToggleDevice(w http.ResponseWriter, req *http.Request) {
	deviceID, err := strconv.Atoi(req.URL.Query().Get("device_id"))
	if err != nil {
		http.Error(w, "Invalid device ID", http.StatusBadRequest)
		return
	}

	err = repositories.ToggleDeviceStatus(deviceID)
	if err != nil {
		http.Error(w, "Failed to toggle device status", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("Device status toggled successfully")
}

func (r *RoomController) SetTemperature(w http.ResponseWriter, req *http.Request) {
    var data struct {
        Temperature float64 `json:"temperature"`
    }

    roomID, err := strconv.Atoi(req.URL.Query().Get("room_id"))
    if err != nil {
        http.Error(w, "Invalid room ID", http.StatusBadRequest)
        return
    }

    err = json.NewDecoder(req.Body).Decode(&data)
    if err != nil || data.Temperature < 10.0 || data.Temperature > 30.0 {
        http.Error(w, "Invalid temperature", http.StatusBadRequest)
        return
    }

    err = repositories.UpdateRoomTemperature(roomID, data.Temperature)
    if err != nil {
        http.Error(w, "Failed to update room temperature", http.StatusInternalServerError)
        return
    }

    err = repositories.UpdateHouseTemperature()
    if err != nil {
        http.Error(w, "Failed to update house temperature", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode("Room temperature and house average temperature updated successfully")
}