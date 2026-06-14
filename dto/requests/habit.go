package requests

type CreateHabitRequest struct {
	Name      string   `json:"name" binding:"required"`
	Frequency string   `json:"frequency" binding:"required"`
	Days      []string `json:"days,omitempty"`
}

type HabitQuery struct {
	Search string
	Page   int
	Limit  int
}

type ToogleHabitLogRequest struct {
	HabitID string `json:"habit_id" binding:"required"`
}
