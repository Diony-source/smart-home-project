package models

type Room struct {
	Name string
	LightOn bool
	Temperature float64
	Devices map[string]Devices
}