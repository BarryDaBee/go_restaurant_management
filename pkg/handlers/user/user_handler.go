package user

import (
	"net/http"

	"github.com/BarryDaBee/go_restaurant_management/internal/models"
	userService "github.com/BarryDaBee/go_restaurant_management/services/user"
	"github.com/BarryDaBee/go_restaurant_management/utility"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	UserService userService.UserService
	Validate    *validator.Validate
}

func NewUserHandler(service userService.UserService, validate *validator.Validate) *UserHandler {
	return &UserHandler{
		UserService: service,
		Validate:    validate,
	}
}

type RegisterRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (h *UserHandler) Register(c *gin.Context) {
	var req RegisterRequest

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

	user := &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	createdUser, err := h.UserService.Register(user)
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
		"user registered successfully",
		createdUser,
	))
}

func (h *UserHandler) Login(c *gin.Context) {
	var req LoginRequest

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

	token, err := h.UserService.Login(req.Email, req.Password)
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
		"login successful",
		gin.H{"token": token},
	))
}
