package config

import (
	"os"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func GetDbConnection() string {
	return getValue("DB_CONN_STRING")
}

func getValue(key string) string {
	return os.Getenv(key)
}
