package requests

type CreateJournalRequest struct {
	Mood      int    `json:"mood" binding:"required"`
	Content   string `json:"content" binding:"required"`
	EntryDate string `json:"entry_date" binding:"required"` // Format: YYYY-MM-DD
}

type JournalQuery struct {
	Search string
	Page   int
	Limit  int
} 
