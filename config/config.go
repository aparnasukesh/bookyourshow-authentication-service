package config

import (
	"fmt"

	"github.com/go-playground/validator"
	"github.com/spf13/viper"
)

type Config struct {
	GrpcPort       string `mapstructure:"GRPCPORT" validate:"required"`
	JWT_secret_key string `mapstructure:"JWT_secret_key" validate:"required"`
}

var envs = []string{
	"GRPCPORT", "JWT_secret_key",
}

func LoadConfig() (Config, error) {
	var cfg Config
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return cfg, fmt.Errorf("error reading config file: %w", err)
	}

	for _, env := range envs {
		if err := viper.BindEnv(env); err != nil {
			return cfg, fmt.Errorf("error binding environment variable %s: %w", env, err)
		}
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return cfg, fmt.Errorf("error unmarshalling config: %w", err)
	}

	validate := validator.New()
	if err := validate.Struct(&cfg); err != nil {
		return cfg, fmt.Errorf("validation error: %w", err)
	}

	return cfg, nil
}
