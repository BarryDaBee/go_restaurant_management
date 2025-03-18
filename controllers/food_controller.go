package controllers

import (
	"net/http"

	"github.com/BarryDaBee/go_restaurant_management/database"
	"github.com/BarryDaBee/go_restaurant_management/models"
	"github.com/gin-gonic/gin"
)

func CreateFood(ctx *gin.Context) {
	var food models.Food

	if err := ctx.ShouldBindJSON(&food); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := database.DB.Create(&food)

	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Unable to create food"})
		return
	}
	ctx.JSON(http.StatusCreated, food)
}

func GetFoodById(ctx *gin.Context) {
	var food models.Food

	id := ctx.Param("id")

	result := database.DB.First(&food, id)

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Food not found"})
		return
	}

	ctx.JSON(http.StatusOK, food)
}

func GetFoods(ctx *gin.Context) {
	var foods []models.Food

	result := database.DB.Find(&foods)

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Foods not found"})
		return
	}

	ctx.JSON(http.StatusOK, foods)
}

func DeleteFood(ctx *gin.Context) {
	var food models.Food

	id := ctx.Param("id")

	// Find before deleting
	result := database.DB.First(&food, id)

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "food not found"})
		return
	}

	database.DB.Delete(&food)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Food deleted successfully",
	})
}

func UpdateFood(ctx *gin.Context) {
	var food models.Food

	id := ctx.Param("id")

	result := database.DB.First(&food, id)

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Food not found"})
		return
	}

	if err := ctx.ShouldBindJSON(&food); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Save(&food)

	ctx.JSON(http.StatusOK, food)
}
