package initializers

import (
	"log"

	"github.com/spf13/viper"
)

// Struct to configure the connection to the database
type Config struct {
	DBHost         string `mapstructure:"POSTGRES_HOST"`
	DBUserName     string `mapstructure:"POSTGRES_USER"`
	DBUserPassword string `mapstructure:"POSTGRES_PASSWORD"`
	DBName         string `mapstructure:"POSTGRES_DB"`
	DBPort         string `mapstructure:"POSTGRES_PORT"`
	ServerPort     string `mapstructure:"PORT"`
	ClientOrigin   string `mapstructure:"CLIENT_ORIGIN"`
}

// This will locate the config file on the path (app.env in this case),
// then return the config struct and error
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal("Failed to read the configuration file")
		return
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatal("Failed to unmarshal the environment")
		return
	}
	return
}
