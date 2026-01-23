package router

import (
	"github.com/BarryDaBee/go_restaurant_management/internal/config"
	"github.com/BarryDaBee/go_restaurant_management/internal/repository/storage"
	foodHandler "github.com/BarryDaBee/go_restaurant_management/pkg/handlers/food"
	"github.com/BarryDaBee/go_restaurant_management/pkg/middleware"
	foodService "github.com/BarryDaBee/go_restaurant_management/services/food"
	"github.com/BarryDaBee/go_restaurant_management/utility"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Food(r *gin.Engine, apiVersion string, validator *validator.Validate, logger *utility.Logger, storage storage.Storage, configuration *config.Configuration) {
	service := foodService.NewFoodService(storage)
	handler := foodHandler.NewFoodHandler(service, validator)

	foodGroup := r.Group("/" + apiVersion + "/foods")
	foodGroup.Use(middleware.Authorize(configuration))
	{
		foodGroup.GET("/", handler.GetAllFoods)
		foodGroup.POST("/", handler.CreateFood)
		foodGroup.GET("/:id", handler.GetFoodByID)
		foodGroup.PUT("/:id", handler.UpdateFood)
		foodGroup.DELETE("/:id", handler.DeleteFood)
	}
}
