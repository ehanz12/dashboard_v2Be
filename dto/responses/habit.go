package responses

type HabitResponse struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	Name      string `json:"name"`
	Frequency string `json:"frequency"`
	Date      string `json:"date,omitempty"`
	CreatedAt string `json:"created_at"`
}

type HabitMini struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Frequency string `json:"frequency"`
	Date      string `json:"date,omitempty"`
	CreatedAt string `json:"created_at"`
}

type HabitSuperMini struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type HabitResponseLog struct {
	ID         string         `json:"id"`
	Habit      HabitSuperMini `json:"habit"`
	Date       string         `json:"date"`
	IsComplete bool           `json:"is_complete"`
}

type HabitLogTodayResponse struct {
	HabitID   string `json:"habit_id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}
