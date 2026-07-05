package main

import (
	"be_dashboard/config"
	"be_dashboard/cron"
	"be_dashboard/database"
	"be_dashboard/routers"
	"be_dashboard/services"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	//load environment variables
	config.LoadEnv()
	//load firebase service account
	services.InitFirebase()
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

scheduler := gocron.NewScheduler(time.Local)

scheduler.Every(1).Minute().Do(func() {
	cron.CheckHabitReminders()
})

scheduler.StartAsync()
// routers.PrintRoutes(app)
	//start server
	app.Listen(":" + config.AppConfig.Port)
}
