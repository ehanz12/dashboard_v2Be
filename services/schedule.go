package services

import (
	"be_dashboard/dto/responses"
	"errors"
	"time"
)

func GetScheduleForDateService(userID, date string) (responses.ScheduleTodayResponse, error) {
	var parsedDate time.Time
	var err error
	if date == "" {
		parsedDate = time.Now()
	} else {
		parsedDate, err = time.Parse("2006-01-02", date)
		if err != nil {
			return responses.ScheduleTodayResponse{}, errors.New("invalid date format, use YYYY-MM-DD")
		}
	}

	dayOfWeek := int(parsedDate.Weekday())
	if dayOfWeek == 0 {
		dayOfWeek = 7
	}

	timeblocks, err := GetTimeblocksByUserID(userID, &dayOfWeek, &parsedDate)
	if err != nil {
		return responses.ScheduleTodayResponse{}, err
	}

	habits, err := GetHabitLogsByDateService(userID, parsedDate.Format("2006-01-02"))
	if err != nil {
		return responses.ScheduleTodayResponse{}, err
	}

	return responses.ScheduleTodayResponse{
		Date:       parsedDate.Format("2006-01-02"),
		DayOfWeek:  dayOfWeek,
		Timeblocks: timeblocks,
		Habits:     habits,
	}, nil
}
