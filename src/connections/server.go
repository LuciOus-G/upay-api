package connections

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"upay/api/v2/src/auth"
	"upay/api/v2/src/users"
)

func RunServer(conf ...interface{}) {
	app := fiber.New()

	//init Database
	DatabaseInit()
	_, err := DB().DB()
	if err != nil {
		panic(err)
	}

	// initiate route and swagger
	app.Get("/swagger/*", swagger.HandlerDefault)
	auth.RoutesAuth(app)
	users.RoutesUser(app)

	errApp := app.Listen(":8000")
	if errApp != nil {
		return
	}
}
