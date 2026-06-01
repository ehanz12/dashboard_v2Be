package routers

import (
	"be_dashboard/handlers"
	"be_dashboard/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupTimeblockRoutes(api fiber.Router) {
	timeblock := api.Group("/timeblocks")

	timeblock.Get("/", middleware.ProtectedRoute, handlers.GetTimeblocksHandler)
	timeblock.Post("/", middleware.ProtectedRoute, handlers.CreateTimeblockHandler)
	timeblock.Patch("/:id", middleware.ProtectedRoute, handlers.UpdateTimeblockHandler)
	timeblock.Delete("/:id", middleware.ProtectedRoute, handlers.DeleteTimeblockHandler)
}
