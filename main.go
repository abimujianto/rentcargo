package main

import (
	"fmt"
	"os"
	"unjukketerampilan/database"
	routes "unjukketerampilan/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	// "unjukketerampilan/migration"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if PORT environment variable is not set
	}

	// Connect to the database
	database.ConnectDB()
	defer database.CloseDB()

	e := echo.New()

	routes.SetupRoutes(e)

	// migration.RunMigrationsForMerks(database.DB)
	// migration.RunMigrationsForCars(database.DB)
	// migration.RunMigrationsForUsers(database.DB)

	e.Start(":8080")
	fmt.Println("Successfully connected to the database")
}
