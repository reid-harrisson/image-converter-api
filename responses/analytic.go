package responses

import "github.com/gofiber/fiber/v2"

type Analytic struct {
	Brightness map[int]int `json:"brightness"`
}

func AnalyticResponse(context *fiber.Ctx, statusCode int, data map[int]int) error {
	return Response(context, statusCode, Analytic{
		Brightness: data,
	})
}
