package config

import (
	"github.com/Template7/common/logger"
	"github.com/spf13/viper"
	"sync"
)

var (
	configPath = "config"
)

type Config struct {
	Backend struct {
		Endpoint string
	}
}

var (
	once     sync.Once
	instance *Config
)

func New() *Config {
	once.Do(func() {
		viper.SetConfigType("yaml")
		instance = &Config{}
		viper.AddConfigPath(configPath)
		viper.SetConfigName("config")
		if err := viper.ReadInConfig(); err != nil {
			panic(err)
		}
		if err := viper.Unmarshal(&instance); err != nil {
			panic(err)
		}

		logger.New().Info("config initialized")
	})
	return instance
}
