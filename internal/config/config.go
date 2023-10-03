package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     uint   `mapstructure:"DB_PORT"`
	DBUsername string `mapstructure:"DB_USERNAME"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
	DBSSLMode  string `mapstructure:"DB_SSLMODE"`
	ServerPort string `mapstructure:"SERVER_PORT"`
	Salt       string `mapstructure:"APP_SALT"`
	SigningKey string `mapstructure:"APP_SIGNING_KEY"`
	TokenTTL   uint   `mapstructure:"AUTH_TOKEN_TTL"`
}

func NewConfig() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()

	if err != nil {
		return nil, err
	}

	var config Config

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}
	logrus.Debug("Parsed config values")
	logrus.Debugln(config)

	return &config, nil
}
