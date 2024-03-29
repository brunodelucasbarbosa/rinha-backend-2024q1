package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type EnvConfig struct {
	DatabaseCredentials DatabaseCredentials
}

func LoadEnvsConfigs() EnvConfig {
	loadEnvs()

	return EnvConfig{
		DatabaseCredentials: DatabaseCredentials{
			Server:   os.Getenv("DB_SERVER"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Port:     os.Getenv("DB_PORT"),
			Database: os.Getenv("DB_DATABASE"),
		},
	}
}

func loadEnvs() {
	err := godotenv.Load(".env")
	if err != nil {
		logrus.Error("Error loading .env file: ", err.Error())
	}
}
