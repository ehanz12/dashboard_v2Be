package models

import "time"

type UserDevice struct {
	BaseModel
	UserID     string `gorm:"type:char(36);not null"`
	FMCToken   string `gorm:"type:varchar(255);not null"`
	DeviceType string `gorm:"type:enum('ios', 'android');not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Users      Users `gorm:"foreignKey:UserID;references:ID"`
}
