package main

import (
	"fmt"
	"unjukketerampilan/database"
	routes "unjukketerampilan/routes"

	"github.com/labstack/echo/v4"
	// "unjukketerampilan/migration"
)

func main() {
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
