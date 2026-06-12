package routers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func SetupRoutes(app *fiber.App) {
	// Konfigurasi Rate Limiting standar profesional
	// Membatasi setiap IP maksimal 100 request per 1 menit
	apiLimiter := limiter.New(limiter.Config{
		Max:        100,             // Maksimal request
		Expiration: 1 * time.Minute, // Durasi limit
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP() // Berdasarkan alamat IP klien
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"status":  "error",
				"message": "Too many requests. Please try again later.",
			})
		},
	})

	// grouping api versioning & apply limiter middleware
	api := app.Group("/api/v1", apiLimiter)

	//setup route auth
	SetupRouteAuth(api)
	//setup route category
	SetupCategoryRoutes(api)
	//setup route transaction
	SetupTransactionRoutes(api)
	//setup route habit
	SetupHabitRoutes(api)
	//setup route journal
	SetupJournalRoutes(api)
	//setup route task
	SetupTaskRoutes(api)
	//setup route timeblock
	SetupTimeblockRoutes(api)
}