package repositories

import (
    "encoding/json"
    "os"
    "smart-home-project/models"
    "log"
)

// RoomRepository manages room data persistence.
type RoomRepository struct {
    filePath string
}

// NewRoomRepository creates a new RoomRepository with a given file path.
func NewRoomRepository(filePath string) *RoomRepository {
    return &RoomRepository{filePath: filePath}
}

// LoadRooms loads rooms data from a JSON file.
func (r *RoomRepository) LoadRooms() (map[string]models.Room, error) {
    file, err := os.Open(r.filePath)
    if err != nil {
        log.Println("Error opening file:", err)
        return nil, err
    }
    defer file.Close()

    rooms := make(map[string]models.Room)
    decoder := json.NewDecoder(file)
    err = decoder.Decode(&rooms)
    if err != nil {
        log.Println("Error decoding JSON:", err)
        return nil, err
    }
    return rooms, nil
}

// SaveRooms saves rooms data to a JSON file.
func (r *RoomRepository) SaveRooms(rooms map[string]models.Room) error {
    file, err := os.Create(r.filePath)
    if err != nil {
        log.Println("Error creating file:", err)
        return err
    }
    defer file.Close()

    encoder := json.NewEncoder(file)
    encoder.SetIndent("", "  ")
    err = encoder.Encode(rooms)
    if err != nil {
        log.Println("Error encoding JSON:", err)
        return err
    }
    return nil
}
