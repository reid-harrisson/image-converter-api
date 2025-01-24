package responses

import (
	"github.com/gofiber/fiber/v2"
)

type Error struct {
	Error string `json:"error"`
}

type Message struct {
	Message string `json:"message"`
}

func Response(context *fiber.Ctx, statusCode int, data interface{}) error {
	return context.Status(statusCode).JSON(data)
}

func ErrorResponse(ctx *fiber.Ctx, status int, message string) error {
	return Response(ctx, status, Error{
		Error: message,
	})
}

func MessageResponse(ctx *fiber.Ctx, status int, message string) error {
	return Response(ctx, status, Message{
		Message: message,
	})
}
