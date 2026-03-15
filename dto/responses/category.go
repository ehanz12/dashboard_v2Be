package responses

type CategoryResponse struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`
	Name   string `json:"name"`
	Type   string `json:"type"`
}
