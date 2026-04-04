package mappers

import (
	"be_dashboard/dto/responses"
	"be_dashboard/models"
)

func ToResponseHabitMini(h models.Habits) responses.HabitMini {
	return responses.HabitMini{
		ID:   h.ID,
		Name: h.Name,
		Frequency: h.Frequency,
		CreatedAt: h.CreatedAt.Format("2006-01-02 15:04:05"),
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
			ID: h.Habits.ID,
			Name: h.Habits.Name,
		},
		Date: h.LogDate.Format("2006-01-02"),
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