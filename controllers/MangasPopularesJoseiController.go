package controllers

import (
	"github.com/gofiber/fiber"
	"github.com/julioolivares90/TumangaOnlineApi/core/tumangaonline"
	"github.com/julioolivares90/TumangaOnlineApi/models"
	"net/http"
)

func GetMangasPopularesJosei(c *fiber.Ctx) {
	mangas := tumangaonline.GetMangasPopularesJosei()
	response := models.Response{
		StatusCode: http.StatusOK,
		Data:       mangas,
	}

	if len(mangas) > 0 {
		c.JSON(response)
	} else {
		errorResponse := models.Response{
			StatusCode: http.StatusBadGateway,
			Data:       "no se encontraron mangas",
		}
		c.JSON(errorResponse)
	}
}
