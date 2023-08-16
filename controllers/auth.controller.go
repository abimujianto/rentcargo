package controllers

import (
	"net/http"
	"os"
	"time"
	"unjukketerampilan/models"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}

func Login(c echo.Context) error {
	var loginRequest models.User

	c.Bind(&loginRequest)

	// Set custom claims
	claims := &jwtCustomClaims{
		"Abi",
		true,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Success: true,
		Message: "Berhasil",
		Data: map[string]string{
			"token": t,
		},
	})
}
