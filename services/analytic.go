package services

import (
	"image"
	"image-converter/utils"
)

type AnalyticService struct {
}

// CreateAccountService initializes a new AccountService with the provided storage
func CreateAnalyticService() *AnalyticService {
	return &AnalyticService{}
}

func (service *AnalyticService) AnalyseImage(img image.Image) map[int]utils.Analytic {
	bounds := img.Bounds()

	data := map[int]utils.Analytic{}

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			pixelColor := img.At(x, y)
			pureR, pureG, pureB, pureA := pixelColor.RGBA()

			// Convert 16-bit color values to 8-bit
			realR := float64(pureR) * 255 / 65535
			realG := float64(pureG) * 255 / 65535
			realB := float64(pureB) * 255 / 65535
			_ = float64(pureA) * 255 / 65535

			brightness := int(colorSum3(realR, realG, realB) * 200 / 255)
			datum, isExisted := data[brightness]
			if !isExisted {
				data[brightness] = utils.Analytic{
					Count: 0,
					Color: "",
				}
			}
			data[brightness] = utils.Analytic{
				Count: datum.Count + 1,
				Color: utils.ColorToHex(pixelColor),
			}
		}
	}

	max := 0
	for _, datum := range data {
		if max < datum.Count {
			max = datum.Count + 1
		}
	}

	for i := 0; i < 200; i++ {
		_, isExited := data[i]
		if !isExited {
			data[i] = utils.Analytic{
				Count: 0,
				Color: "",
			}
		}
	}

	for key, datum := range data {
		data[key] = utils.Analytic{
			Count: datum.Count * 200 / max,
			Color: datum.Color,
		}
	}

	return data
}
