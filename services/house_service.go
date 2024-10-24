package services

import "smart-home-project/models"

type HouseServices struct {
	House *models.House
}

func NewHouseService(house *models.House) *HouseServices {
	return &HouseServices{
		House: house,
	}
}

func (h *HouseServices) ToogleMainDoorLock() {
	h.House.MainDoorLocked = !h.House.MainDoorLocked
}

func (h *HouseServices) SetTotalTemperature(temp int) {
	h.House.TotalTemperature = temp
	for roomName, room := range h.House.Rooms {
		room.Temperature = temp
		h.House.Rooms[roomName] = room
	}
}

func (h *HouseServices) ToogleCorridorLight() {
	h.House.Corridor.LightOn = !h.House.Corridor.LightOn
}