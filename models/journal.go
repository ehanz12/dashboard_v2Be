package models

import "time"

type Journal struct {
	BaseModel
	UserID    string    `gorm:"type:char(36);not null"`
	Mood      int       `gorm:"not null"`
	Content   string    `gorm:"type:text;not null"`
	EntryDate time.Time `gorm:"type:date;not null;uniqueIndex:idx_user_entry_date"`
	CreatedAt time.Time
	UpdatedAt time.Time

	User Users `gorm:"foreignKey:UserID;references:ID"`
}
