package routes

import (
	_ "image-converter/docs"
	"image-converter/server"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func ConfigRoutes(server *server.Server) {
	server.App.Get("/swagger/*", swagger.HandlerDefault)
	server.App.Get("/", redirectToSwagger)
}

func redirectToSwagger(context *fiber.Ctx) error {
	return context.Redirect("/swagger/index.html")
}
