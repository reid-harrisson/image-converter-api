// Package handlers contains HTTP request handlers for the bank account manager
package handlers

// Import necessary packages for handling HTTP requests, responses, and services
import (
	"bytes"
	"image"
	responses "image-converter/responses"
	"image-converter/services"
	"image-converter/utils"
	"image/jpeg"
	"image/png"
	"log"
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

func loadImageFromFile(context *fiber.Ctx) (image.Image, error) {
	// Retrieve the file from the request
	fileHeader, err := context.FormFile("file")
	if err != nil {
		return &image.RGBA{}, responses.ErrorResponse(context, http.StatusBadRequest, "Invalid image file")
	}

	// Check the content type of the uploaded file
	contentType := fileHeader.Header.Get("Content-Type")

	log.Println(contentType)

	// Open the uploaded imageFile
	imageFile, err := fileHeader.Open()
	defer imageFile.Close()
	if err != nil {
		return &image.RGBA{}, responses.ErrorResponse(context, http.StatusBadRequest, "Fail to open file: "+err.Error())
	}

	// Decode the image based on its content type
	switch contentType {
	case "image/jpeg":
		img, err := jpeg.Decode(imageFile)
		if err != nil {
			return &image.RGBA{}, responses.ErrorResponse(context, http.StatusBadRequest, "Fail to decode jpeg file: "+err.Error())
		}
		return img, nil
	case "image/png":
		img, err := png.Decode(imageFile)
		if err != nil {
			return &image.RGBA{}, responses.ErrorResponse(context, http.StatusBadRequest, "Fail to decode jpeg file: "+err.Error())
		}
		return img, nil
	}

	return &image.RGBA{}, responses.ErrorResponse(context, http.StatusBadRequest, "Unsupported image type: "+contentType)
}

func sendImage(context *fiber.Ctx, convertedImage image.Image) error {
	// Use a buffer to encode the image
	var buf bytes.Buffer
	err := png.Encode(&buf, convertedImage)
	if err != nil {
		return responses.ErrorResponse(context, http.StatusInternalServerError, "Failed to encode image: "+err.Error())
	}
	// Set the appropriate content type and return the image
	context.Set("Content-Type", "image/png")
	return context.Status(http.StatusOK).Send(buf.Bytes())
}

// ConvertImage godoc
// @Summary Convert an image
// @Description Convert an image to a specified format
// @Tags Image
// @Accept multipart/form-data
// @Produce image/png
// @Param file formData file true "Image file to convert"
// @Param back_color query string true "Background Color" default(#000000)
// @Param fore_color query string true "Foreground Color" default(#FFFFFF)
// @Success 201 {object} responses.Message "Conversion successful"
// @Failure 400 {object} responses.Error "Invalid request"
// @Failure 500 {object} responses.Error "Internal server error"
// @Router /image [post]
func (handler *ImageHandler) Convert(context *fiber.Ctx) error {
	// Load source images
	sourceImage, err := loadImageFromFile(context)
	if err != nil {
		return err
	}

	log.Println(context.Queries())

	// Process validation request
	log.Println(context.Query("back_color"))
	log.Println(context.Query("fore_color"))

	backColor := utils.NewRGB(context.Query("back_color"))
	foreColor := utils.NewRGB(context.Query("fore_color"))

	log.Println(backColor)
	log.Println(foreColor)

	// Process the image to remove the background
	convertedImage := handler.ImageService.RemoveBackground(sourceImage, backColor, foreColor)

	// Send the image
	return sendImage(context, convertedImage)
}
