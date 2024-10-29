package repositories

import (
	"context"
	"smart-home-project/models"
)

func ToggleMainDoorLock() error {
	_, err := DB.Exec(context.Background(), "UPDATE houses SET main_door_locked = NOT main_door_locked WHERE id = 1")
	return err
}

func UpdateHouseTemperature() error {
    var avgTemperature float64

    err := DB.QueryRow(context.Background(), "SELECT AVG(temperature) FROM rooms").Scan(&avgTemperature)
    if err != nil {
        return err
    }

    _, err = DB.Exec(context.Background(), "UPDATE houses SET total_temperature = $1 WHERE id = 1", avgTemperature)
    return err
}

func GetHouseStatus() (*models.House, error) {
	house := models.House{}
	err := DB.QueryRow(context.Background(), "SELECT main_door_locked, total_temperature, corridor_light_on FROM houses WHERE id = 1").Scan(
		&house.MainDoorLocked,
		&house.TotalTemperature,
		&house.Corridor.LightOn,
	)
	if err != nil {
		return nil, err
	}
	return &house, nil
}

func ToggleCorridorLight() error {
    _, err := DB.Exec(context.Background(), "UPDATE houses SET corridor_light_on = NOT corridor_light_on WHERE id = 1")
    return err
}
