package models

import "time"

type Users struct {
	BaseModel
	Name                 string
	Email                string
	Password             string
	NomorHP              string
	Bio                  *string
	EmailVerified        bool
	VerificationCode     string
	VerificationExpireAt *time.Time
	CreatedAt            time.Time
	UpdatedAt            time.Time

	Categories []Categories `gorm:"foreignKey:UserID;references:ID"`
	Habits     []Habits     `gorm:"foreignKey:UserID;references:ID"`
	Journals   []Journal    `gorm:"foreignKey:UserID;references:ID"`
	Tasks      []Task       `gorm:"foreignKey:UserID;references:ID"`
	Timeblocks []Timeblock  `gorm:"foreignKey:UserID;references:ID"`
}
