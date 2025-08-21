package models

import (
	"gorm.io/gorm"
	"strings"
)

type Tab struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Email string `gorm:"uniqueIndex"`
}

func (c *Tab) BeforeSave(tx *gorm.DB) (err error) {
	c.Email = strings.TrimSpace(strings.ToLower(c.Email))
	return nil
}
