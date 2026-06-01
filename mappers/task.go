package mappers

import (
	"be_dashboard/dto/responses"
	"be_dashboard/models"
)

func ToTaskResponse(t models.Task) responses.TaskResponse {
	var dueDateStr *string
	if t.DueDate != nil {
		formatted := t.DueDate.Format("2006-01-02 15:04:05")
		dueDateStr = &formatted
	}

	return responses.TaskResponse{
		ID:          t.ID,
		UserID:      t.UserID,
		Title:       t.Title,
		Description: t.Description,
		Quadrant:    t.Quadrant,
		IsCompleted: t.IsCompleted,
		DueDate:     dueDateStr,
		CreatedAt:   t.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   t.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

func ToListTaskResponse(tasks []models.Task) []responses.TaskResponse {
	res := make([]responses.TaskResponse, 0, len(tasks))
	for _, t := range tasks {
		res = append(res, ToTaskResponse(t))
	}
	return res
}
