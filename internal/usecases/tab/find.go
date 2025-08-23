package tab

import (
	"github.com/vinihss/bodego-api/internal/domain/tab"
)

type FindTabUseCase struct {
	repo TabRepository
}

func NewFindTabUseCase(repo TabRepository) *FindTabUseCase {
	return &FindTabUseCase{repo: repo}
}

func (uc *FindTabUseCase) Execute(id uint) (tab.Tab, error) {
	entity, err := uc.repo.FindByID(id)
	if err != nil {
		return entity, err
	}
	return entity, nil
}

func (uc *FindTabUseCase) ExecuteAll(page, size int) ([]tab.Tab, error) {
	offset := (page - 1) * size
	tabs, err := uc.repo.FindAll(offset, size)
	if err != nil {
		return nil, err
	}
	return tabs, nil
}
