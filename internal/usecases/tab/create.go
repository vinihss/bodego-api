package tab

import (
	"github.com/vinihss/bodego-api/internal/domain/tab"
)

type TabRepository interface {
	Create(entity tab.Tab) (tab.Tab, error)
	Delete(id uint) error
	FindByID(id uint) (tab.Tab, error)
	Update(entity tab.Tab) (tab.Tab, error)
	FindAll(int, size int) ([]tab.Tab, error)
}

type CreateTabInput struct {
	Name  string
	Email string
}

type CreateTabUseCase struct {
	repo TabRepository
}

func NewCreateTabUseCase(repo TabRepository) *CreateTabUseCase {
	return &CreateTabUseCase{repo: repo}
}

func (uc *CreateTabUseCase) Execute(input CreateTabInput) (tab.Tab, error) {
	fav := tab.Tab{
		Name:  input.Name,
		Email: input.Email,
	}
	return uc.repo.Create(fav)
}
