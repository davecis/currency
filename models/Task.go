package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model

	Status bool    `gorm:"not null" json:"code"`
	Time   float64 `gorm:"not null" json:"time"`
}
