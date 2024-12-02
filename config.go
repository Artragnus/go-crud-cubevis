package main

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
)

type Conf struct {
	Port           string `mapstructure:"PORT"`
	JWTSecret      string `mapstructure:"JWT_SECRET"`
	DBUser         string `mapstructure:"DB_USER"`
	DBPass         string `mapstructure:"DB_PASS"`
	DBPort         string `mapstructure:"DB_PORT"`
	DBName         string `mapstructure:"DB_NAME"`
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

	if !viper.IsSet("DB_USER") {
		return nil, errors.New("DB_USER is not set in config file")
	}

	if !viper.IsSet("DB_PASS") {
		return nil, errors.New("DB_PASS is not set in config file")
	}

	if !viper.IsSet("DB_PORT") {
		return nil, errors.New("DB_PORT is not set in config file")
	}

	if !viper.IsSet("DB_NAME") {
		return nil, errors.New("DB_NAME is not set in config file")
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
