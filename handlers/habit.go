package handlers

import (
	"be_dashboard/dto/requests"
	"be_dashboard/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func CreateHabitHandler(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)

	var req requests.CreateHabitRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	habit, err := services.CreatHabitService(userID, req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"Message" : "Success to create habits", "data" : habit})
}

func GetHabitsByUserIDHandlers(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)

	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	query := requests.HabitQuery{
		Page:  page,
		Limit: limit,
	}

	habits, err := services.GetHabitsByUserIDService(userID, query)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(habits)

}