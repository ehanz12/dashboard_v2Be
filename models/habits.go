package models

import "time"

type Habits struct {
	BaseModel
	UserID       string
	Name         string
	Description  string
	TargetPerDay int
	ReminderTime string
	CreatedAt    time.Time

	Users Users `gorm:"foreignKey:UserID;references:ID"`
}