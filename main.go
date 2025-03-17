package main

import (
	"log"
	"net/http"

	"github.com/BarryDaBee/go_restaurant_management/database"
	"github.com/gin-gonic/gin"
)

func main() {
	// port := os.Getenv("PORT")

	// if port == "" {
	// 	port = "8080"
	// }

	database.con

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to Restaurant API"})
	})

	log.Println("Server is running on port :8080")

	router.Run()
}
