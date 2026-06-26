package models

import (
	"time"

	"gorm.io/datatypes"
)

type Habits struct {
	BaseModel
	UserID          string
	Name            string
	Frequency       string
	Days            datatypes.JSON `gorm:"type:json;default:null"`
	ReminderTime    *time.Time
	ReminderEnabled bool
	CreatedAt       time.Time

	Users Users `gorm:"foreignKey:UserID;references:ID"`
}
