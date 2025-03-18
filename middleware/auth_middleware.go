package middleware

import (
	"net/http"
	"strings"

	"github.com/BarryDaBee/go_restaurant_management/helpers"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get Authorization header
		authHeader := c.GetHeader("Authorization")

		//1. Check if Authorization header is missing
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token missing"})
			c.Abort()
			return
		}

		//2. Ensure token is in "Bearer <token>" format
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		//3. Extract token
		tokenString := parts[1]

		//4. Validate token
		userID, err := helpers.ValidateJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		//5. Store userID and role in Gin context
		c.Set("user_id", userID)

		// Proceed to next handler
		c.Next()
	}
}
