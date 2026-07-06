package services

import (
	"be_dashboard/database"
	"be_dashboard/dto/responses"
	"be_dashboard/models"
	"fmt"
	"time"
)

func GetDashboardData(userID string) (responses.DashboardResponse, error) {
	// Initialize response
	var res responses.DashboardResponse

	// ==========================================
	// 1. Finance Section
	// ==========================================
	var income, expense float64

	if err := database.DB.Model(&models.Transactions{}).
		Where("user_id = ? AND type = ?", userID, "income").
		Select("COALESCE(SUM(amount), 0)").
		Scan(&income).Error; err != nil {
		return responses.DashboardResponse{}, err
	}

	if err := database.DB.Model(&models.Transactions{}).
		Where("user_id = ? AND type = ?", userID, "expense").
		Select("COALESCE(SUM(amount), 0)").
		Scan(&expense).Error; err != nil {
		return responses.DashboardResponse{}, err
	}

	res.Finance = responses.FinanceDashboard{
		Income:  income,
		Expense: expense,
		Balance: income - expense,
		Saving:  income - expense,
	}

	// ==========================================
	// 2. Habits Section
	// ==========================================
	todayStr := time.Now().Format("2006-01-02")
	yesterdayStr := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	weekday := time.Now().Weekday().String()
	jsonVal := fmt.Sprintf("\"%s\"", weekday)

	type habitTodayResult struct {
		HabitID   string `gorm:"column:habit_id"`
		Name      string `gorm:"column:name"`
		Completed bool   `gorm:"column:completed"`
	}
	var habitsToday []habitTodayResult

	err := database.DB.
		Table("habits").
		Select(`
			habits.id as habit_id,
			habits.name,
			COALESCE(habit_logs.completed, false) as completed
		`).
		Joins(`
			LEFT JOIN habit_logs 
			ON habit_logs.habit_id = habits.id 
			AND habit_logs.log_date = ?
		`, todayStr).
		Where("habits.user_id = ? AND (habits.frequency = 'daily' OR JSON_CONTAINS(habits.days, ?))", userID, jsonVal).
		Scan(&habitsToday).Error

	if err != nil {
		return responses.DashboardResponse{}, err
	}

	completedToday := 0
	for _, h := range habitsToday {
		if h.Completed {
			completedToday++
		}
	}

	var completedDays int64
	if err := database.DB.Table("habit_logs").
		Joins("JOIN habits ON habit_logs.habit_id = habits.id").
		Where("habits.user_id = ? AND habit_logs.completed = ?", userID, true).
		Count(&completedDays).Error; err != nil {
		return responses.DashboardResponse{}, err
	}

	var totalLogs int64
	if err := database.DB.Table("habit_logs").
		Joins("JOIN habits ON habit_logs.habit_id = habits.id").
		Where("habits.user_id = ?", userID).
		Count(&totalLogs).Error; err != nil {
		return responses.DashboardResponse{}, err
	}

	habitCompletionRate := 0
	if totalLogs > 0 {
		habitCompletionRate = int(float64(completedDays) / float64(totalLogs) * 100)
	}

	// Fetch unique active dates to compute streak
	var activeDates []time.Time
	if err := database.DB.Table("habit_logs").
		Select("DISTINCT habit_logs.log_date").
		Joins("JOIN habits ON habit_logs.habit_id = habits.id").
		Where("habits.user_id = ? AND habit_logs.completed = ?", userID, true).
		Order("habit_logs.log_date DESC").
		Scan(&activeDates).Error; err != nil {
		return responses.DashboardResponse{}, err
	}

	currentStreak := 0
	longestStreak := 0

	var dateStrings []string
	for _, d := range activeDates {
		dateStrings = append(dateStrings, d.Format("2006-01-02"))
	}

	dateMap := make(map[string]bool)
	for _, s := range dateStrings {
		dateMap[s] = true
	}

	if len(dateStrings) > 0 {
		if dateMap[todayStr] {
			currentStreak = 1
			checkDate := time.Now().AddDate(0, 0, -1)
			for {
				checkStr := checkDate.Format("2006-01-02")
				if dateMap[checkStr] {
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
				checkStr := checkDate.Format("2006-01-02")
				if dateMap[checkStr] {
					currentStreak++
					checkDate = checkDate.AddDate(0, 0, -1)
				} else {
					break
				}
			}
		}

		// Calculate longest streak
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

	res.Habits = responses.HabitsDashboard{
		CompletedToday: completedToday,
		TotalToday:     len(habitsToday),
		CurrentStreak:  currentStreak,
		LongestStreak:  longestStreak,
		CompletionRate: habitCompletionRate,
		CompletedDays:  int(completedDays),
	}

	// ==========================================
	// 3. Tasks Section
	// ==========================================
	todayStart := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Now().Location())
	todayEnd := todayStart.Add(24 * time.Hour)

	var tasksCompletedToday int64
	if err := database.DB.Model(&models.Task{}).
		Where("user_id = ? AND is_completed = ? AND updated_at >= ? AND updated_at < ?", userID, true, todayStart, todayEnd).
		Count(&tasksCompletedToday).Error; err != nil {
		return responses.DashboardResponse{}, err
	}

	var tasksTotalToday int64
	if err := database.DB.Model(&models.Task{}).
		Where("user_id = ? AND due_date >= ? AND due_date < ?", userID, todayStart, todayEnd).
		Count(&tasksTotalToday).Error; err != nil {
		return responses.DashboardResponse{}, err
	}

	var completedCount int64
	var totalCount int64
	if err := database.DB.Model(&models.Task{}).Where("user_id = ?", userID).Count(&totalCount).Error; err != nil {
		return responses.DashboardResponse{}, err
	}
	if err := database.DB.Model(&models.Task{}).Where("user_id = ? AND is_completed = ?", userID, true).Count(&completedCount).Error; err != nil {
		return responses.DashboardResponse{}, err
	}
	taskCompletionRate := 0
	if totalCount > 0 {
		taskCompletionRate = int(float64(completedCount) / float64(totalCount) * 100)
	}

	var overdueCount int64
	if err := database.DB.Model(&models.Task{}).
		Where("user_id = ? AND is_completed = ? AND due_date < ?", userID, false, todayStart).
		Count(&overdueCount).Error; err != nil {
		return responses.DashboardResponse{}, err
	}

	res.Tasks = responses.TasksDashboard{
		CompletedToday: int(tasksCompletedToday),
		TotalToday:     int(tasksTotalToday),
		CompletionRate: taskCompletionRate,
		Overdue:        int(overdueCount),
	}

	// ==========================================
	// 4. Schedule Section
	// ==========================================
	parsedDate := time.Now()
	dayOfWeek := int(parsedDate.Weekday())
	if dayOfWeek == 0 {
		dayOfWeek = 7
	}
	timeblocks, err := GetTimeblocksByUserID(userID, &dayOfWeek, &parsedDate)
	if err != nil {
		return responses.DashboardResponse{}, err
	}

	currentTimeStr := time.Now().Format("15:04")
	nextEventName := "-"
	nextEventTime := "-"

	for _, tb := range timeblocks {
		if tb.StartTime > currentTimeStr {
			nextEventName = tb.ActivityName
			nextEventTime = tb.StartTime
			break
		}
	}

	res.Schedule = responses.ScheduleDashboard{
		TodayEvents: len(timeblocks),
		NextEvent:   nextEventName,
		NextTime:    nextEventTime,
	}

	return res, nil
}
