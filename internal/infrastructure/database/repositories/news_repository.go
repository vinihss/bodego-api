package repositories

import (
	"github.com/vinihss/bodego-api/internal/domain/news"
	"github.com/vinihss/bodego-api/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type NewsRepository struct {
	db *gorm.DB
}

func NewNewsRepository(db *gorm.DB) *NewsRepository {
	return &NewsRepository{db: db}
}

func (r *NewsRepository) Create(n news.News) (news.News, error) {
	model := models.News{
		Name:        n.Name,
		Price:       n.Price,
		Description: n.Description,
	}

	if err := r.db.Create(&model).Error; err != nil {
		return news.News{}, err
	}

	return news.News{
		ID:          model.ID,
		Name:        model.Name,
		Price:       model.Price,
		Description: model.Description,
	}, nil
}

func (r *NewsRepository) FindByID(id uint) (news.News, error) {
	var model models.News

	if err := r.db.First(&model, id).Error; err != nil {
		return news.News{}, err
	}

	return news.News{
		ID:          model.ID,
		Name:        model.Name,
		Price:       model.Price,
		Description: model.Description,
	}, nil
}

func (r *NewsRepository) Delete(id uint) error {
	var model models.News

	if err := r.db.First(&model, id).Error; err != nil {
		return err
	}

	if err := r.db.Delete(&model).Error; err != nil {
		return err
	}

	return nil
}

func (r *NewsRepository) Update(new news.News) (news.News, error) {
	var model models.News

	if err := r.db.First(&model, new.ID).Error; err != nil {
		return news.News{}, err
	}

	model.Name = new.Name
	model.Price = new.Price
	model.Description = new.Description

	if err := r.db.Save(&model).Error; err != nil {
		return news.News{}, err
	}

	return news.News{
		ID:          model.ID,
		Name:        model.Name,
		Price:       model.Price,
		Description: model.Description,
	}, nil
}

func (r *NewsRepository) FindAll(offset, size int) ([]news.News, error) {
	var models []models.News

	if err := r.db.Offset(offset).Limit(size).Find(&models).Error; err != nil {
		return nil, err
	}

	var newsItems []news.News
	for _, model := range models {
		newsItems = append(newsItems, news.News{
			ID:          model.ID,
			Name:        model.Name,
			Price:       model.Price,
			Description: model.Description,
		})
	}

	return newsItems, nil
}