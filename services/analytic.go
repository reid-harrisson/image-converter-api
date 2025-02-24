package services

import (
	"image"
)

type AnalyticService struct {
}

// CreateAccountService initializes a new AccountService with the provided storage
func CreateAnalyticService() *AnalyticService {
	return &AnalyticService{}
}

func (service *AnalyticService) AnalyseImage(img image.Image) map[int]int {
	bounds := img.Bounds()

	data := map[int]int{}

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			pixelColor := img.At(x, y)
			pureR, pureG, pureB, pureA := pixelColor.RGBA()

			// Convert 16-bit color values to 8-bit
			realR := float64(pureR) * 255 / 65535
			realG := float64(pureG) * 255 / 65535
			realB := float64(pureB) * 255 / 65535
			_ = float64(pureA) * 255 / 65535

			currentBrightness := colorSum3(realR, realG, realB) * 1000 / 255
			data[int(currentBrightness)]++
		}
	}

	max := 0
	for _, value := range data {
		if max < value {
			max = value
		}
	}

	for key, value := range data {
		data[key] = value * 1000 / max
	}

	return data
}
