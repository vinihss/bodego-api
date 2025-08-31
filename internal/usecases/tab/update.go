package tab

import (
	"errors"
	"github.com/vinihss/bodego-api/internal/domain/tab"
)

type UpdateTabUseCase struct {
	repo TabRepository
}

type UpdateTabInput struct {
	ID          uint   `json:"id" binding:"required"`
	Description string `json:"description,omitempty"`
}

func NewUpdateTabUseCase(repo TabRepository) *UpdateTabUseCase {
	return &UpdateTabUseCase{repo: repo}
}

func (uc *UpdateTabUseCase) Execute(input UpdateTabInput) (tab.Tab, error) {
	// Find the existing tab
	existingTab, err := uc.repo.FindByID(input.ID)
	if err != nil {
		return tab.Tab{}, err
	}

	// Don't allow updating closed tabs
	if existingTab.Status == tab.TabStatusClosed {
		return tab.Tab{}, errors.New("cannot update a closed tab")
	}

	// Update only the description
	existingTab.Description = input.Description

	return uc.repo.Update(existingTab)
}
