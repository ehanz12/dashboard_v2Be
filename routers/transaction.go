package routers

import (
	"be_dashboard/handlers"
	"be_dashboard/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupTransactionRoutes(api fiber.Router) {
	transaction := api.Group("/transactions")
	transaction.Post("/", middleware.ProtectedRoute, handlers.CreateTransactionHandler)
	transaction.Patch("/:id", middleware.ProtectedRoute, handlers.UpdateTransactionHandler)
	transaction.Delete("/:id", middleware.ProtectedRoute, handlers.DeleteTransactionHandler)
}