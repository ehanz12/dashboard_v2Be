package mappers

import (
	"be_dashboard/dto/responses"
	"be_dashboard/models"
	"encoding/json"
)

func ToResponseHabitMini(h models.Habits) responses.HabitMini {
	var days []string
	if len(h.Days) > 0 {
		_ = json.Unmarshal(h.Days, &days)
	}

	return responses.HabitMini{
		ID:              h.ID,
		Name:            h.Name,
		Frequency:       h.Frequency,
		Days:            days,
		ReminderTime:    h.ReminderTime,
		ReminderEnabled: h.ReminderEnabled,
		CreatedAt:       h.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}

func ToListResponseHabitMini(habits []models.Habits) []responses.HabitMini {
	res := make([]responses.HabitMini, 0, len(habits))
	for _, h := range habits {
		res = append(res, ToResponseHabitMini(h))
	}
	return res
}

func ToResponseHabitLog(h models.HabitLogs) responses.HabitResponseLog {
	return responses.HabitResponseLog{
		ID: h.ID,
		Habit: responses.HabitSuperMini{
			ID:   h.Habits.ID,
			Name: h.Habits.Name,
		},
		Date:       h.LogDate.Format("2006-01-02"),
		IsComplete: h.Completed,
	}
}

func ListToResponseHabitLog(habits []models.HabitLogs) []responses.HabitResponseLog {
	res := make([]responses.HabitResponseLog, 0, len(habits))
	for _, h := range habits {
		res = append(res, ToResponseHabitLog(h))
	}
	return res
}
