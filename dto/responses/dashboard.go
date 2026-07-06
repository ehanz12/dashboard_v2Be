package responses

type DashboardResponse struct {
	Finance  FinanceDashboard  `json:"finance"`
	Habits   HabitsDashboard   `json:"habits"`
	Tasks    TasksDashboard    `json:"tasks"`
	Schedule ScheduleDashboard `json:"schedule"`
}

type FinanceDashboard struct {
	Balance float64 `json:"balance"`
	Income  float64 `json:"income"`
	Expense float64 `json:"expense"`
	Saving  float64 `json:"saving"`
}

type HabitsDashboard struct {
	CompletedToday int `json:"completed_today"`
	TotalToday     int `json:"total_today"`
	CurrentStreak  int `json:"current_streak"`
	LongestStreak  int `json:"longest_streak"`
	CompletionRate int `json:"completion_rate"`
	CompletedDays  int `json:"completed_days"`
}

type TasksDashboard struct {
	CompletedToday int `json:"completed_today"`
	TotalToday     int `json:"total_today"`
	CompletionRate int `json:"completion_rate"`
	Overdue        int `json:"overdue"`
}

type ScheduleDashboard struct {
	TodayEvents int    `json:"today_events"`
	NextEvent   string `json:"next_event"`
	NextTime    string `json:"next_time"`
}
