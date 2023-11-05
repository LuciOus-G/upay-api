package users

import "github.com/gofiber/fiber/v2"

func RoutesUser(app *fiber.App) {
	localRoute := app.Group("/user")

	localRoute.Post("/register", CreateUser)
}
