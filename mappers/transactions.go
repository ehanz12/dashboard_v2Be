package mappers

import (
	"be_dashboard/dto/responses"
	"be_dashboard/models"
)

func ToTransactionResponse(t models.Transactions) responses.TransactionResponse {
	var categoryRes *responses.CategoryMiniResponse
	if t.CategoryID != nil && t.Categories.ID != "" {
		categoryRes = &responses.CategoryMiniResponse{
			ID:   t.Categories.ID,
			Name: t.Categories.Name,
		}
	}

	return responses.TransactionResponse{
		ID:              t.ID,
		Category:        categoryRes,
		Amount:          t.Amount,
		Type:            t.Type,
		Description:     t.Description,
		TransactionDate: t.TransactionDate.Format("2006-01-02"),
	}
}

func ToTransactionResponses(transactions []models.Transactions) []responses.TransactionResponse {
	responses := make([]responses.TransactionResponse, 0, len(transactions))
	for _, t := range transactions {
		responses = append(responses, ToTransactionResponse(t))
	}
	return responses
}
