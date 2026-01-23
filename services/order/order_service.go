package order

import (
	"net/http"

	"github.com/BarryDaBee/go_restaurant_management/internal/repository/storage"
	"github.com/BarryDaBee/go_restaurant_management/internal/models"
	"github.com/BarryDaBee/go_restaurant_management/utility"
)

type orderService struct {
	Storage storage.Storage
}

func NewOrderService(storage storage.Storage) OrderService {
	return &orderService{
		Storage: storage,
	}
}

type OrderService interface {
	CreateOrder(order *models.Order) (*models.Order, error)
	GetOrderByID(id uint) (*models.Order, error)
	GetOrdersByUserID(userID uint) ([]models.Order, error)
	GetAllOrders() ([]models.Order, error)
	UpdateOrderStatus(id uint, status string) (*models.Order, error)
	DeleteOrder(id uint) error
}

func (o *orderService) CreateOrder(order *models.Order) (*models.Order, error) {
	err := o.Storage.Create(order)
	if err != nil {
		return nil, utility.NewAppError("error creating order", http.StatusBadRequest)
	}
	return order, nil
}

func (o *orderService) GetOrderByID(id uint) (*models.Order, error) {
	var order models.Order
	err := o.Storage.FindByID(&order, id)
	if err != nil {
		return nil, utility.NewAppError("order not found", http.StatusNotFound)
	}
	return &order, nil
}

func (o *orderService) GetOrdersByUserID(userID uint) ([]models.Order, error) {
	var orders []models.Order
	err := o.Storage.FindAllWithCondition(&orders, "user_id = ?", userID)
	if err != nil {
		return nil, utility.NewAppError("error fetching orders", http.StatusInternalServerError)
	}
	return orders, nil
}

func (o *orderService) GetAllOrders() ([]models.Order, error) {
	var orders []models.Order
	err := o.Storage.FindAll(&orders)
	if err != nil {
		return nil, utility.NewAppError("error fetching orders", http.StatusInternalServerError)
	}
	return orders, nil
}

func (o *orderService) UpdateOrderStatus(id uint, status string) (*models.Order, error) {
	var order models.Order
	err := o.Storage.FindByID(&order, id)
	if err != nil {
		return nil, utility.NewAppError("order not found", http.StatusNotFound)
	}

	order.Status = status
	err = o.Storage.Update(&order)
	if err != nil {
		return nil, utility.NewAppError("error updating order", http.StatusInternalServerError)
	}

	return &order, nil
}

func (o *orderService) DeleteOrder(id uint) error {
	var order models.Order
	err := o.Storage.FindByID(&order, id)
	if err != nil {
		return utility.NewAppError("order not found", http.StatusNotFound)
	}

	err = o.Storage.Delete(&order)
	if err != nil {
		return utility.NewAppError("error deleting order", http.StatusInternalServerError)
	}

	return nil
}
