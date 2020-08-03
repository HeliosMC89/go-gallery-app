package main

import (
	"os"
	// load dotenv variables
	_ "github.com/joho/godotenv/autoload"
	// load migrations
	_ "github.com/heliosmc89/gallery-app-with-go/migrations"
	"github.com/heliosmc89/gallery-app-with-go/routes"
)

func main() {
	fmt.println("Server started on: %s", os.Getenv("APP_PORT"))
	routes.RegisterRoutes()
}
