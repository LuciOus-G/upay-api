package auth

import "github.com/gofiber/fiber/v2"

func loginSession(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"code":  200,
		"error": "",
		"data": fiber.Map{
			"hello": "World",
		},
	})
}
