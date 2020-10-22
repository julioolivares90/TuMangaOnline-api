package controllers

import (
	"github.com/gofiber/fiber"
	"github.com/julioolivares90/TumangaOnlineApi/core/tumangaonline"
	"github.com/julioolivares90/TumangaOnlineApi/models"
	"net/http"
)

func GetInfoLibrary(c *fiber.Ctx) {
	title := c.Query("title")
	_page := c.Query("_page")
	orderItem := c.Query("orderItem")
	orderDir := c.Query("orderDir")
	filter_by := c.Query("filter_by")
	Type := c.Query("Type")
	demography := c.Query("demography")
	status := c.Query("status")
	translation_status := c.Query("translation_status")
	webcomic := c.Query("webcomic")
	yonkoma := c.Query("yonkoma")
	amateur := c.Query("amateur")
	erotic := c.Query("erotic")

	if filter_by == "" {
		filter_by = "title"
	}
	if orderItem == "" {
		orderItem = "likes_count"
	}
	if orderDir == "" {
		orderDir = "desc"
	}
	mangas := tumangaonline.BuscarMangas(orderItem, orderDir, title, _page, filter_by, Type, demography, status, translation_status, webcomic, yonkoma, amateur, erotic)

	if mangas == nil {
		response := models.Response{
			StatusCode: http.StatusBadRequest,
			Data:       "Ocurrio un error",
		}
		c.JSON(response)
	}
	response := models.Response{
		StatusCode: http.StatusOK,
		Data:       mangas,
	}

	c.JSON(response)
}
