package config

import (
	"fmt"
	"os"
)

type AppEnv struct {
	AppPort             string
	DBHost              string
	DBConnectionString  string
	LoggerUseStructured bool
}

func NewAppEnv() AppEnv {
	appPort := os.Getenv("PORT")
	if appPort == "" {
		appPort = "3000"
	}

	dbHost := os.Getenv("DAN2_TEST_DB_HOST")
	if dbHost == "" {
		dbHost = "localhost"
	}

	connectionString := fmt.Sprintf("postgresql://dev:dev@%s:5400/dan2", dbHost)
	return AppEnv{
		AppPort:            appPort,
		DBHost:             dbHost,
		DBConnectionString: connectionString,
	}
}
