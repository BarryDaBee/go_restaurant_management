package helpers

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret []byte = []byte(os.Getenv("SECRET"))

type Claims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateJWT(userID uint) (string, error) {
	expirationTime := jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 30))

	//1. define claims
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: expirationTime,
		},
	}

	//2. generate new token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//3. sign and return token
	return token.SignedString(jwtSecret)
}

func ValidateJWT(tokenStr string) (uint, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (any, error) {
		return jwtSecret, nil
	})
	if err != nil || !token.Valid {
		return 0, err
	}
	return claims.UserID, nil
}
