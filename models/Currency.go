package models

import "gorm.io/gorm"

type Currency struct {
	gorm.Model

	Code  string  `gorm:"not null" json:"code"`
	Value float64 `gorm:"not null" json:"value"`
}
