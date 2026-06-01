package routers

import (
	"be_dashboard/handlers"
	"be_dashboard/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupJournalRoutes(api fiber.Router) {
	journal := api.Group("/journals")

	journal.Get("/", middleware.ProtectedRoute, handlers.GetJournalsHandler)
	journal.Get("/date/:date", middleware.ProtectedRoute, handlers.GetJournalByDateHandler)
	journal.Post("/", middleware.ProtectedRoute, handlers.CreateOrUpdateJournalHandler)
	journal.Delete("/:id", middleware.ProtectedRoute, handlers.DeleteJournalHandler)
}
