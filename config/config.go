package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

// NewConfig creates new Viper instance from environment and config.yaml file
func NewConfig() *viper.Viper {
	viper.SetConfigFile("config/config.yaml")
	viper.AutomaticEnv()

	// reformat ENV_VAR into env.var for viper
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	config := viper.GetViper()
	if err := config.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err.Error())
	}

	return config
}
