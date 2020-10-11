package main

import (
	"os"

	"github.com/gofiber/fiber"
	"github.com/julioolivares90/TumangaOnlineApi/controllers"
)

func main() {
	startServer()
}
func startServer() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) {

		c.Send("Hello TuMangaOnlineApi")
	})

	app.Get("/api/v1/manga/populares", controllers.GetMangasPopularesWithPagination)
	app.Get("/api/v1/manga/populares-josei", controllers.GetMangasPopularesJosei)
	app.Get("/api/v1/manga/populares-seinen", controllers.GetMangasPopularesSeinen)
	app.Get("/api/v1/manga/info", controllers.GetInfoManga)
	app.Get("/api/v1/manga/paginas", controllers.GetPaginasManga)
	//app.Get("/api/v1/manga/paginas/info", controllers.GetPaginasInfo)
	app.Get("/api/v1/manga/library", controllers.GetInfoLibrary)
	app.Get("api/v1/manga/listas", controllers.GetListasMangas)

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
