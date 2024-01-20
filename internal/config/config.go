package config

import (
	"github.com/spf13/viper"
)

func init() {
	initConfig()
}

func initConfig() {
	viper.SetDefault("gitBinary", "git") // Default git binary path
	viper.SetEnvPrefix("TEMPLATE_CLI")   // Prefix for environment variables
	viper.AutomaticEnv()                 // Read from environment variables
}
