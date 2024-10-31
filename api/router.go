// api/router.go
package api

import (
	"net/http"
	"smart-home-project/controllers"
	"smart-home-project/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	homeController := controllers.NewHomeController()
	roomController := controllers.NewRoomController()
	authController := controllers.NewAuthController()

	// Auth Routes
	router.HandleFunc("/api/auth/register", authController.Register).Methods("POST")
	router.HandleFunc("/api/auth/login", authController.Login).Methods("POST")

	// Protected Routes with Role-based Authorization
	router.Handle("/api/house/status", middleware.RoleAuth("admin")(http.HandlerFunc(homeController.GetHouseStatus))).Methods("GET")
	router.Handle("/api/house/corridor-light", middleware.RoleAuth("admin", "user", "guest")(http.HandlerFunc(homeController.ToggleCorridorLight))).Methods("POST")
	router.Handle("/api/house/lock", middleware.RoleAuth("admin")(http.HandlerFunc(homeController.ToggleMainDoorLock))).Methods("POST")
	router.Handle("/api/house/temperature", middleware.RoleAuth("admin", "user")(http.HandlerFunc(homeController.SetTotalTemperature))).Methods("POST")

	router.Handle("/api/room/light", middleware.RoleAuth("admin", "user", "guest")(http.HandlerFunc(roomController.ToggleLight))).Methods("POST")
	router.Handle("/api/room/temperature", middleware.RoleAuth("admin", "user")(http.HandlerFunc(roomController.SetTemperature))).Methods("POST")

	router.Handle("/api/room/device", middleware.RoleAuth("admin", "user")(http.HandlerFunc(roomController.ToggleDevice))).Methods("POST")

	return router
}
