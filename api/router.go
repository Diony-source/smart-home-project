package api

import (
	"smart-home-project/controllers"
	"smart-home-project/models"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	// Initialize a sample house and create controllers.
	house := initializeHouse()
	homeController := controllers.NewHomeController(&house)

	// Fixing the issue by using temporary variables.
	livingRoom := house.Rooms["Living Room"]
	livingRoomController := controllers.NewRoomController(&livingRoom, &house)

	kitchen := house.Rooms["Kitchen"]
	kitchenController := controllers.NewRoomController(&kitchen, &house)

	bedroom := house.Rooms["Bedroom"]
	bedroomController := controllers.NewRoomController(&bedroom, &house)

	bathroom := house.Rooms["Bathroom"]
	bathroomController := controllers.NewRoomController(&bathroom, &house)

	// General house routes.
	router.HandleFunc("/api/house/lock", homeController.ToggleMainDoorLock).Methods("POST")
	router.HandleFunc("/api/house/temperature", homeController.SetTotalTemperature).Methods("POST")
	router.HandleFunc("/api/house/corridor-light", homeController.ToggleCorridorLight).Methods("POST")
	router.HandleFunc("/api/house/status", homeController.GetHomeStatus).Methods("GET")

	// Routes for rooms.
	router.HandleFunc("/api/rooms/living-room/light", livingRoomController.ToggleLight).Methods("POST")
	router.HandleFunc("/api/rooms/living-room/device", livingRoomController.ToggleDevice).Methods("POST")
	router.HandleFunc("/api/rooms/living-room/temperature", livingRoomController.SetTemperature).Methods("POST")
	router.HandleFunc("/api/rooms/living-room/status", livingRoomController.GetRoomStatus).Methods("GET")

	router.HandleFunc("/api/rooms/kitchen/light", kitchenController.ToggleLight).Methods("POST")
	router.HandleFunc("/api/rooms/kitchen/device", kitchenController.ToggleDevice).Methods("POST")
	router.HandleFunc("/api/rooms/kitchen/temperature", kitchenController.SetTemperature).Methods("POST")
	router.HandleFunc("/api/rooms/kitchen/status", kitchenController.GetRoomStatus).Methods("GET")

	router.HandleFunc("/api/rooms/bedroom/light", bedroomController.ToggleLight).Methods("POST")
	router.HandleFunc("/api/rooms/bedroom/device", bedroomController.ToggleDevice).Methods("POST")
	router.HandleFunc("/api/rooms/bedroom/temperature", bedroomController.SetTemperature).Methods("POST")
	router.HandleFunc("/api/rooms/bedroom/status", bedroomController.GetRoomStatus).Methods("GET")

	router.HandleFunc("/api/rooms/bathroom/light", bathroomController.ToggleLight).Methods("POST")
	router.HandleFunc("/api/rooms/bathroom/device", bathroomController.ToggleDevice).Methods("POST")
	router.HandleFunc("/api/rooms/bathroom/temperature", bathroomController.SetTemperature).Methods("POST")
	router.HandleFunc("/api/rooms/bathroom/status", bathroomController.GetRoomStatus).Methods("GET")

	return router
}

// initializeHouse creates a sample house with rooms and their initial configurations.
func initializeHouse() models.House {
    return models.House{
        Rooms: map[string]models.Room{
            "Living Room": {
                Name: "Living Room",
                LightOn: false,
                Temperature: 22,
                Devices: map[string]models.Devices{
                    "TV": {Name: "TV", IsOn: false},
                },
            },
            "Kitchen": {
                Name: "Kitchen",
                LightOn: false,
                Temperature: 22,
                Devices: map[string]models.Devices{
                    "Coffee Maker": {Name: "Coffee Maker", IsOn: false},
                },
            },
            "Bedroom": {
                Name: "Bedroom",
                LightOn: false,
                Temperature: 20,
                Devices: map[string]models.Devices{
                    "TV": {Name: "TV", IsOn: false},
                },
            },
            "Bathroom": {
                Name: "Bathroom",
                LightOn: false,
                Temperature: 24,
                Devices: map[string]models.Devices{
                    "Flush": {Name: "Flush", IsOn: false},
                },
            },
        },
        Corridor: models.Corridor{
            LightOn: false,
        },
        MainDoorLocked: false,
        TotalTemperature: 22,
    }
}
