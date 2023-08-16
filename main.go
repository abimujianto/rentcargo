package main

import (
	"fmt"
	"os"
	"unjukketerampilan/database"
	routes "unjukketerampilan/routes"

	"github.com/labstack/echo/v4"
	// "unjukketerampilan/migration"
)

func main() {
	e := echo.New()
	// godotenv.Load()

	// Connect to the database
	database.ConnectDB()
	defer database.CloseDB()
	routes.SetupRoutes(e)

	e.Start(getPort())
	// migration.RunMigrationsForMerks(database.DB)
	// migration.RunMigrationsForCars(database.DB)
	// migration.RunMigrationsForUsers(database.DB)

	fmt.Println("Successfully connected to the database")
}

func getPort() string {
	if envPort := os.Getenv("PORT"); envPort != "" {
		return ":" + envPort
	}
	return ":8080"
}
