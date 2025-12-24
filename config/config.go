package config

import (
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	BaseURL      string `mapstructure:"base_url"`
	ClientID     string `mapstructure:"client_id"`
	ClientSecret string `mapstructure:"client_secret"`
	Debug        bool   `mapstructure:"debug"`
}

func Load() (*Config, error) {
	viper.SetEnvPrefix("VERSAFLEET")
	viper.AutomaticEnv()

	// Replace . with _ in env vars if needed, though SetEnvPrefix usually matches prefix_upper_key
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Set defaults
	viper.SetDefault("base_url", "https://api.versafleet.co/api")
	viper.SetDefault("debug", false)

	// Allow reading from a .env file if present
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	_ = viper.ReadInConfig() // Ignore error if config file not found

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
