package migrations

import (
	"github.com/heliosmc89/gallery-app-with-go/models"
	"github.com/heliosmc89/gallery-app-with-go/utils"
)

var database, err = utils.GetDatabaseConnection()

func init() {
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(models.User{})
}

//Refresh function is used to take the tables down form the database and refresh it
func Refresh() {
	database.DropTableIfExists(&models.User{})
	database.AutoMigrate(&models.User{})
}
