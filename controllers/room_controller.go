package controllers

import (
	"encoding/json"
	"net/http"
	"smart-home-project/repositories"
	"strconv"

	"github.com/sirupsen/logrus"
)

type RoomController struct{}

func NewRoomController() *RoomController {
	return &RoomController{}
}

func (r *RoomController) ToggleLight(w http.ResponseWriter, req *http.Request) {
	roomID, err := strconv.Atoi(req.URL.Query().Get("room_id"))
	if err != nil {
		logrus.Error("Invalid room ID for light toggle: ", err)
		http.Error(w, "Invalid room ID", http.StatusBadRequest)
		return
	}

	err = repositories.ToggleRoomLight(roomID)
	if err != nil {
		logrus.Error("Failed to toggle room light: ", err)
		http.Error(w, "Failed to toggle room light", http.StatusInternalServerError)
		return
	}

	logrus.Info("Room light toggled successfully for room ID: ", roomID)
	json.NewEncoder(w).Encode("Room light toggled successfully")
}

func (r *RoomController) ToggleDevice(w http.ResponseWriter, req *http.Request) {
	deviceID, err := strconv.Atoi(req.URL.Query().Get("device_id"))
	if err != nil {
		logrus.Error("Invalid device ID for toggle: ", err)
		http.Error(w, "Invalid device ID", http.StatusBadRequest)
		return
	}

	err = repositories.ToggleDeviceStatus(deviceID)
	if err != nil {
		logrus.Error("Failed to toggle device status: ", err)
		http.Error(w, "Failed to toggle device status", http.StatusInternalServerError)
		return
	}

	logrus.Info("Device status toggled successfully for device ID: ", deviceID)
	json.NewEncoder(w).Encode("Device status toggled successfully")
}

func (r *RoomController) SetTemperature(w http.ResponseWriter, req *http.Request) {
	logrus.Info("Setting room temperature")

	var data struct {
		Temperature float64 `json:"temperature"`
	}

	roomID, err := strconv.Atoi(req.URL.Query().Get("room_id"))
	if err != nil {
		logrus.Error("Invalid room ID for setting temperature: ", err)
		http.Error(w, "Invalid room ID", http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(req.Body).Decode(&data)
	if err != nil || data.Temperature < 10.0 || data.Temperature > 30.0 {
		logrus.Error("Invalid temperature value: ", err)
		http.Error(w, "Invalid temperature", http.StatusBadRequest)
		return
	}

	err = repositories.UpdateRoomTemperature(roomID, data.Temperature)
	if err != nil {
		logrus.Error("Failed to update room temperature: ", err)
		http.Error(w, "Failed to update room temperature", http.StatusInternalServerError)
		return
	}

	err = repositories.UpdateHouseTemperature()
	if err != nil {
		logrus.Error("Failed to update house average temperature: ", err)
		http.Error(w, "Failed to update house temperature", http.StatusInternalServerError)
		return
	}

	logrus.Info("Room temperature and house average temperature updated successfully")
	json.NewEncoder(w).Encode("Room temperature and house average temperature updated successfully")
}
