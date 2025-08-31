package models

import (
	"gorm.io/gorm"
)

type Drink struct {
	ID    uint    `gorm:"primaryKey"`
	Name  string  `gorm:"uniqueIndex"`
	Price float64
}

func (d *Drink) TableName() string {
	return "drinks"
}

func MigrateDrinks(db *gorm.DB) error {
	return db.AutoMigrate(&Drink{})
}

func SeedDrinks(db *gorm.DB) error {
	drinks := []Drink{
		{Name: "Cerveja Heineken Latão", Price: 10.00},
		{Name: "Dose cachaça", Price: 4.00},
	}
	for _, drink := range drinks {
		var existing Drink
		if err := db.Where("name = ?", drink.Name).First(&existing).Error; err == gorm.ErrRecordNotFound {
			if err := db.Create(&drink).Error; err != nil {
				return err
			}
		} else if err != nil && err != gorm.ErrRecordNotFound {
			return err
		}
	}
	return nil
}
