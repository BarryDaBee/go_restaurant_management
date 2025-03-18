package routes

import (
	"github.com/BarryDaBee/go_restaurant_management/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(route *gin.Engine) {
	user := route.Group("/user")
	{
		user.POST("/register", controllers.Register)
		user.POST("/login", controllers.Login2)
	}
}
