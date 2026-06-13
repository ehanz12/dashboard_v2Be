package models

import "time"

type Timeblock struct {
	BaseModel
	UserID       string     `gorm:"type:char(36);not null"`
	ActivityName string     `gorm:"type:varchar(255);not null"`
	StartTime    string     `gorm:"type:varchar(5);not null"`
	EndTime      string     `gorm:"type:varchar(5);not null"`
	ColorCode    string     `gorm:"type:varchar(7);not null;default:'#4F46E5'"`
	DayOfWeek    int        `gorm:"not null"`
	Date         *time.Time `gorm:"type:date;default:null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time

	User Users `gorm:"foreignKey:UserID;references:ID"`
}
