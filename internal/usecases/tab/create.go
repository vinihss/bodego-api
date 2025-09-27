package tab

import (
	"errors"
	"github.com/vinihss/bodego-api/internal/domain/tab"
	"github.com/vinihss/bodego-api/internal/infrastructure/database/repositories"
	"time"
)

type OpenTabInput struct {
	UserID      uint   `json:"user_id" binding:"required"`
	Description string `json:"description,omitempty"`
}

type OpenTabUseCase struct {
	repo tab.Repository
}

func NewOpenTab(repo tab.Repository) *OpenTabUseCase {
	return &OpenTabUseCase{repo: repo}
}

func (uc *OpenTabUseCase) Execute(input OpenTabInput) (tab.Tab, error) {
	// Check if user already has an open tab
	openTabs, err := uc.repo.FindOpenTabsByUserID(input.UserID)
	if err != nil {
		return tab.Tab{}, err
	}

	if len(openTabs) > 0 {
		return tab.Tab{}, errors.New("user already has an open tab")
	}

	newTab := tab.Tab{
		UserID:      input.UserID,
		OpenDate:    time.Now(),
		CloseDate:   nil,
		Description: input.Description,
		Status:      tab.TabStatusOpen,
	}

	return uc.repo.Create(newTab)
}

// Legacy support - keeping the old interface
type CreateTabInput struct {
	Name  string
	Email string
}

type CreateTab struct {
	repo tab.Repository
}

func NewCreateTab(repo repositories.TabRepository) *CreateTab {
	return &CreateTab{repo: repo}
}

// Deprecated: Use OpenTabUseCase instead
func (uc *CreateTab) Execute(input CreateTabInput) (tab.Tab, error) {
	// This is kept for backward compatibility but should not be used
	return tab.Tab{}, errors.New("deprecated: use OpenTabUseCase instead")
}
