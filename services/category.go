package services

import (
	"be_dashboard/database"
	"be_dashboard/dto/responses"
	"be_dashboard/models"
	"errors"
)

func CreateCategoryService(userID, name, categoryType string) (responses.CategoryResponse, error) {

	tx := database.DB.Begin()

	if tx.Error != nil {
		return responses.CategoryResponse{}, tx.Error
	}

	var user models.Users
	if err := tx.First(&user, "id = ?", userID).Error; err != nil {
		tx.Rollback()
		return responses.CategoryResponse{}, errors.New("user not found")
	}

	category := models.Categories{
		UserID: userID,
		Name:   name,
		Type:   categoryType,
	}

	if err := tx.Create(&category).Error; err != nil {
		tx.Rollback()
		return responses.CategoryResponse{}, errors.New("failed to create category")
	}

	if err := tx.Commit().Error; err != nil {
		return responses.CategoryResponse{}, err
	}

	return responses.CategoryResponse{
		ID:   category.ID,
		UserID: category.UserID,
		Name:   category.Name,
		Type:   category.Type,
	}, nil
}
