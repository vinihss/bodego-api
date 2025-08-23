package http_interfaces_tab

import (
	"github.com/vinihss/bodego-api/internal/usecases/tab"
)

type TabController struct {
	createUC *tab.CreateTabUseCase
	deleteUC *tab.DeleteTabUseCase
	findUC   *tab.FindTabUseCase
	updateUC *tab.UpdateTabUseCase
}

func NewTabController(
	createUC *tab.CreateTabUseCase,
	deleteUC *tab.DeleteTabUseCase,
	findUC *tab.FindTabUseCase,
	updateUC *tab.UpdateTabUseCase,
) *TabController {
	return &TabController{createUC: createUC, deleteUC: deleteUC, findUC: findUC, updateUC: updateUC}
}

func (ctrl *TabController) CreateTab(req CreateTabRequest) (TabResponse, error) {
	fav, err := ctrl.createUC.Execute(tab.CreateTabInput{
		Name:  req.Name,
		Email: req.Email,
	})
	if err != nil {
		return TabResponse{}, err
	}

	return TabResponse{
		ID:    fav.ID,
		Name:  fav.Name,
		Email: fav.Email,
	}, nil
}

func (ctrl *TabController) GetTab(id uint) (TabResponse, error) {
	fav, err := ctrl.findUC.Execute(id)
	if err != nil {
		return TabResponse{}, err
	}

	return TabResponse{
		ID:    fav.ID,
		Name:  fav.Name,
		Email: fav.Email,
	}, nil
}

func (ctrl *TabController) DeleteTab(id uint) error {
	err := ctrl.deleteUC.Execute(id)
	if err != nil {
		return err
	}

	return nil
}

func (ctrl *TabController) UpdateTab(id uint, req UpdateTabRequest) (TabResponse, error) {
	input, err := ctrl.updateUC.Execute(tab.UpdateTabInput{
		ID:    id,
		Name:  req.Name,
		Email: req.Email,
	})
	if err != nil {
		return TabResponse{}, err
	}

	return TabResponse{
		ID:    input.ID,
		Name:  input.Name,
		Email: input.Email,
	}, nil
}

func (ctrl *TabController) GetAllTabs(page, size int) ([]TabResponse, error) {
	tabs, err := ctrl.findUC.ExecuteAll(page, size)
	if err != nil {
		return nil, err
	}

	var responses []TabResponse
	for _, c := range tabs {
		responses = append(responses, TabResponse{
			ID:    c.ID,
			Name:  c.Name,
			Email: c.Email,
		})
	}

	return responses, nil
}
