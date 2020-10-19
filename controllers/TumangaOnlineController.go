package controllers

import (
	"net/http"
	"strconv"

	s "strings"

	"github.com/gofiber/fiber"
	"github.com/julioolivares90/TumangaOnlineApi/core/tumangaonline"
	"github.com/julioolivares90/TumangaOnlineApi/models"
)

///GetMangasPopularesWithPagination
func GetMangasPopularesWithPagination(c *fiber.Ctx) {
	param := c.Params("pageNumber")
	id, err := strconv.Atoi(param)
	if err != nil {
		response := models.Response{
			StatusCode: http.StatusBadRequest,
			Data:       "error pageNumber no puede ser null",
		}
		c.JSON(response)
	}
	if id > 0 {
		mangasPopulares := tumangaonline.GetMangasPopulares(id)
		response := models.Response{
			StatusCode: http.StatusOK,
			Data:       mangasPopulares,
		}
		c.JSON(response)
	} else if id < 0 {
		mangasPopulares := tumangaonline.GetMangasPopulares(1)
		response := models.Response{
			StatusCode: http.StatusOK,
			Data:       mangasPopulares,
		}
		c.JSON(response)
	} else if id == 0 {
		mangasPopulares := tumangaonline.GetMangasPopulares(1)
		response := models.Response{
			StatusCode: http.StatusOK,
			Data:       mangasPopulares,
		}
		c.JSON(response)
	}
}
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

func GetMangasPopularesSeinen(c *fiber.Ctx) {
	mangas := tumangaonline.GetMangasPopularesSeinen()
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
