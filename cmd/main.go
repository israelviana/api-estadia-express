package main

import (
	"api-estadia-express/init/db"
	"api-estadia-express/init/enviroments"
	"api-estadia-express/init/logger"
	"api-estadia-express/server/routers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"os"
)

func main() {
	logger.InitZapLogger()
	enviroments.LoadEnv()
	db.InstanceDB()

	//cors
	crs := cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET, POST, DELETE, PUT, OPTIONS",
		AllowHeaders:     "*",
		AllowCredentials: true,
	})

	//Config Fiber Groups
	app := fiber.New()
	app.Use(crs)
	routers.Users(app)

	if err := app.Listen(os.Getenv("port_application")); err != nil {
		logger.Fatal("error to up server", err)
	}

}
