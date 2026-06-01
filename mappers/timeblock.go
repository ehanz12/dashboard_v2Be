package mappers

import (
	"be_dashboard/dto/responses"
	"be_dashboard/models"
)

func ToTimeblockResponse(t models.Timeblock) responses.TimeblockResponse {
	return responses.TimeblockResponse{
		ID:           t.ID,
		UserID:       t.UserID,
		ActivityName: t.ActivityName,
		StartTime:    t.StartTime,
		EndTime:      t.EndTime,
		ColorCode:    t.ColorCode,
		DayOfWeek:    t.DayOfWeek,
		CreatedAt:    t.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:    t.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

func ToListTimeblockResponse(timeblocks []models.Timeblock) []responses.TimeblockResponse {
	res := make([]responses.TimeblockResponse, 0, len(timeblocks))
	for _, t := range timeblocks {
		res = append(res, ToTimeblockResponse(t))
	}
	return res
}
