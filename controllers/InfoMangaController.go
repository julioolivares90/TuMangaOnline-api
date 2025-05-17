package controllers

import (
	"net/http"
	s "strings"

	"github.com/gofiber/fiber/v2"
	"github.com/julioolivares90/TumangaOnlineApi/core/tumangaonline"
	"github.com/julioolivares90/TumangaOnlineApi/models"
)

func GetInfoManga(c *fiber.Ctx) error {
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
	return c.JSON(response)

}

func GetCookiesFromTMO(c *fiber.Ctx) error {
	cookies, err := tumangaonline.GetCookiesFromTMO()

	if err != nil {
		response := models.Response{
			StatusCode: http.StatusBadGateway,
			Data:       err.Error(),
		}
		return c.JSON(response)
	}
	response := models.Response{
		StatusCode: http.StatusOK,
		Data:       cookies,
	}
	return c.JSON(response)
}

func GetPageFromTMOWithCookie(c *fiber.Ctx) error {

	url := c.Query("urlPage")
	urlRefer := c.Query("urlRefer")

	if s.Compare(url, "") == -1 {
		response := models.Response{
			StatusCode: http.StatusBadRequest,
			Data:       "el parametro urlPage es requerido",
		}
		return c.JSON(response)
	}
	if s.Compare(urlRefer, "") == -1 {
		response := models.Response{
			StatusCode: http.StatusBadRequest,
			Data:       "el parametro urlRefer es requerido",
		}
		c.JSON(response)
	}
	pages, err := tumangaonline.GetImageChapter2(urlRefer, url)
	if err != nil {
		response := models.Response{
			StatusCode: http.StatusBadRequest,
			Data:       "No se encontraron datos para mostrar",
		}
		return c.JSON(response)
	}
	if len(pages) <= 0 {
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
	return c.JSON(reponse)
}
