package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type AppMode int

const (
	Development AppMode = iota
	Testing
	Production
)

type Config struct {
	AppMode    AppMode
	PGUrl      string
	HttpServer HTTPServerCfg
}

type HTTPServerCfg struct {
	Port int
}

func Load() *Config {
	var config Config
	mode := os.Getenv("MODE")

	var configFile string
	switch mode {
	case "production":
		config.AppMode = Production
		configFile = "config.prod"
	case "testing":
		config.AppMode = Testing
		configFile = "config.test"
	default:
		config.AppMode = Development
		configFile = "config.dev"
	}

	viper.AddConfigPath("configs")
	viper.SetConfigFile(configFile)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("faied read config file, reason: %v", err)
	}

	return &config
}
