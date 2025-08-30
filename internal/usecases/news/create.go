package news

import (
	"github.com/vinihss/bodego-api/internal/domain/news"
)

type NewsRepository interface {
	Create(entity news.News) (news.News, error)
	Delete(id uint) error
	FindByID(id uint) (news.News, error)
	Update(entity news.News) (news.News, error)
	FindAll(offset, size int) ([]news.News, error)
}

type CreateNewsInput struct {
	Name        string
	Price       float64
	Description string
}

type CreateNewsUseCase struct {
	repo NewsRepository
}

func NewCreateNewsUseCase(repo NewsRepository) *CreateNewsUseCase {
	return &CreateNewsUseCase{repo: repo}
}

func (uc *CreateNewsUseCase) Execute(input CreateNewsInput) (news.News, error) {
	newsItem := news.News{
		Name:        input.Name,
		Price:       input.Price,
		Description: input.Description,
	}
	return uc.repo.Create(newsItem)
}