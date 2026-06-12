package routers

import (
	"be_dashboard/handlers"
	"be_dashboard/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRouteAuth(api fiber.Router) {
	auth := api.Group("/auth")

	//register
	auth.Post("/register", handlers.CreateAuthHandler)

	//login
	auth.Post("/login", handlers.LoginAuthHandler)

	//me
	auth.Get("/me", middleware.ProtectedRoute,handlers.MeAuthHandler)

	//edit
	auth.Put("/edit", middleware.ProtectedRoute, handlers.EditMeAuthHandler)

	//change password
	auth.Patch("/change-password", middleware.ProtectedRoute, handlers.ChangePasswordAuthHandler)
}