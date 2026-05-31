package models

import "time"

type Categories struct {
	BaseModel
	UserID    string
	Name      string
	CreatedAt time.Time

	Users Users `gorm:"foreignKey:UserID;references:ID"`
}
