package routers

import (
	"be_dashboard/handlers"
	"be_dashboard/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupCategoryRoutes(api fiber.Router) {
	category := api.Group("/category")

	category.Get("/", middleware.ProtectedRoute, handlers.GetCategoriesHandler)
	category.Post("/", middleware.ProtectedRoute,handlers.CreateCategoryHandler)
	category.Patch("/:id", middleware.ProtectedRoute, handlers.UpdateCategoryHandler)
	category.Delete("/:id", middleware.ProtectedRoute, handlers.DeleteCategoryHandler)
}