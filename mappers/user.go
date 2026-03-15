package mappers

import (
	"be_dashboard/dto/responses"
	"be_dashboard/models"
)

func ToUserResponse(u models.Users) responses.UserResponse {
	return responses.UserResponse{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
	}
}