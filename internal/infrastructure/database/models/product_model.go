package models

import (
	"gorm.io/gorm"
	"strings"
)

type Product struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Email       string `gorm:"uniqueIndex"`
	Price       float64
	Description string
}

func (c *Product) BeforeSave(tx *gorm.DB) (err error) {
	c.Email = strings.TrimSpace(strings.ToLower(c.Email))
	return nil
}
