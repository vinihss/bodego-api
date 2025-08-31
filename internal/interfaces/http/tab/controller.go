package http_interfaces_tab

import (
	"github.com/vinihss/bodego-api/internal/usecases/tab"
)

type TabController struct {
	openUC   *tab.OpenTabUseCase
	closeUC  *tab.CloseTabUseCase
	deleteUC *tab.DeleteTabUseCase
	findUC   *tab.FindTabUseCase
	updateUC *tab.UpdateTabUseCase
}

func NewTabController(
	openUC *tab.OpenTabUseCase,
	closeUC *tab.CloseTabUseCase,
	deleteUC *tab.DeleteTabUseCase,
	findUC *tab.FindTabUseCase,
	updateUC *tab.UpdateTabUseCase,
) *TabController {
	return &TabController{
		openUC:   openUC,
		closeUC:  closeUC,
		deleteUC: deleteUC,
		findUC:   findUC,
		updateUC: updateUC,
	}
}

func (ctrl *TabController) OpenTab(req OpenTabRequest) (TabResponse, error) {
	result, err := ctrl.openUC.Execute(tab.OpenTabInput{
		UserID:      req.UserID,
		Description: req.Description,
	})
	if err != nil {
		return TabResponse{}, err
	}

	return TabResponse{
		ID:          result.ID,
		UserID:      result.UserID,
		OpenDate:    result.OpenDate,
		CloseDate:   result.CloseDate,
		Description: result.Description,
		Status:      string(result.Status),
	}, nil
}

func (ctrl *TabController) CloseTab(id uint) (TabResponse, error) {
	result, err := ctrl.closeUC.Execute(tab.CloseTabInput{
		ID: id,
	})
	if err != nil {
		return TabResponse{}, err
	}

	return TabResponse{
		ID:          result.ID,
		UserID:      result.UserID,
		OpenDate:    result.OpenDate,
		CloseDate:   result.CloseDate,
		Description: result.Description,
		Status:      string(result.Status),
	}, nil
}

func (ctrl *TabController) GetTab(id uint) (TabResponse, error) {
	result, err := ctrl.findUC.Execute(id)
	if err != nil {
		return TabResponse{}, err
	}

	return TabResponse{
		ID:          result.ID,
		UserID:      result.UserID,
		OpenDate:    result.OpenDate,
		CloseDate:   result.CloseDate,
		Description: result.Description,
		Status:      string(result.Status),
	}, nil
}

func (ctrl *TabController) UpdateTab(id uint, req UpdateTabRequest) (TabResponse, error) {
	result, err := ctrl.updateUC.Execute(tab.UpdateTabInput{
		ID:          id,
		Description: req.Description,
	})
	if err != nil {
		return TabResponse{}, err
	}

	return TabResponse{
		ID:          result.ID,
		UserID:      result.UserID,
		OpenDate:    result.OpenDate,
		CloseDate:   result.CloseDate,
		Description: result.Description,
		Status:      string(result.Status),
	}, nil
}

func (ctrl *TabController) DeleteTab(id uint) error {
	err := ctrl.deleteUC.Execute(id)
	if err != nil {
		return err
	}
	return nil
}

func (ctrl *TabController) GetAllTabs(page, size int) ([]TabResponse, error) {
	tabs, err := ctrl.findUC.ExecuteAll(page, size)
	if err != nil {
		return nil, err
	}

	var responses []TabResponse
	for _, t := range tabs {
		responses = append(responses, TabResponse{
			ID:          t.ID,
			UserID:      t.UserID,
			OpenDate:    t.OpenDate,
			CloseDate:   t.CloseDate,
			Description: t.Description,
			Status:      string(t.Status),
		})
	}

	return responses, nil
}

func (ctrl *TabController) GetTabsByUserID(userID uint) ([]TabResponse, error) {
	tabs, err := ctrl.findUC.ExecuteByUserID(userID)
	if err != nil {
		return nil, err
	}

	var responses []TabResponse
	for _, t := range tabs {
		responses = append(responses, TabResponse{
			ID:          t.ID,
			UserID:      t.UserID,
			OpenDate:    t.OpenDate,
			CloseDate:   t.CloseDate,
			Description: t.Description,
			Status:      string(t.Status),
		})
	}

	return responses, nil
}

func (ctrl *TabController) GetOpenTabsByUserID(userID uint) ([]TabResponse, error) {
	tabs, err := ctrl.findUC.ExecuteOpenTabsByUserID(userID)
	if err != nil {
		return nil, err
	}

	var responses []TabResponse
	for _, t := range tabs {
		responses = append(responses, TabResponse{
			ID:          t.ID,
			UserID:      t.UserID,
			OpenDate:    t.OpenDate,
			CloseDate:   t.CloseDate,
			Description: t.Description,
			Status:      string(t.Status),
		})
	}

	return responses, nil
}
