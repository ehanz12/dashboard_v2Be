package services

import (
	"be_dashboard/database"
	"be_dashboard/dto/requests"
	"be_dashboard/dto/responses"
	"be_dashboard/mappers"
	"be_dashboard/models"
	"errors"
	"time"
)

func GetTasksByUserID(userID string, completed *bool) ([]responses.TaskResponse, error) {
	var tasks []models.Task
	query := database.DB.Where("user_id = ?", userID)

	if completed != nil {
		query = query.Where("is_completed = ?", *completed)
	}

	if err := query.Order("due_date ASC, created_at DESC").Find(&tasks).Error; err != nil {
		return nil, err
	}
	return mappers.ToListTaskResponse(tasks), nil
}

func CreateTaskService(userID string, req requests.CreateTaskRequest) (responses.TaskResponse, error) {
	tx := database.DB.Begin()
	if tx.Error != nil {
		return responses.TaskResponse{}, tx.Error
	}

	var user models.Users
	if err := tx.First(&user, "id = ?", userID).Error; err != nil {
		tx.Rollback()
		return responses.TaskResponse{}, errors.New("user not found")
	}

	var parsedDueDate *time.Time
	if req.DueDate != nil && *req.DueDate != "" {
		t, err := time.Parse("2006-01-02 15:04:05", *req.DueDate)
		if err != nil {
			// Coba alternatif format tanggal saja
			t2, err2 := time.Parse("2006-01-02", *req.DueDate)
			if err2 != nil {
				tx.Rollback()
				return responses.TaskResponse{}, errors.New("invalid due_date format, must be YYYY-MM-DD HH:MM:SS or YYYY-MM-DD")
			}
			parsedDueDate = &t2
		} else {
			parsedDueDate = &t
		}
	}

	quadrant := req.Quadrant
	if quadrant < 1 || quadrant > 4 {
		quadrant = 1 // Default to quadrant 1 (Urgent & Important)
	}

	task := models.Task{
		UserID:      userID,
		Title:       req.Title,
		Description: req.Description,
		Quadrant:    quadrant,
		IsCompleted: false,
		DueDate:     parsedDueDate,
	}

	if err := tx.Create(&task).Error; err != nil {
		tx.Rollback()
		return responses.TaskResponse{}, errors.New("failed to create task")
	}

	if err := tx.Commit().Error; err != nil {
		return responses.TaskResponse{}, err
	}

	return mappers.ToTaskResponse(task), nil
}

func UpdateTaskService(userID string, id string, req requests.UpdateTaskRequest) (responses.TaskResponse, error) {
	tx := database.DB.Begin()
	if tx.Error != nil {
		return responses.TaskResponse{}, tx.Error
	}

	var task models.Task
	if err := tx.First(&task, "id = ? AND user_id = ?", id, userID).Error; err != nil {
		tx.Rollback()
		return responses.TaskResponse{}, errors.New("task not found")
	}

	if req.Title != "" {
		task.Title = req.Title
	}
	task.Description = req.Description

	if req.Quadrant >= 1 && req.Quadrant <= 4 {
		task.Quadrant = req.Quadrant
	}

	if req.IsCompleted != nil {
		task.IsCompleted = *req.IsCompleted
	}

	if req.DueDate != nil {
		if *req.DueDate == "" {
			task.DueDate = nil
		} else {
			t, err := time.Parse("2006-01-02 15:04:05", *req.DueDate)
			if err != nil {
				t2, err2 := time.Parse("2006-01-02", *req.DueDate)
				if err2 != nil {
					tx.Rollback()
					return responses.TaskResponse{}, errors.New("invalid due_date format, must be YYYY-MM-DD HH:MM:SS or YYYY-MM-DD")
				}
				task.DueDate = &t2
			} else {
				task.DueDate = &t
			}
		}
	}

	if err := tx.Save(&task).Error; err != nil {
		tx.Rollback()
		return responses.TaskResponse{}, errors.New("failed to update task")
	}

	if err := tx.Commit().Error; err != nil {
		return responses.TaskResponse{}, err
	}

	return mappers.ToTaskResponse(task), nil
}

func ToggleTaskStatusService(userID string, id string) (responses.TaskResponse, error) {
	tx := database.DB.Begin()
	if tx.Error != nil {
		return responses.TaskResponse{}, tx.Error
	}

	var task models.Task
	if err := tx.First(&task, "id = ? AND user_id = ?", id, userID).Error; err != nil {
		tx.Rollback()
		return responses.TaskResponse{}, errors.New("task not found")
	}

	task.IsCompleted = !task.IsCompleted

	if err := tx.Save(&task).Error; err != nil {
		tx.Rollback()
		return responses.TaskResponse{}, errors.New("failed to toggle task status")
	}

	if err := tx.Commit().Error; err != nil {
		return responses.TaskResponse{}, err
	}

	return mappers.ToTaskResponse(task), nil
}

func DeleteTaskService(userID string, id string) error {
	tx := database.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	var task models.Task
	if err := tx.First(&task, "id = ? AND user_id = ?", id, userID).Error; err != nil {
		tx.Rollback()
		return errors.New("task not found")
	}

	if err := tx.Delete(&task).Error; err != nil {
		tx.Rollback()
		return errors.New("failed to delete task")
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}
	return nil
}
