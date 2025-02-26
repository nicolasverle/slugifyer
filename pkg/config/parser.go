package config

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type (
	APIConfig struct {
		Slugs SlugsConfig `json:"slugs"`
		DB    DBConfig    `json:"db"`
	}

	SlugsConfig struct {
		CharsLength int              `json:"chars_length"`
		Storage     StorageConfig    `json:"storage"`
		Validation  ValidationConfig `json:"validation"`
	}

	ValidationConfig struct {
		URLTimeout time.Duration `json:"url_timeout"`
	}

	StorageConfig struct {
		TableName string `json:"table_name"`
	}

	DBConfig struct {
		Engine        string `json:"engine"`
		Port          int    `json:"port"`
		ConnectionURL string `json:"connection_url"`
	}
)

func Parse() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	// config path for pkg unit testing...
	viper.AddConfigPath("../..")

	viper.AutomaticEnv()
	viper.SetEnvPrefix("env")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("error reading config file, %s", err)
	}

	return nil
}
