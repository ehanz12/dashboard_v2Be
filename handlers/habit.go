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
		Search: c.Query("search"),
	}

	habits, err := services.GetHabitsByUserIDService(userID, query)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(habits)
}

func GetHabitLogTodayByUserIDHandler(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)
	logs, err := services.GetHabitLogsTodayService(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data" : logs, "Message" : "Success to Get Habit Log Today"})
}

func GetHabitSummaryHandler(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)
	HabitID := c.Params("id")

	summary, err := services.GetHabitSummaryService(userID, HabitID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data" : summary, "Message" : "Success to Get Habit Summary"})
}


func UpdateHabitHandler(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)

	HabitID := c.Params("id")

	var req requests.CreateHabitRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error" : "invalid payload"})
	}

	update, err := services.UpdateHabitService(userID, HabitID, req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error" : err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"Message" : "Update SuccessFully", "data" : update})
}

func DeleteHabitHandler(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)

	habitID := c.Params("id")

	if err := services.DeleteHabitService(userID, habitID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error" : err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"Message" : "Success to Delete Habit"})
}


func ToggleHabitLogHandler(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)

	var req requests.ToogleHabitLogRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error" : "invalid payload"})
	}

	log, err := services.TonggleHabitLogService(userID, req.HabitID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error" : err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"Message" : "Success to Toggle Habit Log", "data" : log})
}