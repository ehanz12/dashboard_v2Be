package services

import (
	"be_dashboard/database"
	"be_dashboard/dto/requests"
	"be_dashboard/dto/responses"
	"be_dashboard/models"
	"errors"
)

func CreateCategoryService(userID string, req requests.CategoryReq) (responses.CategoryResponse, error) {

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
		Name:   req.Name,
		Type:   req.Type,
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


func UpdateCategoryService(id string, req requests.CategoryReq) (responses.CategoryResponse, error) {
	tx := database.DB.Begin()

	if tx.Error != nil {
		return responses.CategoryResponse{}, tx.Error
	}

	var category models.Categories
	if err := tx.First(&category, "id = ?", id).Error; err != nil {
		tx.Rollback()
		return responses.CategoryResponse{}, errors.New("category not found")
	}

	category.Name = req.Name
	category.Type = req.Type
	if err := tx.Save(&category).Error; err != nil {
		tx.Rollback()
		return responses.CategoryResponse{}, errors.New("failed to update category")
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


func DeleteCategoryService(id string) error {
	tx := database.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	var category models.Categories
	if err := tx.First(&category, "id = ?", id).Error; err != nil {
		tx.Rollback()
		return errors.New("category not found")
	}
	if err := tx.Delete(&category).Error; err != nil {
		tx.Rollback()
		return errors.New("failed to delete category")
	}
	if err := tx.Commit().Error; err != nil {
		return err
	}
	return nil
}
