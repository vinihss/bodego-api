package http_interfaces_news

import (
	"github.com/vinihss/bodego-api/internal/usecases/news"
)

type NewsController struct {
	createUC *news.CreateNewsUseCase
	deleteUC *news.DeleteNewsUseCase
	findUC   *news.FindNewsUseCase
	updateUC *news.UpdateNewsUseCase
}

func NewNewsController(
	createUC *news.CreateNewsUseCase,
	deleteUC *news.DeleteNewsUseCase,
	findUC *news.FindNewsUseCase,
	updateUC *news.UpdateNewsUseCase,
) *NewsController {
	return &NewsController{createUC: createUC, deleteUC: deleteUC, findUC: findUC, updateUC: updateUC}
}

func (ctrl *NewsController) CreateNews(req CreateNewsRequest) (NewsResponse, error) {
	newsItem, err := ctrl.createUC.Execute(news.CreateNewsInput{
		Name:        req.Name,
		Price:       req.Price,
		Description: req.Description,
	})
	if err != nil {
		return NewsResponse{}, err
	}

	return NewsResponse{
		ID:          newsItem.ID,
		Name:        newsItem.Name,
		Price:       newsItem.Price,
		Description: newsItem.Description,
	}, nil
}

func (ctrl *NewsController) GetNews(id uint) (NewsResponse, error) {
	newsItem, err := ctrl.findUC.Execute(id)
	if err != nil {
		return NewsResponse{}, err
	}

	return NewsResponse{
		ID:          newsItem.ID,
		Name:        newsItem.Name,
		Price:       newsItem.Price,
		Description: newsItem.Description,
	}, nil
}

func (ctrl *NewsController) DeleteNews(id uint) error {
	err := ctrl.deleteUC.Execute(id)
	if err != nil {
		return err
	}

	return nil
}

func (ctrl *NewsController) UpdateNews(id uint, req UpdateNewsRequest) (NewsResponse, error) {
	newsItem, err := ctrl.updateUC.Execute(news.UpdateNewsInput{
		ID:          id,
		Name:        req.Name,
		Price:       req.Price,
		Description: req.Description,
	})
	if err != nil {
		return NewsResponse{}, err
	}

	return NewsResponse{
		ID:          newsItem.ID,
		Name:        newsItem.Name,
		Price:       newsItem.Price,
		Description: newsItem.Description,
	}, nil
}

func (ctrl *NewsController) GetAllNews(page, size int) ([]NewsResponse, error) {
	newsItems, err := ctrl.findUC.ExecuteAll(page, size)
	if err != nil {
		return nil, err
	}

	var responses []NewsResponse
	for _, n := range newsItems {
		responses = append(responses, NewsResponse{
			ID:          n.ID,
			Name:        n.Name,
			Price:       n.Price,
			Description: n.Description,
		})
	}

	return responses, nil
}