package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var config *viper.Viper

func Init(path string) {
	var err error
	config = viper.New()

	config.SetConfigFile(path)

	err = config.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("Error: .env file not found")
			os.Exit(1)
		} else {
			fmt.Printf("Error reading .env file: %s\n", err)
			os.Exit(1)
		}
	}
}

func GetConfig() *viper.Viper {
	return config
}
