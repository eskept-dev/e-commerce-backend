package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	App      AppConfig      `mapstructure:"app"`
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	Cache    CacheConfig    `mapstructure:"cache"`
	JWT      JWTConfig      `mapstructure:"jwt"`
}

type AppConfig struct {
	ActivationURL string `mapstructure:"activation_url"`
}

type ServerConfig struct {
	Port int `mapstructure:"port"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"db_name"`
	SSLMode  string `mapstructure:"ssl_mode"`
}

type CacheConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
}

type JWTConfig struct {
	Secret                        string `mapstructure:"secret"`
	TokenExpirationTime           int    `mapstructure:"token_expiration_time"`
	RefreshTokenExpirationTime    int    `mapstructure:"refresh_token_expiration_time"`
	ActivationTokenExpirationTime int    `mapstructure:"activation_token_expiration_time"`
}

func LoadConfig(path string) (*Config, error) {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	viper.AddConfigPath(path)
	viper.SetConfigName(fmt.Sprintf("app.%s", env))
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	config := &Config{}
	err = viper.Unmarshal(config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return config, nil
}
