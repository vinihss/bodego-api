package news

import (
	"github.com/vinihss/bodego-api/internal/domain/news"
)

type UpdateNewsUseCase struct {
	repo NewsRepository
}

type UpdateNewsInput struct {
	ID          uint
	Name        string
	Price       float64
	Description string
}

func NewUpdateNewsUseCase(repo NewsRepository) *UpdateNewsUseCase {
	return &UpdateNewsUseCase{repo: repo}
}

func (uc *UpdateNewsUseCase) Execute(input UpdateNewsInput) (news.News, error) {
	newsItem := news.News{
		ID:          input.ID,
		Name:        input.Name,
		Price:       input.Price,
		Description: input.Description,
	}
	return uc.repo.Update(newsItem)
}