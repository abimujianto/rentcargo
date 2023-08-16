package controllers

import (
	"net/http"
	"unjukketerampilan/database"
	"unjukketerampilan/models"

	"github.com/labstack/echo/v4"
)

func CreateCar(c echo.Context) error {

	db := database.DB

	car := new(models.Car)
	if err := c.Bind(car); err != nil {
		response := models.BaseResponse{
			Success: false,
			Message: "Gagal request data",
			Data:    nil,
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	if err := db.Create(car).Error; err != nil {
		response := models.BaseResponse{
			Success: false,
			Message: "Gagal tambah data",
			Data:    nil,
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := models.BaseResponse{
		Success: false,
		Message: "Berhasil tambah data mobil",
		Data:    car,
	}
	return c.JSON(http.StatusCreated, response)
}

type Cars struct {
	ID          int     `gorm:"primaryKey:autoIncrement" json:"id"`
	IDMerk      int     `gorm:"foreignKey:IDMerk" json:"id_merk"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

func GetCars(c echo.Context) error {

	db := database.DB

	var cars []Cars

	if err := db.Select("id, name, description, price").Find(&cars).Error; err != nil {
		response := models.BaseResponse{
			Success: false,
			Message: "Gagal mengambil data mobil",
			Data:    nil,
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := models.BaseResponse{
		Success: true,
		Message: "Berhasil mengambil data mobil",
		Data:    cars,
	}
	return c.JSON(http.StatusOK, response)
}

func GetCarsWithMerks(c echo.Context) error {
	db := database.DB

	var cars []models.Car
	if err := db.Preload("Merk").Find(&cars).Error; err != nil {
		response := models.BaseResponse{
			Success: false,
			Message: "Gagal mengambil data mobil",
			Data:    nil,
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := models.BaseResponse{
		Success: true,
		Message: "Berhasil mengambil data mobil",
		Data:    cars,
	}
	return c.JSON(http.StatusOK, response)
}

func GetCarWithMerkByID(c echo.Context) error {
	db := database.DB

	carID := c.Param("id")

	var car models.Car
	if err := db.Joins("JOIN merks ON cars.id_merk = merks.id").Where("cars.id = ?", carID).First(&car).Error; err != nil {
		response := models.BaseResponse{
			Success: false,
			Message: "Gagal mengambil data mobil",
			Data:    nil,
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := models.BaseResponse{
		Success: true,
		Message: "Berhasil mengambil data mobil dengan relasi merek berdasarkan id",
		Data:    car,
	}
	return c.JSON(http.StatusOK, response)
}

func DeleteCar(c echo.Context) error {
	db := database.DB

	carID := c.Param("id")

	var car models.Car

	if err := db.First(&car, carID).Error; err != nil {
		response := models.BaseResponse{
			Success: false,
			Message: "Mobil tidak ditemukan !",
			Data:    nil,
		}
		return c.JSON(http.StatusNotFound, response)
	}

	if err := db.Delete(&car).Error; err != nil {
		response := models.BaseResponse{
			Success: false,
			Message: "Gagal hapus data !",
			Data:    nil,
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := models.BaseResponse{
		Success: true,
		Message: "Berhasil menghapus data mobil",
		Data:    car,
	}

	return c.JSON(http.StatusOK, response)
}
