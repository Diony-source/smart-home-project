package repositories

import (
	"context"
	"smart-home-project/models"
)

func GetRoomsStatus() ([]models.Room, error) {
	rows, err := DB.Query(context.Background(), "SELECT id, name, light_on, temperature FROM rooms")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rooms []models.Room
	for rows.Next() {
		room := models.Room{}
		if err := rows.Scan(&room.ID, &room.Name, &room.LightOn, &room.Temperature); err != nil {
			return nil, err
		}

		device, err := GetDeviceByRoomID(room.ID)
		if err != nil {
			return nil, err
		}

		room.Devices = map[string]models.Devices{
			device.Name: device,
		}

		rooms = append(rooms, room)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return rooms, nil
}

func GetDeviceByRoomID(roomID int) (models.Devices, error) {
	device := models.Devices{}
	err := DB.QueryRow(context.Background(), "SELECT name, is_on FROM devices WHERE room_id = $1", roomID).Scan(
		&device.Name,
		&device.IsOn,
	)
	if err != nil {
		return models.Devices{}, err
	}
	return device, nil
}

func ToggleRoomLight(roomID int) error {
	_, err := DB.Exec(context.Background(), "UPDATE rooms SET light_on = NOT light_on WHERE id = $1", roomID)
	return err
}

func UpdateRoomTemperature(roomID int, newTemp float64) error {
	_, err := DB.Exec(context.Background(), "UPDATE rooms SET temperature = $1 WHERE id = $2", newTemp, roomID)
	return err
}
