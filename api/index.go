package api

import (
	"image-converter/routes"
	s "image-converter/server"
	"net/http"

	"github.com/gofiber/adaptor/v2"
)

var (
	server *s.Server
)

func init() {
	server = s.Create()

	routes.ConfigRoutes(server)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	adaptor.FiberApp(server.App).ServeHTTP(w, r)
}
