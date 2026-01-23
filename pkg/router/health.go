package router

import (
	"github.com/BarryDaBee/go_restaurant_management/internal/config"
	"github.com/BarryDaBee/go_restaurant_management/internal/repository/storage"
	"github.com/BarryDaBee/go_restaurant_management/pkg/handlers/health"
	"github.com/BarryDaBee/go_restaurant_management/utility"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Health(r *gin.Engine, apiVersion string, validator *validator.Validate, logger *utility.Logger, storage storage.Storage, configuration *config.Configuration) {
	healthHandler := health.NewHealthHandler()

	r.GET("/", healthHandler.HealthCheck)
	r.GET("/"+apiVersion+"/health", healthHandler.HealthCheck)
}
