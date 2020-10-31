package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/julioolivares90/TumangaOnlineApi/core/tumangaonline"
	"github.com/julioolivares90/TumangaOnlineApi/models"
	"net/http"
)

func GetListasMangas(c *fiber.Ctx) error {
	listas := tumangaonline.GetLibraryMangas()

	if listas == nil {
		response := models.Response{
			StatusCode: http.StatusBadRequest,
			Data:       "Ocurrio un error",
		}
		return c.JSON(response)
	}
	response := models.Response{
		StatusCode: http.StatusOK,
		Data:       listas,
	}

	return c.JSON(response)
}
