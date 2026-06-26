package cron

import (
	"encoding/json"
	"log"
	"strings"
	"time"

	"be_dashboard/database"
	"be_dashboard/models"
	"be_dashboard/services"
)

func isTodayIncluded(daysJSON string) bool {

	var days []string

	if err := json.Unmarshal(
		[]byte(daysJSON),
		&days,
	); err != nil {
		return false
	}

	today := strings.ToUpper(
		time.Now().Weekday().String()[:3],
	)

	for _, day := range days {

		if day == today {
			return true
		}
	}

	return false
}

func CheckHabitReminders() {

	now := time.Now().Format("15:04")

	var habits []models.Habits

	database.DB.
		Where("reminder_enabled = ?", true).
		Find(&habits)

	for _, habit := range habits {

		if habit.ReminderTime == nil {
			continue
		}

		habitTime :=
			habit.ReminderTime.Format("15:04")

		if habitTime != now {
			continue
		}

		// Daily habits selalu aktif setiap hari, weekly perlu cek hari
		if habit.Frequency == "weekly" && !isTodayIncluded(habit.Days.String()) {
			continue
		}

		var devices []models.UserDevice

		database.DB.
			Where(
				"user_id = ?",
				habit.UserID,
			).
			Find(&devices)

		for _, device := range devices {

			err := services.SendPushNotification(
				device.FMCToken,
				"Reminder Habit",
				"Jangan lupa "+habit.Name,
			)

			if err != nil {
				log.Println(err)
			}
		}
	}
}

