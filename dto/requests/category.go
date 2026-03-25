package requests

type CategoryReq struct {
	UserID string `json:"user_id" validate:"required"`
	Name   string `json:"name" validate:"required"`
}
