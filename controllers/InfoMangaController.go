package controllers

import (
	"github.com/gofiber/fiber"
	"github.com/julioolivares90/TumangaOnlineApi/core/tumangaonline"
	"github.com/julioolivares90/TumangaOnlineApi/models"
	"net/http"
	s "strings"
)

func GetInfoManga(c *fiber.Ctx) {
	urlAvisitar := c.Query("mangaUrl", "")
	if s.Compare(urlAvisitar, "") == -1 {
		response := models.Response{
			StatusCode: http.StatusBadRequest,
			Data:       "el parametro no puede estar vacio",
		}
		c.JSON(response)
	}

	mangaInfo := tumangaonline.GetInfoManga(urlAvisitar)
	response := models.Response{
		StatusCode: http.StatusOK,
		Data:       mangaInfo,
	}
	c.JSON(response)
}

func GetCookiesFromTMO(c *fiber.Ctx) {
	cookies, err := tumangaonline.GetCookiesFromTMO()

	if err != nil {
		response := models.Response{
			StatusCode: http.StatusBadGateway,
			Data:       err.Error(),
		}
		c.JSON(response)
	}
	response := models.Response{
		StatusCode: http.StatusOK,
		Data:       cookies,
	}
	c.JSON(response)
}

func GetPageFromTMOWithCookie(c *fiber.Ctx) {

	url := c.Query("urlPage")

	if url == "" {
		response := models.Response{
			StatusCode: http.StatusBadRequest,
			Data:       "el parametro urlPage es requerido",
		}
		c.JSON(response)
	}
	pages, err := tumangaonline.GetImageChapter(url)
	if err != nil {
		response := models.Response{
			StatusCode: http.StatusBadRequest,
			Data:       "No se encontraron datos para mostrar",
		}
		c.JSON(response)
	}
	if len(pages) < 0 {
		response := models.Response{
			StatusCode: http.StatusBadRequest,
			Data:       "No se encontraron datos para mostrar",
		}
		c.JSON(response)
	}
	reponse := models.Response{
		StatusCode: http.StatusOK,
		Data:       pages,
	}
	c.JSON(reponse)

}
