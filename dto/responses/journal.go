package responses

type JournalResponse struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	Mood      int    `json:"mood"`
	Content   string `json:"content"`
	EntryDate string `json:"entry_date"` // Format: YYYY-MM-DD
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
