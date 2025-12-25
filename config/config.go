package config

import (
	"fmt"
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
	err := viper.ReadInConfig() // Ignore error if config file not found
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("No config file found (using env/defaults only)")
		} else {
			fmt.Printf("Config file found but error occurred: %v\n", err)
		}
	} else {
		fmt.Println("Config file successfully loaded:", viper.ConfigFileUsed())
	}

	// settings := viper.AllSettings()
	// out, _ := json.MarshalIndent(settings, "", "  ")
	// fmt.Printf("--- VIPER DEBUG ---\n%s\n-------------------\n", string(out))

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	// fmt.Println(cfg)

	return &cfg, nil
}
