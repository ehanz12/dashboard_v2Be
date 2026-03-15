package models

import "time"

type Transactions struct {
	BaseModel
	UserID          string
	CategoryID      string
	Amount          float64
	Type            string
	Description     string
	TransactionDate time.Time
	CreatedAt       time.Time
	Users           Users      `gorm:"foreignKey:UserID;references:ID"`
	Categories      Categories `gorm:"foreignKey:CategoryID;references:ID"`
}
