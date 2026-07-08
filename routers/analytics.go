package routers

import (
	"be_dashboard/handlers"
	"be_dashboard/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupAnalyticsRoutes(api fiber.Router) {
	analytics := api.Group("/analytics", middleware.ProtectedRoute)

	analytics.Get("/summary", handlers.GetAnalyticsSummaryHandler)
	analytics.Get("/finance", handlers.GetFinanceAnalyticsHandler)
	analytics.Get("/categories", handlers.GetCategoryAnalyticsHandler)
	analytics.Get("/habits", handlers.GetHabitAnalyticsHandler)
	analytics.Get("/tasks", handlers.GetTaskAnalyticsHandler)
	analytics.Get("/streak", handlers.GetStreakAnalyticsHandler)
}
