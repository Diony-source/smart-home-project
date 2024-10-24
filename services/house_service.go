package services

import "smart-home-project/models"

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

func (h *HouseService) SetTotalTemperature(temp int) {
    h.House.TotalTemperature = temp
    for roomName, room := range h.House.Rooms {
        room.Temperature = temp
        h.House.Rooms[roomName] = room
    }
}

func (h *HouseService) ToggleCorridorLight() {
    h.House.Corridor.LightOn = !h.House.Corridor.LightOn
}
