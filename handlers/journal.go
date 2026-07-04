package handlers

import (
	"be_dashboard/dto/requests"
	"be_dashboard/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetJournalsHandler(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	query := requests.JournalQuery{
		Page:  page,
		Limit: limit,
		Search: c.Query("search"),
	}

	journals, err := services.GetJournalsByUserID(userID, query)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(journals)
}

func GetJournalByDateHandler(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)
	dateStr := c.Params("date")

	journal, err := services.GetJournalByDate(userID, dateStr)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Journal retrieved successfully",
		"data":    journal,
	})
}

func CreateOrUpdateJournalHandler(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)

	var req requests.CreateJournalRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	journal, err := services.CreateOrUpdateJournal(userID, req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Journal saved successfully",
		"data":    journal,
	})
}

func DeleteJournalHandler(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)
	id := c.Params("id")

	if err := services.DeleteJournal(userID, id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Journal entry deleted successfully",
	})
}
