package controllers

import (
	"github.com/gofiber/fiber"
	"github.com/julioolivares90/TumangaOnlineApi/core/tumangaonline"
	"github.com/julioolivares90/TumangaOnlineApi/models"
	"net/http"
)

func GetListasMangas(c *fiber.Ctx) {
	listas := tumangaonline.GetLibraryMangas()

	if listas == nil {
		response := models.Response{
			StatusCode: http.StatusBadRequest,
			Data:       "Ocurrio un error",
		}
		c.JSON(response)
	}
	response := models.Response{
		StatusCode: http.StatusOK,
		Data:       listas,
	}

	c.JSON(response)
}
