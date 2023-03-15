package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Configs struct {
	MapboxToken string
}

func init() {
	godotenv.Load()
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
