package models

import "time"

type Habits struct {
	BaseModel
	UserID    string
	Name      string
	Frequency string
	Date      *time.Time `gorm:"type:date;default:null"`
	CreatedAt time.Time

	Users Users `gorm:"foreignKey:UserID;references:ID"`
}
