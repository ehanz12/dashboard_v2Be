package handlers

import (
	"be_dashboard/database"
	 "github.com/gofiber/fiber/v2"
)

func RedisTest(c *fiber.Ctx) error {

	err := database.Redis.Set(database.Ctx, "nama", "Reihan", 0).Err()

	if err != nil {
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	value, _ := database.Redis.Get(database.Ctx, "nama").Result()

	return c.JSON(fiber.Map{
		"value": value,
	})
}
