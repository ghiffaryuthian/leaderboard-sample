package config

import (
	"log"

	"github.com/spf13/viper"
)

// NewConfig creates new Viper instance from environment and config.yaml file
func NewConfig() *viper.Viper {
	viper.SetConfigFile("config/config.yaml")
	viper.AutomaticEnv()
	config := viper.GetViper()
	if err := config.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err.Error())
	}

	return config
}
