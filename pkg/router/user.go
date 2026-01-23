package router

import (
	"github.com/BarryDaBee/go_restaurant_management/internal/config"
	"github.com/BarryDaBee/go_restaurant_management/internal/repository/storage"
	userHandler "github.com/BarryDaBee/go_restaurant_management/pkg/handlers/user"
	userService "github.com/BarryDaBee/go_restaurant_management/services/user"
	"github.com/BarryDaBee/go_restaurant_management/utility"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func User(r *gin.Engine, apiVersion string, validator *validator.Validate, logger *utility.Logger, storage storage.Storage, configuration *config.Configuration) {
	service := userService.NewUserService(storage, configuration.Server.Secret)
	handler := userHandler.NewUserHandler(service, validator)

	userGroup := r.Group("/" + apiVersion + "/users")
	{
		userGroup.POST("/register", handler.Register)
		userGroup.POST("/login", handler.Login)
	}
}
