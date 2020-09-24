package config

import (
	"log"
	"sync"

	"github.com/sirius1024/overseer/models"
	"github.com/spf13/viper"
)

var (
	once sync.Once
	conf models.Configuration
)

// GetConfig from config file or env vars
func GetConfig() models.Configuration {
	once.Do(func() {
		conf = initializeConfiguration()
	})
	return conf
}

// initializeConfiguration 读取本地配置
func initializeConfiguration() models.Configuration {
	viper.SetConfigName("overseer")
	viper.AddConfigPath("/etc/overseer/")
	viper.AddConfigPath("$HOME/.overseer")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// TODO: Read from env
		} else {
			// other error
		}
	}
	// config parsed
	var conf models.Configuration
	err := viper.Unmarshal(&conf)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	return conf
}
