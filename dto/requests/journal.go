package requests

type CreateJournalRequest struct {
	Mood      int    `json:"mood" binding:"required"`
	Content   string `json:"content" binding:"required"`
	EntryDate string `json:"entry_date" binding:"required"` // Format: YYYY-MM-DD
}
