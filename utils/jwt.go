package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Secret key for JWT (Keep this safe)
var JwtSecret = []byte("k41edsja$#cn")

// GenerateToken creates a JWT token
func GenerateJWT(email string, role string, userID uint) (string) {
	claims := jwt.MapClaims{
		"email":  email,
		"userID": userID,
		"role":   role,
		"expiry": time.Now().Add(time.Hour * 120).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtSecret)
	if err != nil {
		return ""
	}

	return tokenString

}
