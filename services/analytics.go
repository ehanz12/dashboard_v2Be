package services

import (
	"be_dashboard/database"
	"be_dashboard/dto/requests"
	"be_dashboard/dto/responses"
	"fmt"
	"math"
	"time"
)

// shortMonthNames digunakan untuk label pada data tahunan.
var shortMonthNames = []string{
	"Jan", "Feb", "Mar", "Apr", "May", "Jun",
	"Jul", "Aug", "Sep", "Oct", "Nov", "Dec",
}

// ─────────────────────────────────────────────────────────────────────────────
// 1. Summary
// ─────────────────────────────────────────────────────────────────────────────

// GetAnalyticsSummary mengembalikan ringkasan finance, habit, dan task
// dalam rentang waktu yang ditentukan oleh filter.
func GetAnalyticsSummary(userID string, filter requests.AnalyticsFilter) (responses.AnalyticsSummaryResponse, error) {
	var res responses.AnalyticsSummaryResponse
	start, end := filter.DateRange()

	// ── Finance ──────────────────────────────────────────────────────────────
	type financeSumResult struct {
		Income  float64 `gorm:"column:income"`
		Expense float64 `gorm:"column:expense"`
	}
	var finSum financeSumResult

	if err := database.DB.Raw(`
		SELECT
			COALESCE(SUM(CASE WHEN type = 'income'  THEN amount ELSE 0 END), 0) AS income,
			COALESCE(SUM(CASE WHEN type = 'expense' THEN amount ELSE 0 END), 0) AS expense
		FROM transactions
		WHERE user_id = ?
		  AND transaction_date >= ?
		  AND transaction_date < ?
	`, userID, start, end).Scan(&finSum).Error; err != nil {
		return responses.AnalyticsSummaryResponse{}, err
	}

	res.Finance = responses.AnalyticsFinanceSummary{
		Income:  finSum.Income,
		Expense: finSum.Expense,
		Balance: finSum.Income - finSum.Expense,
	}

	// ── Habit ────────────────────────────────────────────────────────────────
	type habitSumResult struct {
		Completed int64 `gorm:"column:completed_count"`
		Total     int64 `gorm:"column:total_count"`
	}
	var habitSum habitSumResult

	if err := database.DB.Raw(`
		SELECT
			COALESCE(SUM(CASE WHEN hl.completed = true THEN 1 ELSE 0 END), 0) AS completed_count,
			COUNT(hl.id)                                                         AS total_count
		FROM habit_logs hl
		JOIN habits h ON hl.habit_id = h.id
		WHERE h.user_id = ?
		  AND hl.log_date >= ?
		  AND hl.log_date < ?
	`, userID, start, end).Scan(&habitSum).Error; err != nil {
		return responses.AnalyticsSummaryResponse{}, err
	}

	habitRate := 0
	if habitSum.Total > 0 {
		habitRate = int(math.Round(float64(habitSum.Completed) / float64(habitSum.Total) * 100))
	}

	// Streak: hitung berdasarkan seluruh data global (bukan hanya period filter)
	currentStreak, longestStreak, err := calcHabitStreak(userID)
	if err != nil {
		return responses.AnalyticsSummaryResponse{}, err
	}

	// Habit hari ini
	todayStr := time.Now().Format("2006-01-02")
	weekday := time.Now().Weekday().String()
	jsonVal := fmt.Sprintf("%q", weekday)

	type habitTodayResult struct {
		Completed bool `gorm:"column:completed"`
	}
	var habitsToday []habitTodayResult

	if err := database.DB.Raw(`
		SELECT COALESCE(hl.completed, false) AS completed
		FROM habits h
		LEFT JOIN habit_logs hl
			ON hl.habit_id = h.id AND hl.log_date = ?
		WHERE h.user_id = ?
		  AND (h.frequency = 'daily' OR JSON_CONTAINS(h.days, ?))
	`, todayStr, userID, jsonVal).Scan(&habitsToday).Error; err != nil {
		return responses.AnalyticsSummaryResponse{}, err
	}

	todayCompleted := 0
	for _, h := range habitsToday {
		if h.Completed {
			todayCompleted++
		}
	}

	res.Habit = responses.AnalyticsHabitSummary{
		TodayCompleted: todayCompleted,
		TodayTotal:     len(habitsToday),
		CompletionRate: habitRate,
		CurrentStreak:  currentStreak,
		LongestStreak:  longestStreak,
	}

	// ── Task ─────────────────────────────────────────────────────────────────
	type taskSumResult struct {
		Completed int64 `gorm:"column:completed_count"`
		Total     int64 `gorm:"column:total_count"`
	}
	var taskSum taskSumResult

	if err := database.DB.Raw(`
		SELECT
			COUNT(*)                                                  AS total_count,
			COALESCE(SUM(CASE WHEN is_completed = true THEN 1 ELSE 0 END), 0) AS completed_count
		FROM tasks
		WHERE user_id = ?
		  AND created_at >= ?
		  AND created_at < ?
	`, userID, start, end).Scan(&taskSum).Error; err != nil {
		return responses.AnalyticsSummaryResponse{}, err
	}

	taskRate := 0
	if taskSum.Total > 0 {
		taskRate = int(math.Round(float64(taskSum.Completed) / float64(taskSum.Total) * 100))
	}

	res.Task = responses.AnalyticsTaskSummary{
		Completed:      int(taskSum.Completed),
		Pending:        int(taskSum.Total - taskSum.Completed),
		CompletionRate: taskRate,
	}

	return res, nil
}

// ─────────────────────────────────────────────────────────────────────────────
// 2. Finance Analytics
// ─────────────────────────────────────────────────────────────────────────────

// GetFinanceAnalytics mengembalikan data income & expense per hari (monthly)
// atau per bulan (yearly), dengan nilai 0 untuk hari/bulan tanpa transaksi.
func GetFinanceAnalytics(userID string, filter requests.AnalyticsFilter) ([]responses.FinanceAnalyticsPoint, error) {
	start, end := filter.DateRange()

	if filter.Type == "monthly" {
		return financeMonthly(userID, filter, start, end)
	}
	return financeYearly(userID, filter, start, end)
}

func financeMonthly(userID string, filter requests.AnalyticsFilter, start, end time.Time) ([]responses.FinanceAnalyticsPoint, error) {
	type row struct {
		Day     int     `gorm:"column:day"`
		Income  float64 `gorm:"column:income"`
		Expense float64 `gorm:"column:expense"`
	}
	var rows []row

	if err := database.DB.Raw(`
		SELECT
			DAY(transaction_date)                                                         AS day,
			COALESCE(SUM(CASE WHEN type = 'income'  THEN amount ELSE 0 END), 0)          AS income,
			COALESCE(SUM(CASE WHEN type = 'expense' THEN amount ELSE 0 END), 0)          AS expense
		FROM transactions
		WHERE user_id = ?
		  AND transaction_date >= ?
		  AND transaction_date < ?
		GROUP BY DAY(transaction_date)
		ORDER BY day
	`, userID, start, end).Scan(&rows).Error; err != nil {
		return nil, err
	}

	// Map hari → data
	dayMap := make(map[int]row, len(rows))
	for _, r := range rows {
		dayMap[r.Day] = r
	}

	days := filter.DaysInMonth()
	monthAbbr := shortMonthNames[filter.Month-1]
	result := make([]responses.FinanceAnalyticsPoint, days)

	for d := 1; d <= days; d++ {
		label := fmt.Sprintf("%d %s", d, monthAbbr)
		point := responses.FinanceAnalyticsPoint{Label: label}
		if r, ok := dayMap[d]; ok {
			point.Income = r.Income
			point.Expense = r.Expense
		}
		result[d-1] = point
	}

	return result, nil
}

func financeYearly(userID string, filter requests.AnalyticsFilter, start, end time.Time) ([]responses.FinanceAnalyticsPoint, error) {
	type row struct {
		Month   int     `gorm:"column:month"`
		Income  float64 `gorm:"column:income"`
		Expense float64 `gorm:"column:expense"`
	}
	var rows []row

	if err := database.DB.Raw(`
		SELECT
			MONTH(transaction_date)                                                       AS month,
			COALESCE(SUM(CASE WHEN type = 'income'  THEN amount ELSE 0 END), 0)          AS income,
			COALESCE(SUM(CASE WHEN type = 'expense' THEN amount ELSE 0 END), 0)          AS expense
		FROM transactions
		WHERE user_id = ?
		  AND transaction_date >= ?
		  AND transaction_date < ?
		GROUP BY MONTH(transaction_date)
		ORDER BY month
	`, userID, start, end).Scan(&rows).Error; err != nil {
		return nil, err
	}

	monthMap := make(map[int]row, len(rows))
	for _, r := range rows {
		monthMap[r.Month] = r
	}

	result := make([]responses.FinanceAnalyticsPoint, 12)
	for m := 1; m <= 12; m++ {
		point := responses.FinanceAnalyticsPoint{Label: shortMonthNames[m-1]}
		if r, ok := monthMap[m]; ok {
			point.Income = r.Income
			point.Expense = r.Expense
		}
		result[m-1] = point
	}

	return result, nil
}

// ─────────────────────────────────────────────────────────────────────────────
// 3. Category Analytics
// ─────────────────────────────────────────────────────────────────────────────

// GetCategoryAnalytics mengembalikan breakdown expense per kategori beserta persentasenya.
func GetCategoryAnalytics(userID string, filter requests.AnalyticsFilter) ([]responses.CategoryAnalyticsResponse, error) {
	start, end := filter.DateRange()

	type row struct {
		Category string  `gorm:"column:category"`
		Amount   float64 `gorm:"column:amount"`
	}
	var rows []row

	if err := database.DB.Raw(`
		SELECT
			COALESCE(c.name, 'Uncategorized') AS category,
			COALESCE(SUM(t.amount), 0)        AS amount
		FROM transactions t
		LEFT JOIN categories c ON t.category_id = c.id
		WHERE t.user_id = ?
		  AND t.type = 'expense'
		  AND t.transaction_date >= ?
		  AND t.transaction_date < ?
		GROUP BY t.category_id, c.name
		ORDER BY amount DESC
	`, userID, start, end).Scan(&rows).Error; err != nil {
		return nil, err
	}

	// Hitung total expense untuk persentase
	var total float64
	for _, r := range rows {
		total += r.Amount
	}

	result := make([]responses.CategoryAnalyticsResponse, len(rows))
	for i, r := range rows {
		pct := 0.0
		if total > 0 {
			pct = math.Round(r.Amount/total*10000) / 100 // dua desimal
		}
		result[i] = responses.CategoryAnalyticsResponse{
			Category:   r.Category,
			Amount:     r.Amount,
			Percentage: pct,
		}
	}

	return result, nil
}

// ─────────────────────────────────────────────────────────────────────────────
// 4. Habit Analytics
// ─────────────────────────────────────────────────────────────────────────────

// GetHabitAnalytics mengembalikan completion habit per hari (monthly) atau per bulan (yearly).
func GetHabitAnalytics(userID string, filter requests.AnalyticsFilter) ([]responses.HabitAnalyticsPoint, error) {
	start, end := filter.DateRange()

	if filter.Type == "monthly" {
		return habitMonthly(userID, filter, start, end)
	}
	return habitYearly(userID, filter, start, end)
}

func habitMonthly(userID string, filter requests.AnalyticsFilter, start, end time.Time) ([]responses.HabitAnalyticsPoint, error) {
	type row struct {
		Day       int `gorm:"column:day"`
		Completed int `gorm:"column:completed_count"`
		Total     int `gorm:"column:total_count"`
	}
	var rows []row

	if err := database.DB.Raw(`
		SELECT
			DAY(hl.log_date)                                                AS day,
			COALESCE(SUM(CASE WHEN hl.completed = true THEN 1 ELSE 0 END), 0) AS completed_count,
			COUNT(hl.id)                                                    AS total_count
		FROM habit_logs hl
		JOIN habits h ON hl.habit_id = h.id
		WHERE h.user_id = ?
		  AND hl.log_date >= ?
		  AND hl.log_date < ?
		GROUP BY DAY(hl.log_date)
		ORDER BY day
	`, userID, start, end).Scan(&rows).Error; err != nil {
		return nil, err
	}

	dayMap := make(map[int]row, len(rows))
	for _, r := range rows {
		dayMap[r.Day] = r
	}

	days := filter.DaysInMonth()
	result := make([]responses.HabitAnalyticsPoint, days)

	for d := 1; d <= days; d++ {
		point := responses.HabitAnalyticsPoint{Label: fmt.Sprintf("%d", d)}
		if r, ok := dayMap[d]; ok {
			pct := 0
			if r.Total > 0 {
				pct = int(math.Round(float64(r.Completed) / float64(r.Total) * 100))
			}
			point.Completed = r.Completed
			point.Total = r.Total
			point.Percentage = pct
		}
		result[d-1] = point
	}

	return result, nil
}

func habitYearly(userID string, filter requests.AnalyticsFilter, start, end time.Time) ([]responses.HabitAnalyticsPoint, error) {
	type row struct {
		Month     int `gorm:"column:month"`
		Completed int `gorm:"column:completed_count"`
		Total     int `gorm:"column:total_count"`
	}
	var rows []row

	if err := database.DB.Raw(`
		SELECT
			MONTH(hl.log_date)                                              AS month,
			COALESCE(SUM(CASE WHEN hl.completed = true THEN 1 ELSE 0 END), 0) AS completed_count,
			COUNT(hl.id)                                                    AS total_count
		FROM habit_logs hl
		JOIN habits h ON hl.habit_id = h.id
		WHERE h.user_id = ?
		  AND hl.log_date >= ?
		  AND hl.log_date < ?
		GROUP BY MONTH(hl.log_date)
		ORDER BY month
	`, userID, start, end).Scan(&rows).Error; err != nil {
		return nil, err
	}

	monthMap := make(map[int]row, len(rows))
	for _, r := range rows {
		monthMap[r.Month] = r
	}

	result := make([]responses.HabitAnalyticsPoint, 12)
	for m := 1; m <= 12; m++ {
		point := responses.HabitAnalyticsPoint{Label: shortMonthNames[m-1]}
		if r, ok := monthMap[m]; ok {
			pct := 0
			if r.Total > 0 {
				pct = int(math.Round(float64(r.Completed) / float64(r.Total) * 100))
			}
			point.Completed = r.Completed
			point.Total = r.Total
			point.Percentage = pct
		}
		result[m-1] = point
	}

	return result, nil
}

// ─────────────────────────────────────────────────────────────────────────────
// 5. Task Analytics
// ─────────────────────────────────────────────────────────────────────────────

// GetTaskAnalytics mengembalikan jumlah task yang diselesaikan per hari (monthly)
// atau per bulan (yearly).
func GetTaskAnalytics(userID string, filter requests.AnalyticsFilter) ([]responses.TaskAnalyticsPoint, error) {
	start, end := filter.DateRange()

	if filter.Type == "monthly" {
		return taskMonthly(userID, filter, start, end)
	}
	return taskYearly(userID, filter, start, end)
}

func taskMonthly(userID string, filter requests.AnalyticsFilter, start, end time.Time) ([]responses.TaskAnalyticsPoint, error) {
	type row struct {
		Day       int `gorm:"column:day"`
		Completed int `gorm:"column:completed_count"`
	}
	var rows []row

	if err := database.DB.Raw(`
		SELECT
			DAY(updated_at)  AS day,
			COUNT(*)         AS completed_count
		FROM tasks
		WHERE user_id = ?
		  AND is_completed = true
		  AND updated_at >= ?
		  AND updated_at < ?
		GROUP BY DAY(updated_at)
		ORDER BY day
	`, userID, start, end).Scan(&rows).Error; err != nil {
		return nil, err
	}

	dayMap := make(map[int]int, len(rows))
	for _, r := range rows {
		dayMap[r.Day] = r.Completed
	}

	days := filter.DaysInMonth()
	result := make([]responses.TaskAnalyticsPoint, days)
	for d := 1; d <= days; d++ {
		result[d-1] = responses.TaskAnalyticsPoint{
			Label:     fmt.Sprintf("%d", d),
			Completed: dayMap[d],
		}
	}

	return result, nil
}

func taskYearly(userID string, filter requests.AnalyticsFilter, start, end time.Time) ([]responses.TaskAnalyticsPoint, error) {
	type row struct {
		Month     int `gorm:"column:month"`
		Completed int `gorm:"column:completed_count"`
	}
	var rows []row

	if err := database.DB.Raw(`
		SELECT
			MONTH(updated_at) AS month,
			COUNT(*)          AS completed_count
		FROM tasks
		WHERE user_id = ?
		  AND is_completed = true
		  AND updated_at >= ?
		  AND updated_at < ?
		GROUP BY MONTH(updated_at)
		ORDER BY month
	`, userID, start, end).Scan(&rows).Error; err != nil {
		return nil, err
	}

	monthMap := make(map[int]int, len(rows))
	for _, r := range rows {
		monthMap[r.Month] = r.Completed
	}

	result := make([]responses.TaskAnalyticsPoint, 12)
	for m := 1; m <= 12; m++ {
		result[m-1] = responses.TaskAnalyticsPoint{
			Label:     shortMonthNames[m-1],
			Completed: monthMap[m],
		}
	}

	return result, nil
}

// ─────────────────────────────────────────────────────────────────────────────
// 6. Streak Analytics
// ─────────────────────────────────────────────────────────────────────────────

// GetStreakAnalytics mengembalikan current_streak, longest_streak, dan completion_rate.
// Untuk type=monthly: completion_rate dihitung dalam bulan tersebut.
// Untuk type=yearly : completion_rate dihitung dalam tahun tersebut.
// current_streak & longest_streak selalu dihitung dari data global (konsisten dengan dashboard).
func GetStreakAnalytics(userID string, filter requests.AnalyticsFilter) (responses.StreakAnalyticsResponse, error) {
	start, end := filter.DateRange()

	// Completion rate dalam period filter
	type habitSumResult struct {
		Completed int64 `gorm:"column:completed_count"`
		Total     int64 `gorm:"column:total_count"`
	}
	var habitSum habitSumResult

	if err := database.DB.Raw(`
		SELECT
			COALESCE(SUM(CASE WHEN hl.completed = true THEN 1 ELSE 0 END), 0) AS completed_count,
			COUNT(hl.id)                                                        AS total_count
		FROM habit_logs hl
		JOIN habits h ON hl.habit_id = h.id
		WHERE h.user_id = ?
		  AND hl.log_date >= ?
		  AND hl.log_date < ?
	`, userID, start, end).Scan(&habitSum).Error; err != nil {
		return responses.StreakAnalyticsResponse{}, err
	}

	completionRate := 0
	if habitSum.Total > 0 {
		completionRate = int(math.Round(float64(habitSum.Completed) / float64(habitSum.Total) * 100))
	}

	// Streak global
	currentStreak, longestStreak, err := calcHabitStreak(userID)
	if err != nil {
		return responses.StreakAnalyticsResponse{}, err
	}

	return responses.StreakAnalyticsResponse{
		CurrentStreak:  currentStreak,
		LongestStreak:  longestStreak,
		CompletionRate: completionRate,
	}, nil
}

// ─────────────────────────────────────────────────────────────────────────────
// Helper: Hitung streak global habit
// ─────────────────────────────────────────────────────────────────────────────

// calcHabitStreak menghitung current_streak dan longest_streak berdasarkan
// seluruh habit_logs user (tanpa filter periode). Logika sama dengan dashboard.go.
func calcHabitStreak(userID string) (currentStreak, longestStreak int, err error) {
	var activeDates []time.Time

	if err = database.DB.Raw(`
		SELECT DISTINCT hl.log_date
		FROM habit_logs hl
		JOIN habits h ON hl.habit_id = h.id
		WHERE h.user_id = ? AND hl.completed = true
		ORDER BY hl.log_date DESC
	`, userID).Scan(&activeDates).Error; err != nil {
		return
	}

	if len(activeDates) == 0 {
		return 0, 0, nil
	}

	todayStr := time.Now().Format("2006-01-02")
	yesterdayStr := time.Now().AddDate(0, 0, -1).Format("2006-01-02")

	dateStrings := make([]string, len(activeDates))
	for i, d := range activeDates {
		dateStrings[i] = d.Format("2006-01-02")
	}

	dateMap := make(map[string]bool, len(dateStrings))
	for _, s := range dateStrings {
		dateMap[s] = true
	}

	// Current streak
	if dateMap[todayStr] {
		currentStreak = 1
		checkDate := time.Now().AddDate(0, 0, -1)
		for {
			if dateMap[checkDate.Format("2006-01-02")] {
				currentStreak++
				checkDate = checkDate.AddDate(0, 0, -1)
			} else {
				break
			}
		}
	} else if dateMap[yesterdayStr] {
		currentStreak = 1
		checkDate := time.Now().AddDate(0, 0, -2)
		for {
			if dateMap[checkDate.Format("2006-01-02")] {
				currentStreak++
				checkDate = checkDate.AddDate(0, 0, -1)
			} else {
				break
			}
		}
	}

	// Longest streak
	if len(dateStrings) > 0 {
		longestStreak = 1
		tempStreak := 1
		for i := 0; i < len(dateStrings)-1; i++ {
			curr, _ := time.Parse("2006-01-02", dateStrings[i])
			next, _ := time.Parse("2006-01-02", dateStrings[i+1])
			if curr.AddDate(0, 0, -1).Equal(next) {
				tempStreak++
			} else {
				if tempStreak > longestStreak {
					longestStreak = tempStreak
				}
				tempStreak = 1
			}
		}
		if tempStreak > longestStreak {
			longestStreak = tempStreak
		}
	}

	return currentStreak, longestStreak, nil
}
