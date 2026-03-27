package services

import (
	"be_dashboard/database"
	"be_dashboard/dto/requests"
	"be_dashboard/dto/responses"
	"be_dashboard/mappers"
	"be_dashboard/models"
	"errors"
)

func CreatHabitService(userID string, req requests.CreateHabitRequest) (responses.HabitResponse, error) {
	tx := database.DB.Begin()
	if tx.Error != nil {
		return responses.HabitResponse{}, tx.Error
	}
	var exits models.Users
	if err := tx.Where("id = ?", userID).First(&exits).Error; err != nil {
		tx.Rollback()
		return responses.HabitResponse{}, errors.New("User Not Found !")
	}

	if req.Frequency != "daily" && req.Frequency != "weekly" {
		tx.Rollback()
		return responses.HabitResponse{}, errors.New("Invalid Frequency !")
	}

	var exitsHabit models.Habits
	if err := tx.Where("user_id = ? AND name = ?", userID, req.Name).First(&exitsHabit).Error; err == nil {
		tx.Rollback()
		return responses.HabitResponse{}, errors.New("Habit Already Exists !")
	}

	habit := models.Habits{
		UserID:    userID,
		Name:      req.Name,
		Frequency: req.Frequency,
	}

	if err := tx.Create(&habit).Error; err != nil {
		tx.Rollback()
		return responses.HabitResponse{}, errors.New("Failed to Create Habit !")
	}
	tx.Commit()
	return responses.HabitResponse{
		ID:        habit.ID,
		UserID:    habit.UserID,
		Name:      habit.Name,
		Frequency: habit.Frequency,
		CreatedAt: habit.CreatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}

func GetHabitsByUserIDService(userID string, query requests.HabitQuery) (map[string]any, error) {
	var habits []models.Habits
	var total int64
	offset := (query.Page - 1) * query.Limit

	db := database.DB.Model(&models.Habits{}).Where("user_id = ?", userID)

	db.Count(&total)
	if err := db.Offset(offset).Limit(query.Limit).Find(&habits).Error; err != nil {
		return nil, errors.New("Failed to Get Habits !")
	}

	res := mappers.ToListResponseHabitMini(habits)
	totalPages := (total + int64(query.Limit) - 1) / int64(query.Limit)
	return map[string]any{
		"data":        res,
		"total":       total,
		"page":        query.Page,
		"limit":       query.Limit,
		"total_pages": totalPages,
	}, nil
}

func UpdateHabitService(userID, id string, req requests.CreateHabitRequest) (responses.HabitResponse, error) {
	tx := database.DB.Begin()
	if tx.Error != nil {
		return responses.HabitResponse{}, tx.Error
	}
	var user models.Users
	if err := tx.First(&user, "id = ?", userID).Error; err != nil {
		tx.Rollback()
		return responses.HabitResponse{}, errors.New("Not Found User !")
	}

	var habit models.Habits
	if err := tx.First(&habit, "id = ? AND user_id = ?", id, userID).Error; err != nil {
		tx.Rollback()
		return responses.HabitResponse{}, errors.New("Not Found Habit !")
	}

	if req.Name != "" {
		habit.Name = req.Name
	}

	if req.Frequency != "" {
		habit.Frequency = req.Frequency
	}

	if req.Name == "" && req.Frequency == "" {
		return responses.HabitResponse{}, errors.New("no data to update")
	}

	var exitsHabit models.Habits
	if err := tx.Where("user_id = ? AND name = ?", userID, req.Name).First(&exitsHabit).Error; err == nil {
		tx.Rollback()
		return responses.HabitResponse{}, errors.New("Habit Already Exists !")
	}

	if err := tx.Save(&habit).Error; err != nil {
		tx.Rollback()
		return responses.HabitResponse{}, errors.New("failed to update")
	}

	tx.Commit()
	return responses.HabitResponse{
		ID:        habit.ID,
		UserID:    habit.UserID,
		Name:      habit.Name,
		Frequency: habit.Frequency,
		CreatedAt: habit.CreatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}
