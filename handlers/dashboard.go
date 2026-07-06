package handlers

import (
	"be_dashboard/services"

	"github.com/gofiber/fiber/v2"
)

func GetDashboardHandler(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)

	data, err := services.GetDashboardData(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Success",
		"data":    data,
	})
}
