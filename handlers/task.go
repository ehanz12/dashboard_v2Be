package handlers

import (
	"be_dashboard/dto/requests"
	"be_dashboard/services"

	"github.com/gofiber/fiber/v2"
)

func GetTasksHandler(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)
	completedStr := c.Query("completed")

	var completed *bool
	if completedStr == "true" {
		t := true
		completed = &t
	} else if completedStr == "false" {
		f := false
		completed = &f
	}

	tasks, err := services.GetTasksByUserID(userID, completed)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Tasks retrieved successfully",
		"data":    tasks,
	})
}

func CreateTaskHandler(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)

	var req requests.CreateTaskRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	task, err := services.CreateTaskService(userID, req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Task created successfully",
		"data":    task,
	})
}

func UpdateTaskHandler(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)
	id := c.Params("id")

	var req requests.UpdateTaskRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	task, err := services.UpdateTaskService(userID, id, req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Task updated successfully",
		"data":    task,
	})
}

func ToggleTaskStatusHandler(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)
	id := c.Params("id")

	task, err := services.ToggleTaskStatusService(userID, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Task status toggled successfully",
		"data":    task,
	})
}

func DeleteTaskHandler(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)
	id := c.Params("id")

	if err := services.DeleteTaskService(userID, id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Task deleted successfully",
	})
}
