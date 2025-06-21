package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port string
	Env  string
}

func Load() *Config {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	return &Config{
		Port: viper.GetString("PORT"),
		Env:  viper.GetString("ENV"),
	}
}
