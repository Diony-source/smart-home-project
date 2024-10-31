package controllers

import (
	"encoding/json"
	"net/http"
	"smart-home-project/models"
	"smart-home-project/repositories"
	"smart-home-project/utils"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct{}

func NewAuthController() *AuthController {
	return &AuthController{}
}

// Kullanıcı Kayıt İşlemi
func (a *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	logrus.Info("Registering new user")

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		logrus.Error("Invalid input for user registration: ", err)
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		logrus.Error("Failed to hash password: ", err)
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)

	if err := repositories.CreateUser(&user); err != nil {
		logrus.Error("Failed to create user: ", err)
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	logrus.Info("User registered successfully")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("User registered successfully")
}

// Kullanıcı Giriş İşlemi
func (a *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	logrus.Info("User login attempt")

	var input models.User
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		logrus.Error("Invalid input for login: ", err)
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	user, err := repositories.GetUserByUsername(input.Username)
	if err != nil {
		logrus.Error("User not found: ", err)
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		logrus.Error("Invalid password for user: ", input.Username)
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateJWT(user.Username, user.Role)
	if err != nil {
		logrus.Error("Failed to generate token: ", err)
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	logrus.Info("User logged in successfully")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
