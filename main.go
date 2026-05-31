package main

import (
	"be_dashboard/config"
	"be_dashboard/database"
	"be_dashboard/routers"
	"github.com/gofiber/fiber/v2/middleware/cors"
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
	app.Use(cors.New(cors.Config{AllowOrigins: "http://localhost:5173,https://www.reihan.biz.id",
		AllowMethods:     "GET,POST,PUT,DELETE,PATCH,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true, //jika pake jwt
	}))
	//start server
	app.Listen(":" + config.AppConfig.Port)
}