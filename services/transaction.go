package services

import (
	"be_dashboard/database"
	"be_dashboard/dto/requests"
	"be_dashboard/dto/responses"
	"be_dashboard/models"
	"errors"
	"time"
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

	if err := tx.Commit().Error; err != nil {
		return responses.TransactionResponse{}, errors.New("failed to commit transaction")
	}

	return responses.TransactionResponse{
		ID:              transaction.ID,
		CategoryID:      transaction.CategoryID,
		Amount:          transaction.Amount,
		Type:            transaction.Type,
		Description:     transaction.Description,
		TransactionDate: transaction.TransactionDate.Format("2006-01-02"),
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

	if err := tx.Save(&transaction).Error; err != nil {
		tx.Rollback()
		return responses.TransactionResponse{}, errors.New("failed to update transaction")
	}
	tx.Commit()
	return responses.TransactionResponse{
		ID:              transaction.ID,
		CategoryID:      transaction.CategoryID,
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

