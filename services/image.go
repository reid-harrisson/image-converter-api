package services

import (
	"image"
	"image-converter/utils"
	"image/color"
	"log"
	"math"
)

type ImageService struct {
}

// CreateAccountService initializes a new AccountService with the provided storage
func CreateImageService() *ImageService {
	return &ImageService{}
}

func colorSumC(color utils.RGB) float64 {
	return 0.2126*float64(color.R) + 0.7152*float64(color.G) + 0.0722*float64(color.B)
}

func colorSum3(r float64, g float64, b float64) float64 {
	return 0.2126*r + 0.7152*g + 0.0722*b
}

func (service *ImageService) RemoveBackground(img image.Image, back utils.RGB, fore utils.RGB) *image.RGBA {
	bounds := img.Bounds()
	convertedImage := image.NewRGBA(bounds)

	totalDiff := colorSumC(fore) - colorSumC(back)

	log.Println(bounds)
	log.Println(totalDiff)

	sum := 0.0

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			pixelColor := img.At(x, y)
			pureR, pureG, pureB, pureA := pixelColor.RGBA()

			// Convert 16-bit color values to 8-bit
			realR := float64(pureR) * 255 / 65535
			realG := float64(pureG) * 255 / 65535
			realB := float64(pureB) * 255 / 65535
			realA := float64(pureA) * 255 / 65535

			currentDiff := colorSum3(realR, realG, realB) - colorSumC(back)

			value := math.Floor(currentDiff * realA / totalDiff)

			sum += value

			if value > 255 {
				value = 255
			}
			if value < 0 {
				value = 0
			}

			alpha := uint8(value)

			newColor := color.NRGBA{
				R: fore.R, // Red component
				G: fore.G, // Green component
				B: fore.B, // Blue component
				A: alpha,  // Alpha channel based on brightness
			}

			convertedImage.Set(x, y, newColor)
		}
	}

	log.Println(sum)

	return convertedImage
}
