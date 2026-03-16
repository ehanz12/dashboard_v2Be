package handlers

import (
	"be_dashboard/dto/requests"
	"be_dashboard/services"

	"github.com/gofiber/fiber/v2"
)

func GetCategoriesHandler(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)
	categories, err := services.GetCategoriesByUserID(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error" : err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data" : categories})
}

//create category
func CreateCategoryHandler(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)

	//validasi payload
	var req requests.CategoryReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error" : "invalid payload"})
	}

	//call service
	category, err := services.CreateCategoryService(userID, req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error" : err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message" : "category created successfully", "data" : category})
}

func UpdateCategoryHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	//validasi payload
	var req requests.CategoryReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error" : "invalid payload"})
	}
	category, err := services.UpdateCategoryService(id, req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error" : err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message" : "category updated successfully", "data" : category})
}

func DeleteCategoryHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	err := services.DeleteCategoryService(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error" : err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message" : "category deleted successfully"})
}