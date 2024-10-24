package repositories

import (
    "encoding/json"
    "os"
    "smart-home-project/models"
    "log"
)

// HouseRepository manages house data persistence.
type HouseRepository struct {
    filePath string
}

// NewHouseRepository creates a new HouseRepository with a given file path.
func NewHouseRepository(filePath string) *HouseRepository {
    return &HouseRepository{filePath: filePath}
}

// LoadHouse loads house data from a JSON file.
func (r *HouseRepository) LoadHouse() (*models.House, error) {
    file, err := os.Open(r.filePath)
    if err != nil {
        log.Println("Error opening file:", err)
        return nil, err
    }
    defer file.Close()

    var house models.House
    decoder := json.NewDecoder(file)
    err = decoder.Decode(&house)
    if err != nil {
        log.Println("Error decoding JSON:", err)
        return nil, err
    }
    return &house, nil
}

// SaveHouse saves house data to a JSON file.
func (r *HouseRepository) SaveHouse(house *models.House) error {
    file, err := os.Create(r.filePath)
    if err != nil {
        log.Println("Error creating file:", err)
        return err
    }
    defer file.Close()

    encoder := json.NewEncoder(file)
    encoder.SetIndent("", "  ")
    err = encoder.Encode(house)
    if err != nil {
        log.Println("Error encoding JSON:", err)
        return err
    }
    return nil
}
