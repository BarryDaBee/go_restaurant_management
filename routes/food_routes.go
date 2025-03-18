package routes

import (
	"github.com/BarryDaBee/go_restaurant_management/controllers"
	"github.com/BarryDaBee/go_restaurant_management/middleware"
	"github.com/gin-gonic/gin"
)

func FoodRoutes(route *gin.Engine) {
	food := route.Group("/food")
	food.Use(middleware.AuthMiddleware())
	{
		food.GET("/", controllers.GetFoods)
		food.POST("/", controllers.CreateFood)
		food.GET("/:id", controllers.GetFoodById)
		food.PUT("/:id", controllers.UpdateFood)
		food.DELETE("/:id", controllers.DeleteFood)
	}
}
