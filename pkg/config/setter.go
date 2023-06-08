package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func Set() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// config file not found; ignore error if desired
		} else {
			// config file was found but another error was produced
		}
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("unable to decode into struct. %v", err)
	}
}
