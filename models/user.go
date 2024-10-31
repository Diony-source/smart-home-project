package models

type User struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	Password     string `json:"password,omitempty"`  // JSON çıktısında şifre alanı olarak kullanılacak
	PasswordHash string `json:"-"`                   // JSON çıktısında görünmeyecek, yalnızca veritabanı için
	Role         string `json:"role"`
}
