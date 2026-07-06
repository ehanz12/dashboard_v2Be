package routers

import (
	"be_dashboard/handlers"
	"be_dashboard/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupDashboardRoutes(api fiber.Router) {
	api.Get("/dashboard", middleware.ProtectedRoute, handlers.GetDashboardHandler)
}
