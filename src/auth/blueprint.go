package auth

import "github.com/gofiber/fiber/v2"

func RoutesAuth(app *fiber.App) {
	localRoute := app.Group("/auth")

	localRoute.Get("/login", loginSession)
}
