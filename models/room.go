package models

type SecurityStatus struct {
	DoorOpen bool
	WindowOpen bool
	MotionDetected bool
}

type Room struct {
	Name string
	LightOn bool
	Temperature int
	Devices map[string]Devices
	security SecurityStatus
}