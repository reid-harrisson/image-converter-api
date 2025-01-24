package services

import (
	"image"
	"image/color"
)

type ImageService struct {
}

// CreateAccountService initializes a new AccountService with the provided storage
func CreateImageService() *ImageService {
	return &ImageService{}
}

func (service *ImageService) RemoveBackground(img image.Image) *image.RGBA {
	bounds := img.Bounds()
	newImg := image.NewRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := img.At(x, y)
			r, g, b, _ := c.RGBA()

			// Convert 16-bit color values to 8-bit
			rr := float64(r) * 255 / 65535
			gg := float64(g) * 255 / 65535
			bb := float64(b) * 255 / 65535

			// Calculate perceived brightness using standard coefficients
			brightness := 0.2126*rr + 0.7152*gg + 0.0722*bb

			alpha := uint8(brightness)
			if brightness < 0 {
				alpha = 0
			}

			newColor := color.NRGBA{
				R: 255,   // Red component
				G: 255,   // Green component
				B: 255,   // Blue component
				A: alpha, // Alpha channel based on brightness
			}

			newImg.Set(x, y, newColor)
		}
	}

	return newImg
}
