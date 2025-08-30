package models

import (
	"time"
	"gorm.io/gorm"
)

type Tab struct {
	ID          uint       `gorm:"primaryKey"`
	UserID      uint       `gorm:"not null;index"`
	OpenDate    time.Time  `gorm:"not null"`
	CloseDate   *time.Time `gorm:"default:null"`
	Description string     `gorm:"type:text"`
	Status      string     `gorm:"type:varchar(20);not null;default:'open'"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
