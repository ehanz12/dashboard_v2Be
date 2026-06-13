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

	if req.Date != nil {
		if *req.Date != "" {
			parsedDate, err := time.Parse("2006-01-02", *req.Date)
			if err != nil {
				tx.Rollback()
				return responses.HabitResponse{}, errors.New("invalid date format, expected YYYY-MM-DD")
			}
			habit.Date = &parsedDate
		}
	}

	if err := tx.Create(&habit).Error; err != nil {
		tx.Rollback()
		return responses.HabitResponse{}, errors.New("Failed to Create Habit !")
	}
	tx.Commit()
	var habitDate string
	if habit.Date != nil {
		habitDate = habit.Date.Format("2006-01-02")
	}
	return responses.HabitResponse{
		ID:        habit.ID,
		UserID:    habit.UserID,
		Name:      habit.Name,
		Frequency: habit.Frequency,
		Date:      habitDate,
		CreatedAt: habit.CreatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}

func GetHabitsByUserIDService(userID string, query requests.HabitQuery) (map[string]any, error) {
	var habits []models.Habits
	var total int64
	offset := (query.Page - 1) * query.Limit

	db := database.DB.Model(&models.Habits{}).Where("user_id = ?", userID)

	if query.Search != "" {
		db = db.Where("name LIKE ?", "%"+query.Search+"%")
	}

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

	originalName := habit.Name

	if req.Name != "" {
		habit.Name = req.Name
	}

	if req.Frequency != "" {
		habit.Frequency = req.Frequency
	}

	if req.Date != nil {
		if *req.Date == "" {
			habit.Date = nil
		} else {
			parsedDate, err := time.Parse("2006-01-02", *req.Date)
			if err != nil {
				tx.Rollback()
				return responses.HabitResponse{}, errors.New("invalid date format, expected YYYY-MM-DD")
			}
			habit.Date = &parsedDate
		}
	}

	if req.Name == "" && req.Frequency == "" && req.Date == nil {
		tx.Rollback()
		return responses.HabitResponse{}, errors.New("no data to update")
	}

	if req.Name != "" && req.Name != originalName {
		var exitsHabit models.Habits
		if err := tx.Where("user_id = ? AND name = ? AND id != ?", userID, req.Name, id).First(&exitsHabit).Error; err == nil {
			tx.Rollback()
			return responses.HabitResponse{}, errors.New("Habit Already Exists !")
		}
	}

	if err := tx.Save(&habit).Error; err != nil {
		tx.Rollback()
		return responses.HabitResponse{}, errors.New("failed to update")
	}

	tx.Commit()
	var habitDate string
	if habit.Date != nil {
		habitDate = habit.Date.Format("2006-01-02")
	}
	return responses.HabitResponse{
		ID:        habit.ID,
		UserID:    habit.UserID,
		Name:      habit.Name,
		Frequency: habit.Frequency,
		Date:      habitDate,
		CreatedAt: habit.CreatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}

func DeleteHabitService(userID, id string) error {
	var habit models.Habits
	tx := database.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	if err := tx.Where("user_id = ? AND id = ?", userID, id).First(&habit).Error; err != nil {
		tx.Rollback()
		return errors.New("Not Found Habit !")
	}

	if err := tx.Delete(&habit).Error; err != nil {
		tx.Rollback()
		return errors.New("Failed To Delete !")
	}

	tx.Commit()
	return nil
}

func TonggleHabitLogService(UserID, HabitID string) (responses.HabitResponseLog, error) {
	tx := database.DB.Begin()
	if tx.Error != nil {
		return responses.HabitResponseLog{}, tx.Error
	}

	var habit models.Habits
	if err := tx.Where("id = ? AND user_id = ?", HabitID, UserID).First(&habit).Error; err != nil {
		tx.Rollback()
		return responses.HabitResponseLog{}, errors.New("habit not found")
	}

	today := time.Now().Format("2006-01-02")

	var log models.HabitLogs
	err := tx.Where("habit_id = ? AND log_date = ?", HabitID, today).First(&log).Error
	if err != nil {
		newLog := models.HabitLogs{
			HabitID:   HabitID,
			LogDate:   time.Now(),
			Completed: true,
		}

		if err := tx.Create(&newLog).Error; err != nil {
			tx.Rollback()
			return responses.HabitResponseLog{}, errors.New("failed to create log")
		}
		tx.Commit()
		return responses.HabitResponseLog{
			ID: newLog.ID,
			Habit: responses.HabitSuperMini{
				ID:   habit.ID,
				Name: habit.Name,
			},
			Date:       today,
			IsComplete: true,
		}, nil
	}

	log.Completed = !log.Completed
	if err := tx.Save(&log).Error; err != nil {
		tx.Rollback()
		return responses.HabitResponseLog{}, errors.New("failed to update log")
	}

	tx.Commit()
	return responses.HabitResponseLog{
		ID: log.ID,
		Habit: responses.HabitSuperMini{
			ID:   habit.ID,
			Name: habit.Name,
		},
		Date:       today,
		IsComplete: log.Completed,
	}, nil
}

func GetHabitLogsTodayService(userID string) ([]responses.HabitLogTodayResponse, error) {

	var result []responses.HabitLogTodayResponse

	today := time.Now().Format("2006-01-02")

	err := database.DB.
		Table("habits").
		Select(`
			habits.id as habit_id,
			habits.name,
			COALESCE(habit_logs.completed, false) as completed
		`).
		Joins(`
			LEFT JOIN habit_logs 
			ON habit_logs.habit_id = habits.id 
			AND habit_logs.log_date = ?
		`, today).
		Where("habits.user_id = ?", userID).
		Order("habits.created_at DESC").
		Scan(&result).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}

func GetHabitLogsByDateService(userID, date string) ([]responses.HabitLogTodayResponse, error) {
	var result []responses.HabitLogTodayResponse

	err := database.DB.
		Table("habits").
		Select(`
			habits.id as habit_id,
			habits.name,
			COALESCE(habit_logs.completed, false) as completed
		`).
		Joins(`
			LEFT JOIN habit_logs 
			ON habit_logs.habit_id = habits.id 
			AND habit_logs.log_date = ?
		`, date).
		Where("habits.user_id = ?", userID).
		Order("habits.created_at DESC").
		Scan(&result).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}
