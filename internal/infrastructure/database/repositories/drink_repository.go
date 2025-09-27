package repositories

import (
	"gorm.io/gorm"
	"github.com/vinihss/bodego-api/internal/domain/drink"
	"github.com/vinihss/bodego-api/internal/infrastructure/database/models"
)

type DrinkRepository struct {
	db *gorm.DB
}

func NewDrinkRepository(db *gorm.DB) *DrinkRepository {
	return &DrinkRepository{db: db}
}

func (r *DrinkRepository) Create(drink drink.Drink) (drink.Drink, error) {
	model := models.Drink{Name: drink.Name, Price: drink.Price}
	if err := r.db.Create(&model).Error; err != nil {
		return drink.Drink{}, err
	}
	drink.ID = model.ID
	return drink, nil
}

func (r *DrinkRepository) FindByID(id uint) (*drink.Drink, error) {
	var model models.Drink
	if err := r.db.First(&model, id).Error; err != nil {
		return nil, err
	}
	return &drink.Drink{ID: model.ID, Name: model.Name, Price: model.Price}, nil
}

func (r *DrinkRepository) FindAll() ([]drink.Drink, error) {
	var modelsList []models.Drink
	if err := r.db.Find(&modelsList).Error; err != nil {
		return nil, err
	}
	drinks := make([]drink.Drink, len(modelsList))
	for i, m := range modelsList {
		drinks[i] = drink.Drink{ID: m.ID, Name: m.Name, Price: m.Price}
	}
	return drinks, nil
}

func (r *DrinkRepository) Update(drink drink.Drink) (*drink.Drink, error) {
	var model models.Drink
	if err := r.db.First(&model, drink.ID).Error; err != nil {
		return nil, err
	}
	model.Name = drink.Name
	model.Price = drink.Price
	if err := r.db.Save(&model).Error; err != nil {
		return nil, err
	}
	return &drink, nil
}

func (r *DrinkRepository) Delete(id uint) error {
	return r.db.Delete(&models.Drink{}, id).Error
}
