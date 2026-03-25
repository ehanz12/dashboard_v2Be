package services

import (
	"be_dashboard/database"
	"be_dashboard/dto/requests"
	"be_dashboard/dto/responses"
	"be_dashboard/mappers"
	"be_dashboard/models"
	"errors"
	"time"

	"gorm.io/gorm"
)

func CreateTransaction(userID string, req requests.CreateTransactionRequest) (responses.TransactionResponse, error) {
	tx := database.DB.Begin()

	if tx.Error != nil {
		return responses.TransactionResponse{}, errors.New("failed to start transaction")
	}

	var users models.Users
	if err := tx.First(&users, "id = ?", userID).Error; err != nil {
		tx.Rollback()
		return responses.TransactionResponse{}, errors.New("user not found")
	}

	if req.CategoryID != nil {
		var category models.Categories
		if err := tx.First(&category, "id = ? AND user_id = ?", *req.CategoryID, userID).Error; err != nil {
			tx.Rollback()
			return responses.TransactionResponse{}, errors.New("category not found")
		}
	}

	if req.Type != "income" && req.Type != "expense" {
		tx.Rollback()
		return responses.TransactionResponse{}, errors.New("invalid transaction type")
	}

	date, err := time.Parse("2006-01-02", req.TransactionDate)
	if err != nil {
		tx.Rollback()
		return responses.TransactionResponse{}, errors.New("invalid date format, expected YYYY-MM-DD")
	}

	transaction := models.Transactions{
		UserID:          userID,
		CategoryID:      req.CategoryID,
		Amount:          req.Amount,
		Type:            req.Type,
		Description:     req.Description,
		TransactionDate: date,
	}

	if err := tx.Create(&transaction).Error; err != nil {
		tx.Rollback()
		return responses.TransactionResponse{}, errors.New("failed to create transaction")
	}

	tx.Preload("Categories", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "name")
	})

	if err := tx.Commit().Error; err != nil {
		return responses.TransactionResponse{}, errors.New("failed to commit transaction")
	}

	return responses.TransactionResponse{
		ID: transaction.ID,
		Category: &responses.CategoryMiniResponse{
			ID:   transaction.Categories.ID,
			Name: transaction.Categories.Name,
		},
		Amount:          transaction.Amount,
		Type:            transaction.Type,
		Description:     transaction.Description,
		TransactionDate: transaction.TransactionDate.Format("2006-01-02"),
	}, nil
}

func GetTransactions(userID string, query requests.TransactionQuery) (map[string]any, error) {
	var transactions []models.Transactions
	var total int64

	db := database.DB.Model(&models.Transactions{}).Where("user_id = ?", userID)

	if query.Type != "" {
		db = db.Where("type = ?", query.Type)
	}

	if query.CategoryID != "" {
		db = db.Where("category_id = ?", query.CategoryID)
	}

	if query.StartDate != "" && query.EndDate != "" {
		db = db.Where("transaction_date BETWEEN ? AND ?", query.StartDate, query.EndDate)
	}

	if query.Search != "" {
		db = db.Where("description LIKE ?", "%"+query.Search+"%")
	}

	db = db.Preload("Categories", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "name")
	})

	db.Count(&total)
	offset := (query.Page - 1) * query.Limit
	if err := db.Order("transaction_date DESC").Limit(query.Limit).Offset(offset).Find(&transactions).Error; err != nil {
		return nil, errors.New("failed to fetch transactions")
	}

	totalPages := (total + int64(query.Limit) - 1) / int64(query.Limit)

	transactionsResponse := mappers.ToTransactionResponses(transactions)

	return map[string]any{
		"data":       transactionsResponse,
		"total":      total,
		"page":       query.Page,
		"limit":      query.Limit,
		"totalPages": totalPages,
	}, nil
}

func UpdateTransaction(userID, id string, req requests.UpdateTransactionRequest) (responses.TransactionResponse, error) {
	tx := database.DB.Begin()

	if tx.Error != nil {
		return responses.TransactionResponse{}, errors.New("failed to start transaction")
	}

	var transaction models.Transactions
	if err := tx.First(&transaction, "id = ? AND user_id = ?", id, userID).Error; err != nil {
		tx.Rollback()
		return responses.TransactionResponse{}, errors.New("transaction not found")
	}

	if req.CategoryID != nil {
		var category models.Categories
		if err := tx.First(&category, "id = ? AND user_id = ?", *req.CategoryID, userID).Error; err != nil {
			tx.Rollback()
			return responses.TransactionResponse{}, errors.New("category not found")
		}
		transaction.CategoryID = req.CategoryID
	}

	if req.Amount != nil {
		transaction.Amount = *req.Amount
	}

	if req.Type != nil {
		if *req.Type != "income" && *req.Type != "expense" {
			tx.Rollback()
			return responses.TransactionResponse{}, errors.New("invalid transaction type")
		}
		transaction.Type = *req.Type
	}

	if req.Description != nil {
		transaction.Description = *req.Description
	}

	if req.TransactionDate != nil {
		date, err := time.Parse("2006-01-02", *req.TransactionDate)
		if err != nil {
			tx.Rollback()
			return responses.TransactionResponse{}, errors.New("invalid date format, expected YYYY-MM-DD")
		}
		transaction.TransactionDate = date
	}

	if req.Amount == nil &&
		req.Type == nil &&
		req.Description == nil &&
		req.CategoryID == nil &&
		req.TransactionDate == nil {
		return responses.TransactionResponse{}, errors.New("no data to update")
	}

	if err := tx.Save(&transaction).Error; err != nil {
		tx.Rollback()
		return responses.TransactionResponse{}, errors.New("failed to update transaction")
	}

	if err := tx.Preload("Categories", func(tx *gorm.DB) *gorm.DB {
		return tx.Select("id", "name")
	}).First(&transaction, "id = ?", transaction.ID).Error; err != nil {
		tx.Rollback()
		return responses.TransactionResponse{}, errors.New("failed to load transaction with category")
	}

	tx.Commit()
	return responses.TransactionResponse{
		ID: transaction.ID,
		Category: &responses.CategoryMiniResponse{
			ID:   transaction.Categories.ID,
			Name: transaction.Categories.Name,
		},
		Amount:          transaction.Amount,
		Type:            transaction.Type,
		Description:     transaction.Description,
		TransactionDate: transaction.TransactionDate.Format("2006-01-02"),
	}, nil

}

func DeleteTransaction(userID, id string) error {
	tx := database.DB.Begin()
	if tx.Error != nil {
		return errors.New("failed to start transaction")
	}
	var transaction models.Transactions
	if err := tx.First(&transaction, "id = ? AND user_id = ?", id, userID).Error; err != nil {
		tx.Rollback()
		return errors.New("transaction not found")
	}
	if err := tx.Delete(&transaction).Error; err != nil {
		tx.Rollback()
		return errors.New("failed to delete transaction")
	}
	if err := tx.Commit().Error; err != nil {
		return errors.New("failed to commit transaction")
	}
	return nil
}
