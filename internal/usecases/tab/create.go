package tab

import (
	"time"
	"errors"
	"github.com/vinihss/bodego-api/internal/domain/tab"
)

type TabRepository interface {
	Create(entity tab.Tab) (tab.Tab, error)
	Delete(id uint) error
	FindByID(id uint) (tab.Tab, error)
	Update(entity tab.Tab) (tab.Tab, error)
	FindAll(offset, size int) ([]tab.Tab, error)
	FindByUserID(userID uint) ([]tab.Tab, error)
	FindOpenTabsByUserID(userID uint) ([]tab.Tab, error)
}

type OpenTabInput struct {
	UserID      uint   `json:"user_id" binding:"required"`
	Description string `json:"description,omitempty"`
}

type OpenTabUseCase struct {
	repo TabRepository
}

func NewOpenTabUseCase(repo TabRepository) *OpenTabUseCase {
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

type CreateTabUseCase struct {
	repo TabRepository
}

func NewCreateTabUseCase(repo TabRepository) *CreateTabUseCase {
	return &CreateTabUseCase{repo: repo}
}

// Deprecated: Use OpenTabUseCase instead
func (uc *CreateTabUseCase) Execute(input CreateTabInput) (tab.Tab, error) {
	// This is kept for backward compatibility but should not be used
	return tab.Tab{}, errors.New("deprecated: use OpenTabUseCase instead")
}
