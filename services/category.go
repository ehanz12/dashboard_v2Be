package services

import (
	"be_dashboard/database"
	"be_dashboard/dto/requests"
	"be_dashboard/dto/responses"
	"be_dashboard/mappers"
	"be_dashboard/models"
	"errors"
)

func GetCategoriesByUserID(userID string) ([]responses.CategoryResponse, error) {
	var categories []models.Categories
	if err := database.DB.Where("user_id = ?", userID).Find(&categories).Error; err != nil {
		return nil, err
	}
	return mappers.ListToCategoryRes(categories), nil
}

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
	}, nil
}


func UpdateCategoryService(userID string, id string, req requests.CategoryReq) (responses.CategoryResponse, error) {
	tx := database.DB.Begin()

	if tx.Error != nil {
		return responses.CategoryResponse{}, tx.Error
	}

	var category models.Categories
	if err := tx.First(&category, "id = ? AND user_id = ?", id, userID).Error; err != nil {
		tx.Rollback()
		return responses.CategoryResponse{}, errors.New("category not found")
	}

	category.Name = req.Name
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
	}, nil
}


func DeleteCategoryService(userID string, id string) error {
	tx := database.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	var category models.Categories
	if err := tx.First(&category, "id = ? AND user_id = ?", id, userID).Error; err != nil {
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
