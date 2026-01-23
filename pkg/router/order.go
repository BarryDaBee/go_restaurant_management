package router

import (
	"github.com/BarryDaBee/go_restaurant_management/internal/config"
	"github.com/BarryDaBee/go_restaurant_management/internal/repository/storage"
	orderHandler "github.com/BarryDaBee/go_restaurant_management/pkg/handlers/order"
	"github.com/BarryDaBee/go_restaurant_management/pkg/middleware"
	orderService "github.com/BarryDaBee/go_restaurant_management/services/order"
	"github.com/BarryDaBee/go_restaurant_management/utility"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Order(r *gin.Engine, apiVersion string, validator *validator.Validate, logger *utility.Logger, storage storage.Storage, configuration *config.Configuration) {
	service := orderService.NewOrderService(storage)
	handler := orderHandler.NewOrderHandler(service, validator)

	orderGroup := r.Group("/" + apiVersion + "/orders")
	orderGroup.Use(middleware.Authorize(configuration))
	{
		orderGroup.GET("/", handler.GetAllOrders)
		orderGroup.POST("/", handler.CreateOrder)
		orderGroup.GET("/:id", handler.GetOrderByID)
		orderGroup.PUT("/:id/status", handler.UpdateOrderStatus)
		orderGroup.DELETE("/:id", handler.DeleteOrder)
		orderGroup.GET("/user/:user_id", handler.GetOrdersByUser)
	}
}
