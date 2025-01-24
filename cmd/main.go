package main

import (
	"image-converter/routes"
	"image-converter/server"
	"log"
)

// @Title Image Converter API
// @Version 1.0
// @BasePath /api/v1/
// @Description RESTful API endpoints for Image Conversion
func main() {
	server := server.Create()

	routes.ConfigRoutes(server)

	if err := server.Listen("3000"); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
