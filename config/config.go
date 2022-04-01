package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	MOVIE_API_KEY string `mapstructure:"MOVIE_API_KEY"`
}

func InitializeViper(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig() // Find and read the config file
	if err != nil {            // Handle errors reading the config file
		fmt.Print("Fatal error config file: %w \n", err)
		return
	}
	err = viper.Unmarshal(&config)
	return
}
