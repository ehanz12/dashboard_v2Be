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

func isTodayIncluded(daysJSON string, now time.Time) bool {

	var days []string

	if err := json.Unmarshal([]byte(daysJSON), &days); err != nil {
		log.Println("Failed parse days:", err)
		return false
	}

	today := strings.ToUpper(now.Weekday().String()[:3])

	for _, day := range days {
		if strings.ToUpper(day) == today {
			return true
		}
	}

	return false
}

func CheckHabitReminders() {

	// ==========================
	// Gunakan timezone Indonesia
	// ==========================
	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		log.Println("Failed load timezone:", err)
		return
	}

	now := time.Now().In(location)
	currentTime := now.Format("15:04")

	// log.Println("========== HABIT REMINDER ==========")
	// log.Println("Current Time :", currentTime)

	var habits []models.Habits

	if err := database.DB.
		Where("reminder_enabled = ?", true).
		Find(&habits).Error; err != nil {

		log.Println("Failed get habits:", err)
		return
	}

	// log.Printf("Found %d reminder habits\n", len(habits))

	for _, habit := range habits {

		// log.Println("--------------------------------")
		// log.Println("Habit :", habit.Name)
		// log.Println("User  :", habit.UserID)

		if habit.ReminderTime == nil {
			log.Println("Skip -> reminder time is NULL")
			continue
		}

		habitTime := strings.TrimSpace(*habit.ReminderTime)

if len(habitTime) >= 5 {
	habitTime = habitTime[:5]
}

		// log.Println("Reminder Time :", habitTime)

		// ==========================
		// Cek jam reminder
		// ==========================
		if habitTime != currentTime {
			log.Println("Skip -> time not matched")
			continue
		}

		// ==========================
		// Weekly hanya jalan sesuai hari
		// ==========================
		if habit.Frequency == "weekly" {

			if !isTodayIncluded(habit.Days.String(), now) {
				log.Println("Skip -> today not included")
				continue
			}
		}

		var devices []models.UserDevice

		if err := database.DB.
			Where("user_id = ?", habit.UserID).
			Find(&devices).Error; err != nil {

			log.Println("Failed get devices:", err)
			continue
		}

		// log.Printf("Found %d devices\n", len(devices))

		if len(devices) == 0 {
			log.Println("Skip -> no registered devices")
			continue
		}

		for _, device := range devices {

			// log.Println("Sending notification...")
			// log.Println("Device :", device.DeviceType)
			// log.Println("Token  :", device.FMCToken)

			err := services.SendPushNotification(
				device.FMCToken,
				"Reminder Habit",
				"Jangan lupa "+habit.Name,
			)

			if err != nil {
				log.Println("FAILED :", err)
			} else {
				log.Println("SUCCESS -> Notification sent")
			}
		}
	}

	// log.Println("========== END REMINDER ==========")
}
