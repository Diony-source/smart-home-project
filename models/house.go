package models

type Corridor struct {
	LightOn bool
}

type House struct {
	Rooms map[string]Room
	Corridor Corridor
	MainDoorLocked bool
	TotalTemperature int
}