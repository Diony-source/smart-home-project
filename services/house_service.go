package services

import (
	"errors"
	"math"
	"smart-home-project/models"
)

type HouseService struct {
	House *models.House
}

func NewHouseService(house *models.House) *HouseService {
	return &HouseService{
		House: house,
	}
}

func (h *HouseService) ToggleMainDoorLock() {
	h.House.MainDoorLocked = !h.House.MainDoorLocked
}

func (h *HouseService) SetTotalTemperature(newTemp float64) error {
	currentTemp := h.House.TotalTemperature

	if math.Abs(currentTemp-newTemp) > 10.0 {
		return errors.New("temperature change cannot exceed 10.0 degrees at once")
	}

	h.House.TotalTemperature = newTemp

	for roomName, room := range h.House.Rooms {
		room.Temperature = newTemp
		h.House.Rooms[roomName] = room
	}

	return nil
}

func (h *HouseService) ToggleCorridorLight() {
	h.House.Corridor.LightOn = !h.House.Corridor.LightOn
}
