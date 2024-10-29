package models

type Room struct {
	ID          int                `json:"id"`
	Name        string             `json:"name"`
	LightOn     bool               `json:"light_on"`
	Temperature float64            `json:"temperature"`
	Devices     map[string]Devices `json:"devices"`
}
