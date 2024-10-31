package repositories

import (
	"context"
	"log"
	"smart-home-project/models"
)

// Kullanıcıyı veritabanına ekler
func CreateUser(user *models.User) error {
	_, err := DB.Exec(context.Background(), "INSERT INTO users (username, email, password_hash, role) VALUES ($1, $2, $3, $4)", user.Username, user.Email, user.PasswordHash, user.Role)
	if err != nil {
		log.Printf("Error executing insert query: %v", err) // Sorgu hatasını logla
	}
	return err
}

// Kullanıcıyı kullanıcı adına göre getirir
func GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := DB.QueryRow(context.Background(), "SELECT id, username, email, password_hash, role FROM users WHERE username=$1", username).
		Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.Role)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
