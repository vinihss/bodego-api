package tab

import (
	"github.com/vinihss/bodego-api/internal/domain/tab"
)

type UpdateTabUseCase struct {
	repo TabRepository
}
type UpdateTabInput struct {
	ID    uint
	Name  string
	Email string
}

func NewUpdateTabUseCase(repo TabRepository) *UpdateTabUseCase {
	return &UpdateTabUseCase{repo: repo}
}

func (uc *UpdateTabUseCase) Execute(input UpdateTabInput) (tab.Tab, error) {
	fav := tab.Tab{
		ID:    input.ID,
		Name:  input.Name,
		Email: input.Email,
	}
	return uc.repo.Update(fav)
}
