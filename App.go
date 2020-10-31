package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/julioolivares90/TumangaOnlineApi/controllers"
	"os"
)

type ResponsePath struct {
}

func main() {
	startServer()
}
func startServer() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {

		var paths []string

		paths = append(paths, "/api/v1/manga/populares")
		paths = append(paths, "/api/v1/manga/populares-josei")
		paths = append(paths, "/api/v1/manga/populares-seinen")
		paths = append(paths, "/api/v1/manga/info")
		paths = append(paths, "/api/v1/manga/library")
		paths = append(paths, "api/v1/manga/listas")
		paths = append(paths, "api/v1/get-cookies")
		paths = append(paths, "api/v1/get-manga")

		return c.JSON(paths)
	})

	app.Get("/api/v1/manga/populares", controllers.GetMangasPopularesWithPagination)
	app.Get("/api/v1/manga/populares-josei", controllers.GetMangasPopularesJosei)
	app.Get("/api/v1/manga/populares-seinen", controllers.GetMangasPopularesSeinen)
	app.Get("/api/v1/manga/info", controllers.GetInfoManga)
	//app.Get("/api/v1/manga/paginas", controllers.GetPaginasManga)
	//app.Get("/api/v1/manga/paginas/info", controllers.GetPaginasInfo)
	app.Get("/api/v1/manga/library", controllers.GetInfoLibrary)
	app.Get("api/v1/manga/listas", controllers.GetListasMangas)

	app.Get("api/v1/get-cookies", controllers.GetCookiesFromTMO)
	app.Get("api/v1/get-manga", controllers.GetPageFromTMOWithCookie)

	port := getPort()
	app.Listen(port)
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":5000"
	}
	return port
}
