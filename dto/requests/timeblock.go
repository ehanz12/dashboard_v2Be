package requests

type CreateTimeblockRequest struct {
	ActivityName string `json:"activity_name" binding:"required"`
	StartTime    string `json:"start_time" binding:"required"` // "HH:MM"
	EndTime      string `json:"end_time" binding:"required"`   // "HH:MM"
	ColorCode    string `json:"color_code"`                    // Optional hex code
	DayOfWeek    int    `json:"day_of_week" binding:"required"` // 1-7 (Monday-Sunday)
}
