package health

import (
	"net/http"

	"github.com/BarryDaBee/go_restaurant_management/utility"
	"github.com/gin-gonic/gin"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, utility.BuildSuccessResponse(
		http.StatusOK,
		"Welcome to Restaurant API",
		gin.H{"status": "healthy"},
	))
}
