// Package handlers contains HTTP request handlers for the bank account manager
package handlers

// Import necessary packages for handling HTTP requests, responses, and services
import (
	"bytes"
	"image"
	responses "image-converter/responses"
	"image-converter/services"
	"image/jpeg"
	"image/png"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type ImageHandler struct {
	ImageService *services.ImageService
}

func CreateImageHandler() *ImageHandler {
	return &ImageHandler{
		ImageService: services.CreateImageService(),
	}
}

// ConvertImage godoc
// @Summary Convert an image
// @Description Convert an image to a specified format
// @Tags Image
// @Accept multipart/form-data, application/json
// @Produce multipart/form-data
// @Param file formData file true "Image file to convert"
// @Param params body requests.ImageRequest true "Conversion configuration details"
// @Success 201 {object} responses.Message "Conversion successful"
// @Failure 400 {object} responses.Error "Invalid request"
// @Failure 500 {object} responses.Error "Internal server error"
// @Router /image [post]
func (handler *ImageHandler) Convert(context *fiber.Ctx) error {
	// Retrieve the file from the request
	fileHeader, err := context.FormFile("file")
	if err != nil {
		return responses.ErrorResponse(context, http.StatusBadRequest, "Invalid image file")
	}

	// Check the content type of the uploaded file
	contentType := fileHeader.Header.Get("Content-Type")

	// Open the uploaded imageFile
	imageFile, err := fileHeader.Open()
	if err != nil {
		return responses.ErrorResponse(context, http.StatusInternalServerError, "Failed to open image file: "+err.Error())
	}
	defer imageFile.Close()

	// Decode the image based on its content type
	var sourceImage image.Image
	switch contentType {
	case "image/jpeg":
		sourceImage, err = jpeg.Decode(imageFile)
		if err != nil {
			return responses.ErrorResponse(context, http.StatusBadRequest, "Failed to decode JPEG image: "+err.Error())
		}
	case "image/png":
		sourceImage, err = png.Decode(imageFile)
		if err != nil {
			return responses.ErrorResponse(context, http.StatusBadRequest, "Failed to decode PNG image: "+err.Error())
		}
	default:
		return responses.ErrorResponse(context, http.StatusUnsupportedMediaType, "Unsupported image type: "+contentType)
	}

	// Process the image to remove the background
	convertedImage := handler.ImageService.RemoveBackground(sourceImage)

	// Use a buffer to encode the image
	var buf bytes.Buffer
	err = png.Encode(&buf, convertedImage)
	if err != nil {
		return responses.ErrorResponse(context, http.StatusInternalServerError, "Failed to encode image: "+err.Error())
	}

	// Set the appropriate content type and return the image
	context.Set("Content-Type", "image/png")
	return context.Send(buf.Bytes())
}
