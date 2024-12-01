package main

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
)

type Conf struct {
	Port           string `mapstructure:"PORT"`
	JWTSecret      string `mapstructure:"JWT_SECRET"`
	DataSourceName string `mapstructure:"DATA_SOURCE_NAME"`
}

func LoadConfig(path string) (*Conf, error) {
	var cfg *Conf
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.AutomaticEnv()
	err := viper.ReadInConfig()

	if !viper.IsSet("PORT") {
		return nil, fmt.Errorf("PORT is not set in config file")
	}

	if !viper.IsSet("JWT_SECRET") {
		return nil, errors.New("JWT_SECRET is not set in config file")
	}

	if !viper.IsSet("DATA_SOURCE_NAME") {
		return nil, errors.New("DATA_SOURCE_NAME is not set in config file")
	}

	if err != nil {
		return nil, errors.New("error reading config file")
	}

	err = viper.Unmarshal(&cfg)

	if err != nil {
		return nil, errors.New("unable to decode into struct")
	}

	return cfg, nil
}
