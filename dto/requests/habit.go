package requests

type CreateHabitRequest struct {
	Name      string  `json:"name" binding:"required"`
	Frequency string  `json:"frequency" binding:"required"`
	Date      *string `json:"date"`
}

type HabitQuery struct {
	Search string
	Page   int
	Limit  int
}

type ToogleHabitLogRequest struct {
	HabitID string `json:"habit_id" binding:"required"`
}
