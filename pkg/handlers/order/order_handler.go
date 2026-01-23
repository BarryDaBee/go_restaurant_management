package order

import (
	"net/http"
	"strconv"

	"github.com/BarryDaBee/go_restaurant_management/models"
	orderService "github.com/BarryDaBee/go_restaurant_management/services/order"
	"github.com/BarryDaBee/go_restaurant_management/utility"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type OrderHandler struct {
	OrderService orderService.OrderService
	Validate     *validator.Validate
}

func NewOrderHandler(service orderService.OrderService, validate *validator.Validate) *OrderHandler {
	return &OrderHandler{
		OrderService: service,
		Validate:     validate,
	}
}

type CreateOrderRequest struct {
	UserID uint    `json:"user_id" validate:"required"`
	Total  float64 `json:"total" validate:"required,gt=0"`
}

type UpdateOrderStatusRequest struct {
	Status string `json:"status" validate:"required,oneof=pending preparing ready delivered cancelled"`
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var req CreateOrderRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utility.BuildErrorResponse(
			http.StatusBadRequest,
			"invalid request body",
			err.Error(),
		))
		return
	}

	if err := h.Validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, utility.BuildErrorResponse(
			http.StatusBadRequest,
			"validation failed",
			utility.ValidationResponse(err, h.Validate),
		))
		return
	}

	order := &models.Order{
		UserID: req.UserID,
		Total:  req.Total,
		Status: "pending",
	}

	createdOrder, err := h.OrderService.CreateOrder(order)
	if err != nil {
		appErr, ok := err.(*utility.AppError)
		if ok {
			c.JSON(appErr.StatusCode(), utility.BuildErrorResponse(
				appErr.StatusCode(),
				appErr.Error(),
				nil,
			))
			return
		}
		c.JSON(http.StatusInternalServerError, utility.BuildErrorResponse(
			http.StatusInternalServerError,
			"internal server error",
			nil,
		))
		return
	}

	c.JSON(http.StatusCreated, utility.BuildSuccessResponse(
		http.StatusCreated,
		"order created successfully",
		createdOrder,
	))
}

func (h *OrderHandler) GetOrderByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, utility.BuildErrorResponse(
			http.StatusBadRequest,
			"invalid order id",
			nil,
		))
		return
	}

	order, err := h.OrderService.GetOrderByID(uint(id))
	if err != nil {
		appErr, ok := err.(*utility.AppError)
		if ok {
			c.JSON(appErr.StatusCode(), utility.BuildErrorResponse(
				appErr.StatusCode(),
				appErr.Error(),
				nil,
			))
			return
		}
		c.JSON(http.StatusInternalServerError, utility.BuildErrorResponse(
			http.StatusInternalServerError,
			"internal server error",
			nil,
		))
		return
	}

	c.JSON(http.StatusOK, utility.BuildSuccessResponse(
		http.StatusOK,
		"order retrieved successfully",
		order,
	))
}

func (h *OrderHandler) GetAllOrders(c *gin.Context) {
	orders, err := h.OrderService.GetAllOrders()
	if err != nil {
		appErr, ok := err.(*utility.AppError)
		if ok {
			c.JSON(appErr.StatusCode(), utility.BuildErrorResponse(
				appErr.StatusCode(),
				appErr.Error(),
				nil,
			))
			return
		}
		c.JSON(http.StatusInternalServerError, utility.BuildErrorResponse(
			http.StatusInternalServerError,
			"internal server error",
			nil,
		))
		return
	}

	c.JSON(http.StatusOK, utility.BuildSuccessResponse(
		http.StatusOK,
		"orders retrieved successfully",
		orders,
	))
}

func (h *OrderHandler) GetOrdersByUser(c *gin.Context) {
	userIDParam := c.Param("user_id")
	userID, err := strconv.ParseUint(userIDParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, utility.BuildErrorResponse(
			http.StatusBadRequest,
			"invalid user id",
			nil,
		))
		return
	}

	orders, err := h.OrderService.GetOrdersByUserID(uint(userID))
	if err != nil {
		appErr, ok := err.(*utility.AppError)
		if ok {
			c.JSON(appErr.StatusCode(), utility.BuildErrorResponse(
				appErr.StatusCode(),
				appErr.Error(),
				nil,
			))
			return
		}
		c.JSON(http.StatusInternalServerError, utility.BuildErrorResponse(
			http.StatusInternalServerError,
			"internal server error",
			nil,
		))
		return
	}

	c.JSON(http.StatusOK, utility.BuildSuccessResponse(
		http.StatusOK,
		"orders retrieved successfully",
		orders,
	))
}

func (h *OrderHandler) UpdateOrderStatus(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, utility.BuildErrorResponse(
			http.StatusBadRequest,
			"invalid order id",
			nil,
		))
		return
	}

	var req UpdateOrderStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utility.BuildErrorResponse(
			http.StatusBadRequest,
			"invalid request body",
			err.Error(),
		))
		return
	}

	if err := h.Validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, utility.BuildErrorResponse(
			http.StatusBadRequest,
			"validation failed",
			utility.ValidationResponse(err, h.Validate),
		))
		return
	}

	updatedOrder, err := h.OrderService.UpdateOrderStatus(uint(id), req.Status)
	if err != nil {
		appErr, ok := err.(*utility.AppError)
		if ok {
			c.JSON(appErr.StatusCode(), utility.BuildErrorResponse(
				appErr.StatusCode(),
				appErr.Error(),
				nil,
			))
			return
		}
		c.JSON(http.StatusInternalServerError, utility.BuildErrorResponse(
			http.StatusInternalServerError,
			"internal server error",
			nil,
		))
		return
	}

	c.JSON(http.StatusOK, utility.BuildSuccessResponse(
		http.StatusOK,
		"order status updated successfully",
		updatedOrder,
	))
}

func (h *OrderHandler) DeleteOrder(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, utility.BuildErrorResponse(
			http.StatusBadRequest,
			"invalid order id",
			nil,
		))
		return
	}

	err = h.OrderService.DeleteOrder(uint(id))
	if err != nil {
		appErr, ok := err.(*utility.AppError)
		if ok {
			c.JSON(appErr.StatusCode(), utility.BuildErrorResponse(
				appErr.StatusCode(),
				appErr.Error(),
				nil,
			))
			return
		}
		c.JSON(http.StatusInternalServerError, utility.BuildErrorResponse(
			http.StatusInternalServerError,
			"internal server error",
			nil,
		))
		return
	}

	c.JSON(http.StatusOK, utility.BuildSuccessResponse(
		http.StatusOK,
		"order deleted successfully",
		nil,
	))
}
