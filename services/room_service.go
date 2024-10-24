package services

import "smart-home-project/models"

type RoomService struct {
    Room *models.Room
}

func NewRoomService(room *models.Room) *RoomService {
    return &RoomService{
        Room: room,
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

func (r *RoomService) SetTemperature(temp int) {
    r.Room.Temperature = temp
}
