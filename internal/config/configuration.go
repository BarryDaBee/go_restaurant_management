package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

var Config *Configuration

func Setup(name string) *Configuration {
	var baseConfiguration *BaseConfig

	appEnv := os.Getenv("APP_ENV")

	switch appEnv {
	case "local", "":
		viper.SetConfigFile(".env")
		if err := viper.ReadInConfig(); err != nil {
			log.Printf("Warning: Error reading .env file: %v", err)
		}
		log.Println("Loaded configurations from .env (local mode)")

	case "production":
		viper.AutomaticEnv()
		log.Println("Loaded configurations from environment variables (production mode)")

	default:
		viper.SetConfigFile(".env")
		if err := viper.ReadInConfig(); err != nil {
			log.Printf("Warning: Error reading .env file: %v", err)
		}
	}

	viper.AutomaticEnv()

	if err := viper.Unmarshal(&baseConfiguration); err != nil {
		log.Fatalf("Unable to decode configuration into struct: %v", err)
	}

	configuration := baseConfiguration.SetupConfiguration()
	Config = configuration
	log.Println("Configurations loaded successfully")
	return configuration
}

func GetConfig() *Configuration {
	return Config
}
