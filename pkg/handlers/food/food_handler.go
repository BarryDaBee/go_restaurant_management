package food

import (
	"net/http"
	"strconv"

	"github.com/BarryDaBee/go_restaurant_management/models"
	foodService "github.com/BarryDaBee/go_restaurant_management/services/food"
	"github.com/BarryDaBee/go_restaurant_management/utility"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type FoodHandler struct {
	FoodService foodService.FoodService
	Validate    *validator.Validate
}

func NewFoodHandler(service foodService.FoodService, validate *validator.Validate) *FoodHandler {
	return &FoodHandler{
		FoodService: service,
		Validate:    validate,
	}
}

func (h *FoodHandler) CreateFood(c *gin.Context) {
	var food models.Food

	if err := c.ShouldBindJSON(&food); err != nil {
		c.JSON(http.StatusBadRequest, utility.BuildErrorResponse(
			http.StatusBadRequest,
			"invalid request body",
			err.Error(),
		))
		return
	}

	if err := h.Validate.Struct(food); err != nil {
		c.JSON(http.StatusBadRequest, utility.BuildErrorResponse(
			http.StatusBadRequest,
			"validation failed",
			utility.ValidationResponse(err, h.Validate),
		))
		return
	}

	createdFood, err := h.FoodService.CreateFood(&food)
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
		"food created successfully",
		createdFood,
	))
}

func (h *FoodHandler) GetFoodByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, utility.BuildErrorResponse(
			http.StatusBadRequest,
			"invalid food id",
			nil,
		))
		return
	}

	food, err := h.FoodService.GetFoodByID(uint(id))
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
		"food retrieved successfully",
		food,
	))
}

func (h *FoodHandler) GetAllFoods(c *gin.Context) {
	foods, err := h.FoodService.GetAllFoods()
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
		"foods retrieved successfully",
		foods,
	))
}

func (h *FoodHandler) UpdateFood(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, utility.BuildErrorResponse(
			http.StatusBadRequest,
			"invalid food id",
			nil,
		))
		return
	}

	var food models.Food
	if err := c.ShouldBindJSON(&food); err != nil {
		c.JSON(http.StatusBadRequest, utility.BuildErrorResponse(
			http.StatusBadRequest,
			"invalid request body",
			err.Error(),
		))
		return
	}

	updatedFood, err := h.FoodService.UpdateFood(uint(id), &food)
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
		"food updated successfully",
		updatedFood,
	))
}

func (h *FoodHandler) DeleteFood(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, utility.BuildErrorResponse(
			http.StatusBadRequest,
			"invalid food id",
			nil,
		))
		return
	}

	err = h.FoodService.DeleteFood(uint(id))
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
		"food deleted successfully",
		nil,
	))
}
