package models

import "time"

type Task struct {
	BaseModel
	UserID      string     `gorm:"type:char(36);not null"`
	Title       string     `gorm:"type:varchar(255);not null"`
	Description string     `gorm:"type:text"`
	Quadrant    int        `gorm:"not null;default:1"`
	IsCompleted bool       `gorm:"not null;default:false"`
	DueDate     *time.Time `gorm:"default:null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time

	User Users `gorm:"foreignKey:UserID;references:ID"`
}
