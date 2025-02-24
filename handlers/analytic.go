// Package handlers contains HTTP request handlers for the bank account manager
package handlers

// Import necessary packages for handling HTTP requests, responses, and services
import (
	"image-converter/responses"
	"image-converter/services"

	"github.com/gofiber/fiber/v2"
)

type AnalyticHandler struct {
	AnalyticService *services.AnalyticService
}

func CreateAnalyticHandler() *AnalyticHandler {
	return &AnalyticHandler{
		AnalyticService: services.CreateAnalyticService(),
	}
}

// AnalyseImage godoc
// @Summary Analyse an image
// @Description Analyse an image
// @Tags Image
// @Accept multipart/form-data
// @Produce image/png
// @Param file formData file true "Image file to analyse"
// @Success 201 {object} responses.Analytic "Conversion successful"
// @Failure 400 {object} responses.Error "Invalid request"
// @Failure 500 {object} responses.Error "Internal server error"
// @Router /analytic [post]
func (handler *AnalyticHandler) Analyse(context *fiber.Ctx) error {
	// Load source images
	sourceImage, err := loadImageFromFile(context)
	if err != nil {
		return err
	}

	// Process the image to remove the background
	data := handler.AnalyticService.AnalyseImage(sourceImage)

	// Send the image
	return responses.AnalyticResponse(context, fiber.StatusOK, data)
}
