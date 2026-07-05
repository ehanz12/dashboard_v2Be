package responses

type HabitResponse struct {
	ID              string   `json:"id"`
	UserID          string   `json:"user_id"`
	Name            string   `json:"name"`
	Frequency       string   `json:"frequency"`
	Days            []string `json:"days,omitempty"`
	ReminderTime    *string  `json:"reminder_time,omitempty"`
	ReminderEnabled bool     `json:"reminder_enabled"`
	CreatedAt       string   `json:"created_at"`
}

type HabitMini struct {
	ID              string   `json:"id"`
	Name            string   `json:"name"`
	Frequency       string   `json:"frequency"`
	Days            []string `json:"days,omitempty"`
	ReminderTime    *string  `json:"reminder_time,omitempty"`
	ReminderEnabled bool     `json:"reminder_enabled"`
	CreatedAt       string   `json:"created_at"`
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

type HabitSummaryResponse struct {
	HabitID        string  `json:"habit_id"`
	CurrentStreak  int     `json:"current_streak"`
	LongestStreak  int     `json:"longest_streak"`
	CompletedDays  int     `json:"completed_days"`
	CompletionRate float64 `json:"completion_rate"`
}
