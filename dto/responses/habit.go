package responses

type HabitResponse struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	Name      string `json:"name"`
	Frequency string `json:"frequency"`
	CreatedAt string `json:"created_at"`
}

type HabitMini struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Frequency string `json:"frequency"`
	CreatedAt string `json:"created_at"`
}
