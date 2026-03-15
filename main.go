package main

import (
	"be_dashboard/config"
	"be_dashboard/database"
	"be_dashboard/routers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	//load environment variables
	config.LoadEnv()
	// connect to database
	database.Connect()
	//setup routes
	app := fiber.New()
	routers.SetupRoutes(app)
	//start server
	app.Listen(":" + config.AppConfig.Port)
}