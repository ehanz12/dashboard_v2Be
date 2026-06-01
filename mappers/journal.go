package mappers

import (
	"be_dashboard/dto/responses"
	"be_dashboard/models"
)

func ToJournalResponse(j models.Journal) responses.JournalResponse {
	return responses.JournalResponse{
		ID:        j.ID,
		UserID:    j.UserID,
		Mood:      j.Mood,
		Content:   j.Content,
		EntryDate: j.EntryDate.Format("2006-01-02"),
		CreatedAt: j.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: j.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

func ToListJournalResponse(journals []models.Journal) []responses.JournalResponse {
	res := make([]responses.JournalResponse, 0, len(journals))
	for _, j := range journals {
		res = append(res, ToJournalResponse(j))
	}
	return res
}
