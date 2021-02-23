package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/julioolivares90/TumangaOnlineApi/core/tumangaonline"
	"github.com/julioolivares90/TumangaOnlineApi/models"
)

///GetMangasPopularesWithPagination
func GetMangasPopularesWithPagination(c *fiber.Ctx) error {

	param := c.Query("pageNumber", "0")
	fmt.Println(param)
	id, err := strconv.Atoi(param)
	if err != nil {
		response := models.Response{
			StatusCode: http.StatusBadRequest,
			Data:       "error pageNumber no puede ser null",
		}
		return c.JSON(response)
	}
	if id > 0 {
		mangasPopulares := tumangaonline.GetMangasPopulares(id)
		response := models.Response{
			StatusCode: http.StatusOK,
			Data:       mangasPopulares,
		}
		return c.JSON(response)
	} else if id < 0 {
		mangasPopulares := tumangaonline.GetMangasPopulares(1)
		response := models.Response{
			StatusCode: http.StatusOK,
			Data:       mangasPopulares,
		}
		return c.JSON(response)
	} else if id == 0 {
		mangasPopulares := tumangaonline.GetMangasPopulares(1)
		response := models.Response{
			StatusCode: http.StatusOK,
			Data:       mangasPopulares,
		}
		return c.JSON(response)
	}
	return c.JSON("")
}
