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

func GetJournalsByUserID(userID string, query requests.JournalQuery) (map[string]any, error) {
	var journals []models.Journal
	offset := (query.Page - 1) * query.Limit
	var total int64

	db := database.DB.Model(&models.Journal{}).Where("user_id = ?", userID)

	if query.Search != "" {
		db = db.Where("content LIKE ?", "%"+query.Search+"%")
	}

	db.Count(&total)
	if err := db.Offset(offset).Limit(query.Limit).Find(&journals).Error; err != nil {
		return nil, errors.New("failed to get journals")
	}

	res := mappers.ToListJournalResponse(journals)
	totalPages := (total + int64(query.Limit) - 1) / int64(query.Limit)
	return map[string]any{
		"data":        res,
		"total":       total,
		"page":        query.Page,
		"limit":       query.Limit,
		"total_pages": totalPages,
	}, nil
}

func GetJournalByDate(userID string, dateStr string) (responses.JournalResponse, error) {
	parsedDate, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return responses.JournalResponse{}, errors.New("invalid date format, must be YYYY-MM-DD")
	}

	var journal models.Journal
	if err := database.DB.Where("user_id = ? AND entry_date = ?", userID, parsedDate).First(&journal).Error; err != nil {
		return responses.JournalResponse{}, errors.New("journal entry not found for this date")
	}
	return mappers.ToJournalResponse(journal), nil
}

func CreateOrUpdateJournal(userID string, req requests.CreateJournalRequest) (responses.JournalResponse, error) {
	parsedDate, err := time.Parse("2006-01-02", req.EntryDate)
	if err != nil {
		return responses.JournalResponse{}, errors.New("invalid date format, must be YYYY-MM-DD")
	}

	if req.Mood < 1 || req.Mood > 5 {
		return responses.JournalResponse{}, errors.New("mood must be between 1 and 5")
	}

	tx := database.DB.Begin()
	if tx.Error != nil {
		return responses.JournalResponse{}, tx.Error
	}

	// Cek apakah user valid
	var user models.Users
	if err := tx.First(&user, "id = ?", userID).Error; err != nil {
		tx.Rollback()
		return responses.JournalResponse{}, errors.New("user not found")
	}

	var journal models.Journal
	err = tx.Where("user_id = ? AND entry_date = ?", userID, parsedDate).First(&journal).Error

	if err == nil {
		// Update existing
		journal.Mood = req.Mood
		journal.Content = req.Content
		if err := tx.Save(&journal).Error; err != nil {
			tx.Rollback()
			return responses.JournalResponse{}, errors.New("failed to update journal")
		}
	} else {
		// Create new
		journal = models.Journal{
			UserID:    userID,
			Mood:      req.Mood,
			Content:   req.Content,
			EntryDate: parsedDate,
		}
		if err := tx.Create(&journal).Error; err != nil {
			tx.Rollback()
			return responses.JournalResponse{}, errors.New("failed to create journal")
		}
	}

	if err := tx.Commit().Error; err != nil {
		return responses.JournalResponse{}, err
	}

	return mappers.ToJournalResponse(journal), nil
}

func DeleteJournal(userID string, id string) error {
	tx := database.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	var journal models.Journal
	if err := tx.First(&journal, "id = ? AND user_id = ?", id, userID).Error; err != nil {
		tx.Rollback()
		return errors.New("journal entry not found")
	}

	if err := tx.Delete(&journal).Error; err != nil {
		tx.Rollback()
		return errors.New("failed to delete journal entry")
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}
	return nil
}
