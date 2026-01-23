package postgresql

import (
	"fmt"
	"log"

	"github.com/BarryDaBee/go_restaurant_management/internal/config"
	"github.com/BarryDaBee/go_restaurant_management/utility"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresStore struct {
	DB            *gorm.DB
	Configuration *config.Configuration
}

func NewPostgresStore(configuration *config.Configuration) *PostgresStore {
	return &PostgresStore{
		Configuration: configuration,
	}
}

func (p *PostgresStore) Connect(logger *utility.Logger, configDatabases config.Databases) {
	var dsn string

	if configDatabases.Host != "" {
		dsn = fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
			configDatabases.Host,
			configDatabases.Username,
			configDatabases.Password,
			configDatabases.DBName,
			configDatabases.Port,
			configDatabases.SSLMode,
		)
	} else {
		log.Fatal("Database configuration is missing")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	p.DB = db
	utility.LogAndPrint(logger, "Database connection established")
}

func (p *PostgresStore) GetDB() interface{} {
	return p.DB
}

func (p *PostgresStore) Migrate(models ...interface{}) error {
	if len(models) == 0 {
		return nil
	}
	return p.DB.AutoMigrate(models...)
}

func (p *PostgresStore) Create(model interface{}) error {
	return p.DB.Create(model).Error
}

func (p *PostgresStore) FindByID(model interface{}, id uint) error {
	return p.DB.First(model, id).Error
}

func (p *PostgresStore) FindOne(model interface{}, query interface{}, args ...interface{}) error {
	return p.DB.Where(query, args...).First(model).Error
}

func (p *PostgresStore) FindAll(models interface{}) error {
	return p.DB.Find(models).Error
}

func (p *PostgresStore) FindAllWithCondition(models interface{}, query interface{}, args ...interface{}) error {
	return p.DB.Where(query, args...).Find(models).Error
}

func (p *PostgresStore) Update(model interface{}) error {
	return p.DB.Save(model).Error
}

func (p *PostgresStore) Delete(model interface{}) error {
	return p.DB.Delete(model).Error
}

func (p *PostgresStore) CheckExists(model interface{}, query interface{}, args ...interface{}) bool {
	var count int64
	p.DB.Model(model).Where(query, args...).Count(&count)
	return count > 0
}
