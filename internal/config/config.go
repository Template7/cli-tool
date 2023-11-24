package config

import (
	"fmt"
	"github.com/Template7/common/logger"
	"github.com/spf13/viper"
	"sync"
)

var (
	configPath = "configs"
	log = logger.GetLogger()
)

type config struct {
	JwtSign string
	Log struct {
		Level     string `yaml:"level"`
		Formatter string
	} `yaml:"log"`
	Mongo struct {
		Db               string `yaml:"db"`
		Host             string `yaml:"host"`
		Port             int    `yaml:"port"`
		Username         string `yaml:"username"`
		Password         string `yaml:"password"`
		ConnectionString string
	} `yaml:"mongo"`
	MySql struct {
		Db               string
		Host             string
		Port             int
		Username         string
		Password         string
		ConnectionString string
		Root             string
		RootPassword     string
	}
	Backend struct {
		Endpoint string
		Username string
		Password string
	}
	Redis struct {
		Host     string
		Password string
		//PollSize int
		//ReadTimeout int
	}
}

var (
	once     sync.Once
	instance *config
)

func New() *config {
	once.Do(func() {
		viper.SetConfigType("yaml")
		instance = &config{}
		viper.AddConfigPath(configPath)
		viper.SetConfigName("config")
		if err := viper.ReadInConfig(); err != nil {
			log.Fatal("fail to load config file: ", err.Error())
		}
		if err := viper.Unmarshal(&instance); err != nil {
			log.Fatal(err)
		}
		instance.initLog()

		if instance.Mongo.Username != "" && instance.Mongo.Password != "" {
			instance.Mongo.ConnectionString = fmt.Sprintf("mongodb://%s:%s@%s:%d", instance.Mongo.Username, instance.Mongo.Password, instance.Mongo.Host, instance.Mongo.Port)
		} else {
			instance.Mongo.ConnectionString = fmt.Sprintf("mongodb://%s:%d", instance.Mongo.Host, instance.Mongo.Port)
		}
		instance.MySql.ConnectionString = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", instance.MySql.Username, instance.MySql.Password, instance.MySql.Host, instance.MySql.Port, instance.MySql.Db)

		log.Debug("config initialized")
	})
	return instance
}

func SetDbConnectionString() {
	if instance.Mongo.Username != "" && instance.Mongo.Password != "" {
		instance.Mongo.ConnectionString = fmt.Sprintf("mongodb://%s:%s@%s:%d", instance.Mongo.Username, instance.Mongo.Password, instance.Mongo.Host, instance.Mongo.Port)
	} else {
		instance.Mongo.ConnectionString = fmt.Sprintf("mongodb://%s:%d", instance.Mongo.Host, instance.Mongo.Port)
	}
	//instance.MySql.ConnectionString = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", instance.MySql.Username, instance.MySql.Password, instance.MySql.Host, instance.MySql.Port, instance.MySql.Db)
	instance.MySql.ConnectionString = fmt.Sprintf("%s:%s@tcp(%s:%d)/?charset=utf8mb4&parseTime=True&loc=Local", instance.MySql.Root, instance.MySql.RootPassword, instance.MySql.Host, instance.MySql.Port)
}

func (c *config) initLog() {
	logger.SetLevel(c.Log.Level)
	logger.SetFormatter(c.Log.Formatter)
	log.Debug("logger initialized")
}
