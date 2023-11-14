package routers

import (
	"api-estadia-express/server/controllers"
	"github.com/gofiber/fiber/v2"
)

func Users(app *fiber.App) {
	user := app.Group("/user")

	user.Post("/register", controllers.CreateUser)
	user.Post("/login", controllers.Login)
}
