package controllers

import (
	"net/http"
	"unjukketerampilan/database"
	"unjukketerampilan/models"

	"github.com/labstack/echo/v4"
)

func CreateMerk(c echo.Context) error {
	db := database.DB

	merk := new(models.Merk)
	if err := c.Bind(merk); err != nil {
		response := models.BaseResponse{
			Success: false,
			Message: "Gagal request data",
			Data:    nil,
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	//untuk post data
	if err := db.Create(merk).Error; err != nil {
		response := models.BaseResponse{
			Success: false,
			Message: "Gagal tambah data",
			Data:    nil,
		}
		c.JSON(http.StatusInternalServerError, response)
		return err
	}

	response := models.BaseResponse{
		Success: true,
		Message: "Berhasil tambah data merek",
		Data:    merk,
	}
	return c.JSON(http.StatusCreated, response)
}

func GetMerks(c echo.Context) error {
	db := database.DB

	var merks []models.Merk
	if err := db.Find(&merks).Error; err != nil {
		response := models.BaseResponse{
			Success: false,
			Message: "Gagal mengambil data merek",
			Data:    nil,
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := models.BaseResponse{
		Success: true,
		Message: "Berhasil mengambil data merek",
		Data:    merks,
	}
	return c.JSON(http.StatusOK, response)
}

func DeleteMerk(c echo.Context) error {

	db := database.DB

	merkId := c.Param("id")

	var merk models.Merk

	if err := db.First(&merk, merkId).Error; err != nil {
		response := models.BaseResponse{
			Success: false,
			Message: "Merk tidak ditemukan !",
			Data:    nil,
		}
		return c.JSON(http.StatusNotFound, response)
	}

	if err := db.Delete(&merk).Error; err != nil {
		response := models.BaseResponse{
			Success: false,
			Message: "Gagal Hapus data",
			Data:    nil,
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := models.BaseResponse{
		Success: true,
		Message: "Berhasil menghapus data merk",
		Data:    merk,
	}

	return c.JSON(http.StatusOK, response)
}
