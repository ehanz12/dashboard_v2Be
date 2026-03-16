package mappers

import (
	"be_dashboard/dto/responses"
	"be_dashboard/models"
)

func ToCategoryRes(c *models.Categories) *responses.CategoryResponse {
	return &responses.CategoryResponse{
		ID:     c.ID,
		UserID: c.UserID,
		Name:   c.Name,
		Type:   c.Type,
	}
}

func ListToCategoryRes(c []models.Categories) []responses.CategoryResponse {
	res := make([]responses.CategoryResponse, 0, len(c))
	for _, category := range c {
		res = append(res, *ToCategoryRes(&category))
	}
	return res
}
