package main

import (
	"log"
	"net/http"
	"os"

	"github.com/BarryDaBee/go_restaurant_management/database"
	"github.com/BarryDaBee/go_restaurant_management/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// port := os.Getenv("PORT")

	// if port == "" {
	// 	port = "8080"
	// }

	database.ConnectDatabase()

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to Restaurant API"})
	})

	routes.UserRoutes(router)
	routes.FoodRoutes(router)

	log.Printf("Server is running on port %s", os.Getenv("PORT"))

	router.Run()
}
