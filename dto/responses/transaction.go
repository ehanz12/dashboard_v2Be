package responses

type TransactionResponse struct {
	ID              string               `json:"id"`
	Category        *CategoryMiniResponse `json:"category,omitempty"`
	Amount          float64              `json:"amount"`
	Type            string               `json:"type"`
	Description     string               `json:"description"`
	TransactionDate string               `json:"transaction_date"`
}
