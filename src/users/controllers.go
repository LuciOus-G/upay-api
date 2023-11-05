package users

import (
	"github.com/gofiber/fiber/v2"
	db "upay/api/v2/src/connections"
	m "upay/api/v2/src/models"
)

func CreateUser(ctx *fiber.Ctx) error {
	//	check if user with the number exists
	userExists := m.UserModel{PhoneNumber: "08"}
	db.DB().First(&userExists)

	return ctx.JSON(fiber.Map{
		"test": "test",
	})
}
