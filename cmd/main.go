package main

import (
	"api-estadia-express/init/db"
	"api-estadia-express/init/enviroments"
	"api-estadia-express/init/logger"
	"github.com/gofiber/fiber/v2"
	"os"
)

func main() {
	logger.InitZapLogger()
	enviroments.LoadEnv()
	db.ConnectionToPostgres()

	app := fiber.New()

	if err := app.Listen(os.Getenv("port_application")); err != nil {
		logger.Fatal("error to up server", err)
	}

}
