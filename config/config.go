package config

import (
	"log"

	"github.com/spf13/viper"
)

var config *viper.Viper

func Init() {
	var err error
	config = viper.New()

	config.SetConfigFile(".env")

	err = config.ReadInConfig()
	if err != nil {
		log.Fatal("Error on parsing configuration file")
	}
}

func GetConfig() *viper.Viper {
	return config
}
