package router

import (
	"net/http"

	"github.com/BarryDaBee/go_restaurant_management/internal/config"
	"github.com/BarryDaBee/go_restaurant_management/internal/repository/storage"
	"github.com/BarryDaBee/go_restaurant_management/pkg/middleware"
	"github.com/BarryDaBee/go_restaurant_management/utility"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Setup(logger *utility.Logger, validator *validator.Validate, storage storage.Storage, configuration *config.Configuration) *gin.Engine {
	if configuration.App.Mode == "release" || configuration.App.Mode == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()

	r.ForwardedByClientIP = true
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.CORS())
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.MaxMultipartMemory = 1 << 20

	ApiVersion := "v1"
	Health(r, ApiVersion, validator, logger, storage, configuration)
	User(r, ApiVersion, validator, logger, storage, configuration)
	Food(r, ApiVersion, validator, logger, storage, configuration)
	Order(r, ApiVersion, validator, logger, storage, configuration)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, utility.BuildErrorResponse(
			http.StatusNotFound,
			"page not found",
			nil,
		))
	})

	return r
}
