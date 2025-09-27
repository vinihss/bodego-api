package repositories

import (
	"github.com/vinihss/bodego-api/internal/domain/user"
	"github.com/vinihss/bodego-api/internal/infrastructure/database/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(t user.User) (user.User, error) {
	model := models.User{}

	if err := r.db.Create(&model).Error; err != nil {
		return user.User{}, err
	}

	return user.User{
		ID:       model.ID,
		Name:     model.Name,
		Email:    model.Email,
		Password: model.Password,
		Role:     user.Role(model.Role),
	}, nil
}

func (r *UserRepository) FindByID(id uint) (user.User, error) {
	var model models.User

	if err := r.db.First(&model, id).Error; err != nil {
		return user.User{}, err
	}

	return user.User{
		ID:       model.ID,
		Name:     model.Name,
		Email:    model.Email,
		Password: model.Password,
		Role:     user.Role(model.Role),
	}, nil
}

func (r *UserRepository) Update(t user.User) (user.User, error) {
	var model models.User

	if err := r.db.First(&model, t.ID).Error; err != nil {
		return user.User{}, err
	}

	model.ID = t.ID
	model.Name = t.Name
	model.Email = t.Email
	model.Password = t.Password

	if err := r.db.Save(&model).Error; err != nil {
		return user.User{}, err
	}

	return user.User{
		ID:       model.ID,
		Name:     model.Name,
		Email:    model.Email,
		Password: model.Password,
		Role:     user.Role(model.Role),
	}, nil
}

func (r *UserRepository) Delete(id uint) error {
	var model models.User

	if err := r.db.First(&model, id).Error; err != nil {
		return err
	}

	if err := r.db.Delete(&model).Error; err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) FindAll(offset, size int) ([]user.User, error) {
	var models []models.User

	if err := r.db.Offset(offset).Limit(size).Find(&models).Error; err != nil {
		return nil, err
	}

	var users []user.User
	for _, model := range models {
		users = append(users, user.User{
			ID:       model.ID,
			Name:     model.Name,
			Email:    model.Email,
			Password: model.Password,
			Role:     user.Role(model.Role),
		})
	}

	return users, nil
}

func (r *UserRepository) FindByUserID(userID uint) ([]user.User, error) {
	var models []models.User

	if err := r.db.Where("user_id = ?", userID).Find(&models).Error; err != nil {
		return nil, err
	}

	var users []user.User
	for _, model := range models {
		users = append(users, user.User{
			ID:       model.ID,
			Name:     model.Name,
			Email:    model.Email,
			Password: model.Password,
			Role:     user.Role(model.Role),
		})
	}

	return users, nil

}

func (r *UserRepository) FindByEmail(email string) ([]user.User, error) {
	var models []models.User

	if err := r.db.Where("email = ?", email).Find(&models).Error; err != nil {
		return nil, err
	}

	var users []user.User
	for _, model := range models {
		users = append(users, user.User{
			ID:       model.ID,
			Name:     model.Name,
			Email:    model.Email,
			Password: model.Password,
			Role:     user.Role(model.Role),
		})
	}

	return users, nil

}
