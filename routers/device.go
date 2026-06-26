package routers

import (
	"be_dashboard/handlers"
	"be_dashboard/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupDeviceRoutes(api fiber.Router) {
	device := api.Group("/devices")

	device.Post("/", middleware.ProtectedRoute, handlers.SaveUserDevice)
	device.Delete("/", middleware.ProtectedRoute, handlers.DeleteUserDevice)
}
