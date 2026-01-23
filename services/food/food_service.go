package food

import (
	"net/http"

	"github.com/BarryDaBee/go_restaurant_management/internal/repository/storage"
	"github.com/BarryDaBee/go_restaurant_management/internal/models"
	"github.com/BarryDaBee/go_restaurant_management/utility"
)

type foodService struct {
	Storage storage.Storage
}

func NewFoodService(storage storage.Storage) FoodService {
	return &foodService{
		Storage: storage,
	}
}

type FoodService interface {
	CreateFood(food *models.Food) (*models.Food, error)
	GetFoodByID(id uint) (*models.Food, error)
	GetAllFoods() ([]models.Food, error)
	UpdateFood(id uint, food *models.Food) (*models.Food, error)
	DeleteFood(id uint) error
}

func (f *foodService) CreateFood(food *models.Food) (*models.Food, error) {
	err := f.Storage.Create(food)
	if err != nil {
		return nil, utility.NewAppError("error creating food", http.StatusBadRequest)
	}
	return food, nil
}

func (f *foodService) GetFoodByID(id uint) (*models.Food, error) {
	var food models.Food
	err := f.Storage.FindByID(&food, id)
	if err != nil {
		return nil, utility.NewAppError("food not found", http.StatusNotFound)
	}
	return &food, nil
}

func (f *foodService) GetAllFoods() ([]models.Food, error) {
	var foods []models.Food
	err := f.Storage.FindAll(&foods)
	if err != nil {
		return nil, utility.NewAppError("error fetching foods", http.StatusInternalServerError)
	}
	return foods, nil
}

func (f *foodService) UpdateFood(id uint, food *models.Food) (*models.Food, error) {
	var existingFood models.Food
	err := f.Storage.FindByID(&existingFood, id)
	if err != nil {
		return nil, utility.NewAppError("food not found", http.StatusNotFound)
	}

	existingFood.Name = food.Name
	existingFood.Price = food.Price
	existingFood.Category = food.Category

	err = f.Storage.Update(&existingFood)
	if err != nil {
		return nil, utility.NewAppError("error updating food", http.StatusInternalServerError)
	}

	return &existingFood, nil
}

func (f *foodService) DeleteFood(id uint) error {
	var food models.Food
	err := f.Storage.FindByID(&food, id)
	if err != nil {
		return utility.NewAppError("food not found", http.StatusNotFound)
	}

	err = f.Storage.Delete(&food)
	if err != nil {
		return utility.NewAppError("error deleting food", http.StatusInternalServerError)
	}

	return nil
}
