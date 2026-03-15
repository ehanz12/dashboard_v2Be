package routers

import (
	"be_dashboard/handlers"
	"be_dashboard/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupCategoryRoutes(api fiber.Router) {
	category := api.Group("/category")

	category.Post("/", middleware.ProtectedRoute,handlers.CreateCategoryHandler)
}