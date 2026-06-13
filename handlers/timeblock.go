package handlers

import (
	"be_dashboard/dto/requests"
	"be_dashboard/services"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetTimeblocksHandler(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)
	dayOfWeekStr := c.Query("day_of_week")
	dateStr := c.Query("date")

	var dayOfWeek *int
	if dayOfWeekStr != "" {
		val, err := strconv.Atoi(dayOfWeekStr)
		if err == nil {
			dayOfWeek = &val
		}
	}

	var date *time.Time
	if dateStr != "" {
		d, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid date format, use YYYY-MM-DD"})
		}
		date = &d
	}

	timeblocks, err := services.GetTimeblocksByUserID(userID, dayOfWeek, date)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Timeblocks retrieved successfully",
		"data":    timeblocks,
	})
}

func CreateTimeblockHandler(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)

	var req requests.CreateTimeblockRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	timeblock, err := services.CreateTimeblockService(userID, req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Timeblock created successfully",
		"data":    timeblock,
	})
}

func UpdateTimeblockHandler(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)
	id := c.Params("id")

	var req requests.CreateTimeblockRequest // use same struct for update
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	timeblock, err := services.UpdateTimeblockService(userID, id, req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Timeblock updated successfully",
		"data":    timeblock,
	})
}

func DeleteTimeblockHandler(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)
	id := c.Params("id")

	if err := services.DeleteTimeblockService(userID, id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Timeblock deleted successfully",
	})
}
