package services

import (
	"errors"
	"math"
	"smart-home-project/models"
)

type RoomService struct {
	Room  *models.Room
	House *models.House
}

func NewRoomService(room *models.Room, house *models.House) *RoomService {
	return &RoomService{
		Room:  room,
		House: house,
	}
}

func (r *RoomService) ToggleLight() {
	r.Room.LightOn = !r.Room.LightOn
}

func (r *RoomService) ToggleDevice(deviceName string) {
	if device, exists := r.Room.Devices[deviceName]; exists {
		device.IsOn = !device.IsOn
		r.Room.Devices[deviceName] = device
	}
}

func (r *RoomService) SetTemperature(newTemp float64) error {
	currentTemp := r.Room.Temperature

	if math.Abs(currentTemp-newTemp) > 10.0 {
		return errors.New("temperature change cannot exceed 10.0 degrees at once")
	}

	r.Room.Temperature = newTemp

	if r.House != nil {
		r.updateHouseTemperature()
	}

	return nil
}

func (r *RoomService) updateHouseTemperature() {
	totalTemp := 0.0
	roomCount := 0

	for _, room := range r.House.Rooms {
		totalTemp += room.Temperature
		roomCount++
	}

	if roomCount > 0 {
		r.House.TotalTemperature = totalTemp / float64(roomCount)
	}
}
