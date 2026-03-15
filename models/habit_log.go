package models

import "time"

type HabitLogs struct {
	BaseModel
	HabitID   string
	LogDate   time.Time
	Completed *bool
	CreatedAt time.Time
	Habits    Habits `gorm:"foreignKey:HabitID;references:ID"`
}
