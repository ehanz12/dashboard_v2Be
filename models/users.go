package models

import "time"

type Users struct {
	BaseModel
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time

	Categories []Categories `gorm:"foreignKey:UserID;references:ID"`
	Habits     []Habits     `gorm:"foreignKey:UserID;references:ID"`
}