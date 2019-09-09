package common

import (
	"github.com/spf13/viper"
	"log"
	"sync"
)

var once sync.Once

type Config struct {
}

func GetConfig() *Config {
	once.Do(func() {
		viper.SetConfigFile("./config.toml")
		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("Error reading config file, %s", err)
		}
	})
	return &Config{}
}

func (*Config) GetToken() string {
	return viper.GetString("telegram.token")
}
