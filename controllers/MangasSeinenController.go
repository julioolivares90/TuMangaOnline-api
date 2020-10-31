package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/julioolivares90/TumangaOnlineApi/core/tumangaonline"
	"github.com/julioolivares90/TumangaOnlineApi/models"
	"net/http"
)

func GetMangasPopularesSeinen(c *fiber.Ctx) error {
	mangas := tumangaonline.GetMangasPopularesSeinen()
	response := models.Response{
		StatusCode: http.StatusOK,
		Data:       mangas,
	}

	if len(mangas) > 0 {
		return c.JSON(response)
	} else {
		errorResponse := models.Response{
			StatusCode: http.StatusBadGateway,
			Data:       "no se encontraron mangas",
		}
		return c.JSON(errorResponse)
	}
}
