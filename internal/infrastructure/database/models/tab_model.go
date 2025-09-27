package models

import (
	"gorm.io/gorm"
	"time"
)

type Tab struct {
	ID          uint       `gorm:"primaryKey"`
	UserID      uint       `gorm:"not null;index"`
	Customer    Customer   `gorm:"foreignKey:UserID"`
	OpenDate    time.Time  `gorm:"not null"`
	CloseDate   *time.Time `gorm:"default:null"`
	Description string     `gorm:"type:text"`
	Status      string     `gorm:"type:varchar(20);not null;default:'open'"`
	Items       []TabItem  `gorm:"foreignKey:TabID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
