package handlers

import (
	"be_dashboard/dto/requests"
	"be_dashboard/services"

	"github.com/gofiber/fiber/v2"
)

func CreateTransactionHandler(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)

	var req requests.CreateTransactionRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}
	transaction, err := services.CreateTransaction(userID, req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message" : "successfuly to create transaction", "data" : transaction})
}

func UpdateTransactionHandler(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)
	transactionID := c.Params("id")

	var req requests.UpdateTransactionRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}
	transaction, err := services.UpdateTransaction(userID, transactionID, req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message" : "update successfuly", "data" : transaction})
}


func DeleteTransactionHandler(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)
	transactionID := c.Params("id")
	err := services.DeleteTransaction(userID, transactionID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message" : "delete successfuly"})
}