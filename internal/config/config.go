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

	DBHostReplica1     string `mapstructure:"DB_HOST_REPLICA_1"`
	DBPortReplica1     uint   `mapstructure:"DB_PORT_REPLICA_1"`
	DBUsernameReplica1 string `mapstructure:"DB_USERNAME_REPLICA_1"`
	DBPasswordReplica1 string `mapstructure:"DB_PASSWORD_REPLICA_1"`
	DBNameReplica1     string `mapstructure:"DB_NAME_REPLICA_1"`
	DBSSLModeReplica1  string `mapstructure:"DB_SSLMODE_REPLICA_1"`

	DBHostReplica2     string `mapstructure:"DB_HOST_REPLICA_2"`
	DBPortReplica2     uint   `mapstructure:"DB_PORT_REPLICA_2"`
	DBUsernameReplica2 string `mapstructure:"DB_USERNAME_REPLICA_2"`
	DBPasswordReplica2 string `mapstructure:"DB_PASSWORD_REPLICA_2"`
	DBNameReplica2     string `mapstructure:"DB_NAME_REPLICA_2"`
	DBSSLModeReplica2  string `mapstructure:"DB_SSLMODE_REPLICA_2"`

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
