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