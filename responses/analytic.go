package responses

import (
	"image-converter/utils"

	"github.com/gofiber/fiber/v2"
)

type Analytic struct {
	Data map[int]utils.Analytic `json:"data"`
}

func AnalyticResponse(context *fiber.Ctx, statusCode int, data map[int]utils.Analytic) error {
	return Response(context, statusCode, Analytic{
		Data: data,
	})
}
