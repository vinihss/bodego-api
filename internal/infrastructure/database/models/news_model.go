package models

import (
	"gorm.io/gorm"
)

type News struct {
	ID          uint    `gorm:"primaryKey"`
	Name        string  `gorm:"not null"`
	Price       float64 `gorm:"not null"`
	Description string
}

func (n *News) BeforeSave(tx *gorm.DB) (err error) {
	return nil
}