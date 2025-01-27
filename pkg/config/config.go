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
	SMTP     SMTPConfig     `mapstructure:"smtp"`
	Template TemplateConfig `mapstructure:"template"`
	Location LocationConfig `mapstructure:"location"`
}

type AppConfig struct {
	ActivationURL     string `mapstructure:"activation_url"`
	AuthenticationURL string `mapstructure:"authentication_url"`
	RegistrationURL   string `mapstructure:"registration_url"`
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
	Secret                            string `mapstructure:"secret"`
	TokenExpirationTime               int    `mapstructure:"token_expiration_time"`
	RefreshTokenExpirationTime        int    `mapstructure:"refresh_token_expiration_time"`
	ActivationTokenExpirationTime     int    `mapstructure:"activation_token_expiration_time"`
	AuthenticationTokenExpirationTime int    `mapstructure:"authentication_token_expiration_time"`
	RegistrationTokenExpirationTime   int    `mapstructure:"registration_token_expiration_time"`
}

type TemplateConfig struct {
	EmailActivation     string `mapstructure:"email_activation"`
	EmailResetPassword  string `mapstructure:"email_reset_password"`
	EmailAuthentication string `mapstructure:"email_authentication"`
	EmailRegistration   string `mapstructure:"email_registration"`
}

type SMTPConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Email    string `mapstructure:"email"`
	Password string `mapstructure:"password"`
}

type LocationConfig struct {
	VietnamProvinceAPI string `mapstructure:"vietnam_province_api"`
	OSMURL             string `mapstructure:"osm_url"`
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
