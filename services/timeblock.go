package services

import (
	"be_dashboard/database"
	"be_dashboard/dto/requests"
	"be_dashboard/dto/responses"
	"be_dashboard/mappers"
	"be_dashboard/models"
	"errors"
	"regexp"
)

func isValidTime(t string) bool {
	match, _ := regexp.MatchString(`^(0[0-9]|1[0-9]|2[0-3]):[0-5][0-9]$`, t)
	return match
}

func GetTimeblocksByUserID(userID string, dayOfWeek *int) ([]responses.TimeblockResponse, error) {
	var timeblocks []models.Timeblock
	query := database.DB.Where("user_id = ?", userID)

	if dayOfWeek != nil {
		query = query.Where("day_of_week = ?", *dayOfWeek)
	}

	if err := query.Order("day_of_week ASC, start_time ASC").Find(&timeblocks).Error; err != nil {
		return nil, err
	}
	return mappers.ToListTimeblockResponse(timeblocks), nil
}

func CreateTimeblockService(userID string, req requests.CreateTimeblockRequest) (responses.TimeblockResponse, error) {
	if req.DayOfWeek < 1 || req.DayOfWeek > 7 {
		return responses.TimeblockResponse{}, errors.New("day_of_week must be between 1 (Monday) and 7 (Sunday)")
	}

	if !isValidTime(req.StartTime) || !isValidTime(req.EndTime) {
		return responses.TimeblockResponse{}, errors.New("start_time and end_time must be in HH:MM format (24-hour)")
	}

	colorCode := req.ColorCode
	if colorCode == "" {
		colorCode = "#4F46E5"
	} else {
		match, _ := regexp.MatchString(`^#[0-9A-Fa-f]{6}$`, colorCode)
		if !match {
			return responses.TimeblockResponse{}, errors.New("color_code must be a valid 6-character HEX code (e.g. #4F46E5)")
		}
	}

	tx := database.DB.Begin()
	if tx.Error != nil {
		return responses.TimeblockResponse{}, tx.Error
	}

	var user models.Users
	if err := tx.First(&user, "id = ?", userID).Error; err != nil {
		tx.Rollback()
		return responses.TimeblockResponse{}, errors.New("user not found")
	}

	timeblock := models.Timeblock{
		UserID:       userID,
		ActivityName: req.ActivityName,
		StartTime:    req.StartTime,
		EndTime:      req.EndTime,
		ColorCode:    colorCode,
		DayOfWeek:    req.DayOfWeek,
	}

	if err := tx.Create(&timeblock).Error; err != nil {
		tx.Rollback()
		return responses.TimeblockResponse{}, errors.New("failed to create timeblock")
	}

	if err := tx.Commit().Error; err != nil {
		return responses.TimeblockResponse{}, err
	}

	return mappers.ToTimeblockResponse(timeblock), nil
}

func UpdateTimeblockService(userID string, id string, req requests.CreateTimeblockRequest) (responses.TimeblockResponse, error) {
	if req.DayOfWeek < 1 || req.DayOfWeek > 7 {
		return responses.TimeblockResponse{}, errors.New("day_of_week must be between 1 and 7")
	}

	if !isValidTime(req.StartTime) || !isValidTime(req.EndTime) {
		return responses.TimeblockResponse{}, errors.New("start_time and end_time must be in HH:MM format")
	}

	colorCode := req.ColorCode
	if colorCode == "" {
		colorCode = "#4F46E5"
	} else {
		match, _ := regexp.MatchString(`^#[0-9A-Fa-f]{6}$`, colorCode)
		if !match {
			return responses.TimeblockResponse{}, errors.New("color_code must be a valid HEX code")
		}
	}

	tx := database.DB.Begin()
	if tx.Error != nil {
		return responses.TimeblockResponse{}, tx.Error
	}

	var timeblock models.Timeblock
	if err := tx.First(&timeblock, "id = ? AND user_id = ?", id, userID).Error; err != nil {
		tx.Rollback()
		return responses.TimeblockResponse{}, errors.New("timeblock not found")
	}

	timeblock.ActivityName = req.ActivityName
	timeblock.StartTime = req.StartTime
	timeblock.EndTime = req.EndTime
	timeblock.ColorCode = colorCode
	timeblock.DayOfWeek = req.DayOfWeek

	if err := tx.Save(&timeblock).Error; err != nil {
		tx.Rollback()
		return responses.TimeblockResponse{}, errors.New("failed to update timeblock")
	}

	if err := tx.Commit().Error; err != nil {
		return responses.TimeblockResponse{}, err
	}

	return mappers.ToTimeblockResponse(timeblock), nil
}

func DeleteTimeblockService(userID string, id string) error {
	tx := database.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	var timeblock models.Timeblock
	if err := tx.First(&timeblock, "id = ? AND user_id = ?", id, userID).Error; err != nil {
		tx.Rollback()
		return errors.New("timeblock not found")
	}

	if err := tx.Delete(&timeblock).Error; err != nil {
		tx.Rollback()
		return errors.New("failed to delete timeblock")
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}
	return nil
}
