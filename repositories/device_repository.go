package repositories

import (
	"context"
	"smart-home-project/models"
)

func ToggleDeviceStatus(deviceID int) error {
	_, err := DB.Exec(context.Background(), "UPDATE devices SET is_on = NOT is_on WHERE id = $1", deviceID)
	return err
}

func GetDeviceByID(deviceID int) (*models.Devices, error) {
	device := models.Devices{}
	err := DB.QueryRow(context.Background(), "SELECT name, is_on FROM devices WHERE id = $1", deviceID).Scan(
		&device.Name,
		&device.IsOn,
	)
	if err != nil {
		return nil, err
	}
	return &device, nil
}
