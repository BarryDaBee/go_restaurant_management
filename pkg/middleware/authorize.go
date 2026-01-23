package middleware

import (
	"net/http"
	"strings"

	"github.com/BarryDaBee/go_restaurant_management/internal/config"
	"github.com/BarryDaBee/go_restaurant_management/utility"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

func Authorize(configuration *config.Configuration) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, utility.BuildErrorResponse(
				http.StatusUnauthorized,
				"authorization token missing",
				nil,
			))
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, utility.BuildErrorResponse(
				http.StatusUnauthorized,
				"invalid token format",
				nil,
			))
			c.Abort()
			return
		}

		tokenString := parts[1]
		claims := &Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(configuration.Server.Secret), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, utility.BuildErrorResponse(
				http.StatusUnauthorized,
				"invalid or expired token",
				nil,
			))
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Next()
	}
}
