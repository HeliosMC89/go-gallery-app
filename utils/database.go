package utils

import (
	"fmt"
	// setting up gorm with postgres connection
	"github.com/heliosmc89/gallery-app-with-go/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

var databaseCredentials = config.GetDatabase()
var psqlInfo = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	databaseCredentials["host"],
	databaseCredentials["port"],
	databaseCredentials["username"],
	databaseCredentials["password"],
	databaseCredentials["database"],
)

var db *gorm.DB

// GetDatabaseConnection function to return the current database connection.
func GetDatabaseConnection() (*gorm.DB, error) {
	var err error
	db, err = gorm.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	db.LogMode(true)
	return db, nil
}
