package responses

type CategoryResponse struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`
	Name   string `json:"name"`
}

type CategoryMiniResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
