package mappers

import (
	"be_dashboard/dto/responses"
	"be_dashboard/models"
)

func ToUserResponse(u models.Users) responses.UserResponse {
	bio := ""
	if u.Bio != nil {
		bio = *u.Bio
	}
	return responses.UserResponse{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		NomorHP:   u.NomorHP,
		Bio:       bio,
		CreatedAt: u.CreatedAt,
	}
}