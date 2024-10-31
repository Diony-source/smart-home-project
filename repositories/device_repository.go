package repositories

import (
	"context"
	"errors"
	"smart-home-project/models"
)

func ToggleDeviceStatus(deviceID int) error {
	var device models.Devices
	err := DB.QueryRow(context.Background(), "SELECT id, is_on FROM devices WHERE id = $1", deviceID).Scan(&device.Name, &device.IsOn)
	if err != nil {
		return errors.New("device not found")
	}

	// Cihaz durumunu tersine çevir (açık/kapalı)
	newStatus := !device.IsOn
	_, err = DB.Exec(context.Background(), "UPDATE devices SET is_on = $1 WHERE id = $2", newStatus, deviceID)
	if err != nil {
		return errors.New("failed to update device status")
	}

	return nil
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
