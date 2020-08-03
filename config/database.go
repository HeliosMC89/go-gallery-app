package config

import (
	"os"
)

var db = map[string]string{
	"port":       os.Getenv("DB_PORT"),
	"database":   os.Getenv("DB_DATABASE"),
	"host":       os.Getenv("DB_HOST"),
	"password":   os.Getenv("DB_PASSWORD"),
	"username":   os.Getenv("DB_USERNAME"),
	"connection": os.Getenv("DB_CONNECTION"),
}

func GetDatabase() map[string]string {
	return db
}
