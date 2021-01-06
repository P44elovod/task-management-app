package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	ServerPort string `mapstructure:"server_port"`
	LogLevel   string `mapstructure:"log_level"`
}

func NewConfig() *Config {
	return &Config{
		ServerPort: ":8080",
		LogLevel:   "debug",
	}
}

func LoadConfig(path string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("toml")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	config := NewConfig()
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, fmt.Errorf("Failed to read the configuration file: %s", err)
	}

	return config, nil

}
