package routers

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App) {
	//grouping api versioning
	api := app.Group("/api/v1")

	//setup route auth
	SetupRouteAuth(api)
	//setup route category
	SetupCategoryRoutes(api)
	//setup route transaction
	SetupTransactionRoutes(api)
}