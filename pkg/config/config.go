// Package config contains configuration for the application
package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/viper"
)

// Config represents application config.
type Config struct {
	DB struct {
		Host       string `mapstructure:"host"`
		Port       int    `mapstructure:"port"`
		DBUsername string `mapstructure:"db_username"`
		DBPassword string `mapstructure:"db_password"`
	} `mapstructure:"db"`
	Server struct {
		Host string `mapstructure:"host"`
		Port int    `mapstructure:"port"`
	} `mapstructure:"server"`
}

// New loads config from environment variables and viper config file. Returns error if config is not valid.
func New(folder, filename string) (*Config, error) {
	cfg := &Config{}

	viper.AddConfigPath(folder)
	viper.SetConfigName(filename)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(cfg); err != nil {
		return nil, err
	}

	if err := envconfig.Process("db", cfg); err != nil {
		return nil, err
	}

	if err := envconfig.Process("server", cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
