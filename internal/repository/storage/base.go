package storage

import (
	"strconv"

	"github.com/BarryDaBee/go_restaurant_management/internal/config"
	"github.com/BarryDaBee/go_restaurant_management/utility"
	"github.com/gin-gonic/gin"
)

type DeleteType string

var (
	SoftDelete DeleteType = "soft"
	HardDelete DeleteType = "hard"
)

type Storage interface {
	Connect(logger *utility.Logger, configDatabases config.Databases)
	GetDB() interface{}
	Migrate(models ...interface{}) error

	Create(model interface{}) error
	FindByID(model interface{}, id uint) error
	FindOne(model interface{}, query interface{}, args ...interface{}) error
	FindAll(models interface{}) error
	FindAllWithCondition(models interface{}, query interface{}, args ...interface{}) error
	Update(model interface{}) error
	Delete(model interface{}) error
	CheckExists(model interface{}, query interface{}, args ...interface{}) bool
}

var (
	defaultPage  = 1
	defaultLimit = 20
)

type Pagination struct {
	Page  int
	Limit int
}

type PaginationResponse struct {
	CurrentPage int `json:"current_page"`
	Page        int `json:"page"`
	TotalPages  int `json:"total_pages"`
}

func GetPagination(c *gin.Context) Pagination {
	var (
		page  *int
		limit *int
	)
	if c.Query("page") != "" {
		pageInt, err := strconv.Atoi(c.Query("page"))
		if err == nil {
			page = &pageInt
		}
	}
	if c.Query("limit") != "" {
		limitInt, err := strconv.Atoi(c.Query("limit"))
		if err == nil {
			limit = &limitInt
		}
	}

	if page != nil && limit != nil {
		return Pagination{Page: *page, Limit: *limit}
	} else if page == nil && limit != nil {
		return Pagination{Page: defaultPage, Limit: *limit}
	} else if page != nil && limit == nil {
		return Pagination{Page: *page, Limit: defaultLimit}
	} else {
		return Pagination{Page: defaultPage, Limit: defaultLimit}
	}
}
