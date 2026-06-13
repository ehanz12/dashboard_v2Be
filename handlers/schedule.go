package handlers

import (
	"be_dashboard/services"

	"github.com/gofiber/fiber/v2"
)

func GetScheduleHandler(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)
	date := c.Query("date", "")

	schedule, err := services.GetScheduleForDateService(userID, date)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Schedule retrieved successfully",
		"data":    schedule,
	})
}
