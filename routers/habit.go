package routers

import (
	"be_dashboard/handlers"
	"be_dashboard/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupHabitRoutes(api fiber.Router) {
	habit := api.Group("/habits")

	habit.Get("/", middleware.ProtectedRoute, handlers.GetHabitsByUserIDHandlers)
	habit.Post("/", middleware.ProtectedRoute, handlers.CreateHabitHandler)
	habit.Patch("/:id", middleware.ProtectedRoute, handlers.UpdateHabitHandler)
	habit.Delete("/:id", middleware.ProtectedRoute, handlers.DeleteHabitHandler)


	habit.Get("/logs/today", middleware.ProtectedRoute, handlers.GetHabitLogTodayByUserIDHandler)
	habit.Get("/:id/streak", middleware.ProtectedRoute, handlers.GetHabitSummaryHandler)
	habit.Post("/toggle", middleware.ProtectedRoute, handlers.ToggleHabitLogHandler)

}
