package utils

import (
    "errors"
    "os"
    "time"

    "github.com/dgrijalva/jwt-go"
)

var jwtKey []byte

// JWT Anahtarını .env Dosyasından Alır
func InitJWTKey() {
    key := os.Getenv("JWT_SECRET_KEY")
    if key == "" {
        panic("JWT_SECRET_KEY is not set in environment variables")
    }
    jwtKey = []byte(key)
}

// JWT Token oluşturur
func GenerateJWT(username, role string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "username": username,
        "role":     role,
        "exp":      time.Now().Add(time.Hour * 24).Unix(),
    })

    return token.SignedString(jwtKey)
}

func ParseJWT(tokenString string) (*jwt.Token, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, errors.New("unexpected signing method")
        }
        return jwtKey, nil
    })

    return token, err
}
