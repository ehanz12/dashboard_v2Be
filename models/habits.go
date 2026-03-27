package models

import "time"

type Habits struct {
	BaseModel
	UserID    string
	Name      string
	Frequency string
	CreatedAt time.Time

	Users Users `gorm:"foreignKey:UserID;references:ID"`
}