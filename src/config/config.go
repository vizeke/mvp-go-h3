package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Configs struct {
	MapboxToken string
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err.Error())
	}
}

func GetConfigs() Configs {
	return Configs{
		MapboxToken: GetMapboxToken(),
	}
}

func GetDbConnection() string {
	return getValue("DB_CONN_STRING")
}

func GetMapboxToken() string {
	return getValue("MAPBOX_TOKEN")
}

func getValue(key string) string {
	return os.Getenv(key)
}
