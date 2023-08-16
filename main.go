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
	e := echo.New()
	godotenv.Load()
	e.Start(getPort())

	// Connect to the database
	database.ConnectDB()
	defer database.CloseDB()
	routes.SetupRoutes(e)

	// migration.RunMigrationsForMerks(database.DB)
	// migration.RunMigrationsForCars(database.DB)
	// migration.RunMigrationsForUsers(database.DB)

	e.Start(":8080")
	fmt.Println("Successfully connected to the database")
}

func getPort() string {
	if envPort := os.Getenv("PORT"); envPort != "" {
		return ":" + envPort
	}
	return ":8080"
}
