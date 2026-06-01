package routers

import (
	"be_dashboard/handlers"
	"be_dashboard/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupTaskRoutes(api fiber.Router) {
	task := api.Group("/tasks")

	task.Get("/", middleware.ProtectedRoute, handlers.GetTasksHandler)
	task.Post("/", middleware.ProtectedRoute, handlers.CreateTaskHandler)
	task.Patch("/:id", middleware.ProtectedRoute, handlers.UpdateTaskHandler)
	task.Patch("/:id/toggle", middleware.ProtectedRoute, handlers.ToggleTaskStatusHandler)
	task.Delete("/:id", middleware.ProtectedRoute, handlers.DeleteTaskHandler)
}
