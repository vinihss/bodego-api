package repositories

import (
	"github.com/vinihss/bodego-api/internal/domain/tab"
	"github.com/vinihss/bodego-api/internal/infrastructure/database/models"
	"gorm.io/gorm"
)

type TabRepository struct {
	db *gorm.DB
}

func NewTabRepository(db *gorm.DB) *TabRepository {
	return &TabRepository{db: db}
}

func (r *TabRepository) Create(t tab.Tab) (tab.Tab, error) {
	model := models.Tab{
		UserID:      t.UserID,
		OpenDate:    t.OpenDate,
		CloseDate:   t.CloseDate,
		Description: t.Description,
		Status:      string(t.Status),
	}

	if err := r.db.Create(&model).Error; err != nil {
		return tab.Tab{}, err
	}

	return tab.Tab{
		ID:          model.ID,
		UserID:      model.UserID,
		OpenDate:    model.OpenDate,
		CloseDate:   model.CloseDate,
		Description: model.Description,
		Status:      tab.TabStatus(model.Status),
	}, nil
}

func (r *TabRepository) FindByID(id uint) (tab.Tab, error) {
	var model models.Tab

	if err := r.db.First(&model, id).Error; err != nil {
		return tab.Tab{}, err
	}

	return tab.Tab{
		ID:          model.ID,
		UserID:      model.UserID,
		OpenDate:    model.OpenDate,
		CloseDate:   model.CloseDate,
		Description: model.Description,
		Status:      tab.TabStatus(model.Status),
	}, nil
}

func (r *TabRepository) Update(t tab.Tab) (tab.Tab, error) {
	var model models.Tab

	if err := r.db.First(&model, t.ID).Error; err != nil {
		return tab.Tab{}, err
	}

	model.UserID = t.UserID
	model.CloseDate = t.CloseDate
	model.Description = t.Description
	model.Status = string(t.Status)

	if err := r.db.Save(&model).Error; err != nil {
		return tab.Tab{}, err
	}

	return tab.Tab{
		ID:          model.ID,
		UserID:      model.UserID,
		OpenDate:    model.OpenDate,
		CloseDate:   model.CloseDate,
		Description: model.Description,
		Status:      tab.TabStatus(model.Status),
	}, nil
}

func (r *TabRepository) Delete(id uint) error {
	var model models.Tab

	if err := r.db.First(&model, id).Error; err != nil {
		return err
	}

	if err := r.db.Delete(&model).Error; err != nil {
		return err
	}

	return nil
}

func (r *TabRepository) FindAll(offset, size int) ([]tab.Tab, error) {
	var models []models.Tab

	if err := r.db.Offset(offset).Limit(size).Find(&models).Error; err != nil {
		return nil, err
	}

	var tabs []tab.Tab
	for _, model := range models {
		tabs = append(tabs, tab.Tab{
			ID:          model.ID,
			UserID:      model.UserID,
			OpenDate:    model.OpenDate,
			CloseDate:   model.CloseDate,
			Description: model.Description,
			Status:      tab.TabStatus(model.Status),
		})
	}

	return tabs, nil
}

func (r *TabRepository) FindByUserID(userID uint) ([]tab.Tab, error) {
	var models []models.Tab

	if err := r.db.Where("user_id = ?", userID).Find(&models).Error; err != nil {
		return nil, err
	}

	var tabs []tab.Tab
	for _, model := range models {
		tabs = append(tabs, tab.Tab{
			ID:          model.ID,
			UserID:      model.UserID,
			OpenDate:    model.OpenDate,
			CloseDate:   model.CloseDate,
			Description: model.Description,
			Status:      tab.TabStatus(model.Status),
		})
	}

	return tabs, nil
}

func (r *TabRepository) FindOpenTabsByUserID(userID uint) ([]tab.Tab, error) {
	var models []models.Tab

	if err := r.db.Where("user_id = ? AND status = ?", userID, "open").Find(&models).Error; err != nil {
		return nil, err
	}

	var tabs []tab.Tab
	for _, model := range models {
		tabs = append(tabs, tab.Tab{
			ID:          model.ID,
			UserID:      model.UserID,
			OpenDate:    model.OpenDate,
			CloseDate:   model.CloseDate,
			Description: model.Description,
			Status:      tab.TabStatus(model.Status),
		})
	}

	return tabs, nil
}