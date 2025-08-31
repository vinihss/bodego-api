package news

import (
	"github.com/vinihss/bodego-api/internal/domain/news"
)

type FindNewsUseCase struct {
	repo NewsRepository
}

func NewFindNewsUseCase(repo NewsRepository) *FindNewsUseCase {
	return &FindNewsUseCase{repo: repo}
}

func (uc *FindNewsUseCase) Execute(id uint) (news.News, error) {
	return uc.repo.FindByID(id)
}

func (uc *FindNewsUseCase) ExecuteAll(page, size int) ([]news.News, error) {
	offset := (page - 1) * size
	return uc.repo.FindAll(offset, size)
}