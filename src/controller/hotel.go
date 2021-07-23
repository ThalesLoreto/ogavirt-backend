package controller

import (
	"net/http"

	model "github.com/ThalesLoreto/ogavirt-backend/model"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

var (
	hotels = []*model.Hotel{}
)

// Create hotel method
func CreateHotel(c echo.Context) error {
	h := &model.Hotel{
		HotelID: uuid.New(),
	}

	if err := c.Bind(h); err != nil {
		return err
	}
	hotels = append(hotels, h)

	return c.JSON(http.StatusOK, h)
}

// Return all created hotels
func GetAllHotel(c echo.Context) error {
	return c.JSON(http.StatusOK, hotels)
}
