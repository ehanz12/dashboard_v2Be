package responses

type TaskResponse struct {
	ID          string  `json:"id"`
	UserID      string  `json:"user_id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Quadrant    int     `json:"quadrant"`
	IsCompleted bool    `json:"is_completed"`
	DueDate     *string `json:"due_date"` // Optional format: YYYY-MM-DD HH:MM:SS
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}
