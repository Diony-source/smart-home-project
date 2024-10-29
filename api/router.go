package api

import (
	"smart-home-project/controllers"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
    router := mux.NewRouter()

    homeController := controllers.NewHomeController()
    roomController := controllers.NewRoomController()

    router.HandleFunc("/api/house/lock", homeController.ToggleMainDoorLock).Methods("POST")
    router.HandleFunc("/api/house/temperature", homeController.SetTotalTemperature).Methods("POST")
    router.HandleFunc("/api/house/status", homeController.GetHouseStatus).Methods("GET")
    router.HandleFunc("/api/house/corridor-light", homeController.ToggleCorridorLight).Methods("POST")

    router.HandleFunc("/api/room/light", roomController.ToggleLight).Methods("POST")
    router.HandleFunc("/api/room/device", roomController.ToggleDevice).Methods("POST")
    router.HandleFunc("/api/room/temperature", roomController.SetTemperature).Methods("POST")

    return router
}
