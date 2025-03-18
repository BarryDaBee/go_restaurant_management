package helpers

import (
	"errors"
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
	// set expiration time
	expirationTime := jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 30))

	// define claims
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: expirationTime,
		},
	}

	// generate new token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// sign and return token
	return token.SignedString(jwtSecret)
}

func ValidateJWT(tokenStr string) (uint, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (any, error) {
		return jwtSecret, nil
	})
	if err != nil || !token.Valid {
		return 0, errors.New("invalid token")
	}
	return claims.UserID, nil
}
