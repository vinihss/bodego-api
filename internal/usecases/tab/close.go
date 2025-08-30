package tab

import (
	"errors"
	"github.com/vinihss/bodego-api/internal/domain/tab"
	"time"
)

type CloseTabInput struct {
	ID uint `json:"id" binding:"required"`
}

type CloseTabUseCase struct {
	repo TabRepository
}

func NewCloseTabUseCase(repo TabRepository) *CloseTabUseCase {
	return &CloseTabUseCase{repo: repo}
}

func (uc *CloseTabUseCase) Execute(input CloseTabInput) (tab.Tab, error) {
	// Find the tab
	existingTab, err := uc.repo.FindByID(input.ID)
	if err != nil {
		return tab.Tab{}, err
	}

	// Check if tab is already closed
	if existingTab.Status == tab.TabStatusClosed {
		return tab.Tab{}, errors.New("tab is already closed")
	}

	// Close the tab
	now := time.Now()
	existingTab.CloseDate = &now
	existingTab.Status = tab.TabStatusClosed

	return uc.repo.Update(existingTab)
}
