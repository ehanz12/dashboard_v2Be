package requests

type CreateHabitRequest struct {
	Name      string `json:"name" binding:"required"`
	Frequency string `json:"frequency" binding:"required"`
}

type HabitQuery struct {
	Page  int
	Limit int
}
