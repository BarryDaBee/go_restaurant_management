package config

type Configuration struct {
	Server    ServerConfiguration
	Databases Databases
	App       App
}

type BaseConfig struct {
	SERVER_PORT string `mapstructure:"SERVER_PORT"`
	JWT_SECRET  string `mapstructure:"JWT_SECRET"`
	APP_ENV     string `mapstructure:"APP_ENV"`

	DB_HOST     string `mapstructure:"DB_HOST"`
	DB_PORT     string `mapstructure:"DB_PORT"`
	DB_SSLMODE  string `mapstructure:"DB_SSLMODE"`
	DB_USER     string `mapstructure:"DB_USER"`
	DB_PASSWORD string `mapstructure:"DB_PASSWORD"`
	DB_NAME     string `mapstructure:"DB_NAME"`
}

func (config *BaseConfig) SetupConfiguration() *Configuration {
	return &Configuration{
		Server: ServerConfiguration{
			Port:   config.SERVER_PORT,
			Secret: config.JWT_SECRET,
		},
		App: App{
			Mode: config.APP_ENV,
		},
		Databases: Databases{
			Host:     config.DB_HOST,
			Port:     config.DB_PORT,
			Username: config.DB_USER,
			Password: config.DB_PASSWORD,
			SSLMode:  config.DB_SSLMODE,
			DBName:   config.DB_NAME,
		},
	}
}
