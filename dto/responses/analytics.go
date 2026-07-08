package responses

// ─────────────────────────────────────────────────────────────────────────────
// Summary
// ─────────────────────────────────────────────────────────────────────────────

// AnalyticsSummaryResponse adalah response untuk GET /analytics/summary.
type AnalyticsSummaryResponse struct {
	Finance AnalyticsFinanceSummary `json:"finance"`
	Habit   AnalyticsHabitSummary   `json:"habit"`
	Task    AnalyticsTaskSummary    `json:"task"`
}

type AnalyticsFinanceSummary struct {
	Income  float64 `json:"income"`
	Expense float64 `json:"expense"`
	Balance float64 `json:"balance"`
}

type AnalyticsHabitSummary struct {
	TodayCompleted int `json:"today_completed"`
	TodayTotal     int `json:"today_total"`
	CompletionRate int `json:"completion_rate"`
	CurrentStreak  int `json:"current_streak"`
	LongestStreak  int `json:"longest_streak"`
}

type AnalyticsTaskSummary struct {
	Completed      int `json:"completed"`
	Pending        int `json:"pending"`
	CompletionRate int `json:"completion_rate"`
}

// ─────────────────────────────────────────────────────────────────────────────
// Finance Analytics
// ─────────────────────────────────────────────────────────────────────────────

// FinanceAnalyticsPoint adalah satu titik data finance.
// Untuk monthly: label = "1 Jul", "2 Jul", …
// Untuk yearly:  label = "Jan", "Feb", …
type FinanceAnalyticsPoint struct {
	Label   string  `json:"label"`
	Income  float64 `json:"income"`
	Expense float64 `json:"expense"`
}

// ─────────────────────────────────────────────────────────────────────────────
// Category Analytics
// ─────────────────────────────────────────────────────────────────────────────

// CategoryAnalyticsResponse adalah response untuk GET /analytics/categories.
type CategoryAnalyticsResponse struct {
	Category   string  `json:"category"`
	Amount     float64 `json:"amount"`
	Percentage float64 `json:"percentage"`
}

// ─────────────────────────────────────────────────────────────────────────────
// Habit Analytics
// ─────────────────────────────────────────────────────────────────────────────

// HabitAnalyticsPoint adalah satu titik data habit.
// Untuk monthly: label = "1", "2", …  (nomor hari)
// Untuk yearly:  label = "Jan", "Feb", …
type HabitAnalyticsPoint struct {
	Label      string `json:"label"`
	Completed  int    `json:"completed"`
	Total      int    `json:"total"`
	Percentage int    `json:"percentage"`
}

// ─────────────────────────────────────────────────────────────────────────────
// Task Analytics
// ─────────────────────────────────────────────────────────────────────────────

// TaskAnalyticsPoint adalah satu titik data task.
// Untuk monthly: label = "1", "2", …
// Untuk yearly:  label = "Jan", "Feb", …
type TaskAnalyticsPoint struct {
	Label     string `json:"label"`
	Completed int    `json:"completed"`
}

// ─────────────────────────────────────────────────────────────────────────────
// Streak Analytics
// ─────────────────────────────────────────────────────────────────────────────

// StreakAnalyticsResponse adalah response untuk GET /analytics/streak.
type StreakAnalyticsResponse struct {
	CurrentStreak  int `json:"current_streak"`
	LongestStreak  int `json:"longest_streak"`
	CompletionRate int `json:"completion_rate"`
}
