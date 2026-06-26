package handlers

import (
	"be_dashboard/dto/requests"
	"be_dashboard/services"

	"github.com/gofiber/fiber/v2"
)

func SaveUserDevice(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)
	var req requests.UserDeviceRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}
	
	if err := services.SaveFCMToken(userID, req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save device information"})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Device information saved successfully"})
}

func DeleteUserDevice(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)
	var req requests.UserDeviceRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := services.DeleteFCMToken(userID, req.FMCToken); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Device removed successfully"})
}