package config

import "github.com/spf13/viper"

type Config struct {
	DBHost     string `mapstructure:"DB_HOST"`
	DBName     string `mapstructure:"DB_NAME"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBPort     string `mapstructure:"DB_PORT"`
	ACCOUNTSID string `mapstructure:"DB_ACCOUNTSID"`
	SERVICEID  string `mapstructure:"DB_SERVICEID"`
	AUTHTOKEN  string `mapstructure:"DB_AUTHTOKEN"`
	UNIDOCKEY  string `mapstructure:"UNIDOC_LICENSE_API_KEY"`
}

var envs = []string{
	"DB_HOST",
	"DB_NAME",
	"DB_USER",
	"DB_PASSWORD",
	"DB_PORT",
	"DB_ACCOUNTSID",
	"DB_SERVICEID",
	"DB_AUTHTOKEN",
	"UNIDOC_LICENSE_API_KEY",
}