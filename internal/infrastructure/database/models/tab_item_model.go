package models

import (
	"gorm.io/gorm"
	"time"
)

type TabItem struct {
	ID        uint    `gorm:"primaryKey"`
	Tab       Tab     `gorm:"foreignKey:TabID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Product   Product `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Price     float64 `gorm:"not null"`
	Quantity  int     `gorm:"not null"`
	Obs       string  `gorm:"type:text"`
	CreatedAt time.Time
}

func (n *TabItem) BeforeSave(tx *gorm.DB) (err error) {
	return nil
}
