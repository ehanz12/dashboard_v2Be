package routers

import (
	// "fmt"
	"time"
	"be_dashboard/handlers"
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
	//testing redis
	api.Get("/redis", handlers.RedisTest)
	//setup route auth
	SetupRouteAuth(api)
	//setup route category
	SetupCategoryRoutes(api)
	//setup route transaction
	SetupTransactionRoutes(api)
	//setup route habit
	SetupHabitRoutes(api)
	//setup route task
	SetupTaskRoutes(api)
	//setup route timeblock
	SetupTimeblockRoutes(api)
	//setup route schedule
	SetupScheduleRoutes(api)
	//setup route device
	SetupDeviceRoutes(api)
	//setup route dashboard
	SetupDashboardRoutes(api)
	//setup route analytics
	SetupAnalyticsRoutes(api)
}

// func PrintRoutes(app *fiber.App) {
// 	fmt.Println("================================================================")
// 	fmt.Printf("%-8s %-35s %s\n", "METHOD", "PATH", "NAME")
// 	fmt.Println("================================================================")

// 	count := 0

// 	for _, stack := range app.Stack() {
// 		for _, route := range stack {
// 			count++
// 			fmt.Printf("%-8s %-35s %s\n",
// 				route.Method,
// 				route.Path,
// 				route.Name,
// 			)
// 		}
// 	}

// 	fmt.Println("================================================================")
// 	fmt.Printf("Total Routes: %d\n", count)
// }
