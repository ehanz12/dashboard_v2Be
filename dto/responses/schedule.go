package responses

type ScheduleTodayResponse struct {
	Date       string                  `json:"date"`
	DayOfWeek  int                     `json:"day_of_week"`
	Timeblocks []TimeblockResponse     `json:"timeblocks"`
	Habits     []HabitLogTodayResponse `json:"habits"`
}
