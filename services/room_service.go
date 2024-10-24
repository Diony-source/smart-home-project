package services

import "smart-home-project/models"

type RoomServices struct {
	Room *models.Room
}

func NewRoomServices(room *models.Room) *RoomServices {
	return &RoomServices{
		Room: room,
	}
}

func (r *RoomServices) ToogleLight() {
	r.Room.LightOn = !r.Room.LightOn
}

func (r *RoomServices) ToogleDivice(deviceName string) {
	if device, exist := r.Room.Devices[deviceName]; exist {
		device.IsOn = !device.IsOn
		r.Room.Devices[deviceName] = device
	}
}

func (r *RoomServices) SetTemperature(temp int) {
	r.Room.Temperature = temp
}
