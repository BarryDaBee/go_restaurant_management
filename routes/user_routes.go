package routes

import (
	"github.com/BarryDaBee/go_restaurant_management/controllers"
	"github.com/BarryDaBee/go_restaurant_management/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(route *gin.Engine) {
	user := route.Group("/user")
	user.Use(middleware.AuthMiddleware())
	{
		user.POST("/register", controllers.Register)
		user.POST("/login", controllers.Login2)
	}
}
