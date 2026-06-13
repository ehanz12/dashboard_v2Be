package responses

type TimeblockResponse struct {
	ID           string `json:"id"`
	UserID       string `json:"user_id"`
	ActivityName string `json:"activity_name"`
	StartTime    string `json:"start_time"`
	EndTime      string `json:"end_time"`
	ColorCode    string `json:"color_code"`
	DayOfWeek    int    `json:"day_of_week"`
	Date         string `json:"date,omitempty"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}
