package server

import (
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	App *fiber.App
}

func Create() *Server {
	app := fiber.New()

	return &Server{
		App: app,
	}
}

func (server *Server) Listen(port string) error {
	return server.App.Listen(":" + port)
}
