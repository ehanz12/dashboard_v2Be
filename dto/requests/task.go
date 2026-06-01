package requests

type CreateTaskRequest struct {
	Title       string  `json:"title" binding:"required"`
	Description string  `json:"description"`
	Quadrant    int     `json:"quadrant"` // 1, 2, 3, or 4
	DueDate     *string `json:"due_date"` // Optional format: RFC3339 or "2006-01-02 15:04:05"
}

type UpdateTaskRequest struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Quadrant    int     `json:"quadrant"`
	IsCompleted *bool   `json:"is_completed"`
	DueDate     *string `json:"due_date"`
}
