package route

import (
	"os"
	"unjukketerampilan/controllers"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupRoutes(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.POST("/login", controllers.Login)

	auth := e.Group("")
	auth.Use(echojwt.JWT([]byte(os.Getenv("SECRET_KEY"))))

	auth.POST("/addmerks", controllers.CreateMerk)
	auth.POST("/addcars", controllers.CreateCar)
	auth.GET("/cars", controllers.GetCars)
	auth.GET("/merks", controllers.GetMerks)
	auth.GET("/get-cars-merk", controllers.GetCarsWithMerks)
	auth.GET("/get-cars-merk/:id", controllers.GetCarWithMerkByID)
	auth.DELETE("merk-del/:id", controllers.DeleteMerk)
	auth.DELETE("car-del/:id", controllers.DeleteCar)
}
