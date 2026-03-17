package requests

type CreateTransactionRequest struct {
	CategoryID      *string  `json:"category_id"`
	Amount          float64  `json:"amount"`
	Type            string   `json:"type"`
	Description     string   `json:"description"`
	TransactionDate string   `json:"transaction_date"`
}

type UpdateTransactionRequest struct {
	CategoryID      *string  `json:"category_id"`
	Amount          *float64 `json:"amount"`
	Type            *string  `json:"type"`
	Description     *string  `json:"description"`
	TransactionDate *string  `json:"transaction_date"`
}