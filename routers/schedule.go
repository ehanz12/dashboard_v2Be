package routers

import (
	"be_dashboard/handlers"
	"be_dashboard/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupScheduleRoutes(api fiber.Router) {
	schedule := api.Group("/schedule")

	schedule.Get("/today", middleware.ProtectedRoute, handlers.GetScheduleHandler)
}
