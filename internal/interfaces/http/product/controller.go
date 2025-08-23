package http_interfaces_product

import (
	"github.com/vinihss/bodego-api/internal/usecases/product"
)

type ProductController struct {
	createUC *product.CreateProductUseCase
	deleteUC *product.DeleteProductUseCase
	findUC   *product.FindProductUseCase
	updateUC *product.UpdateProductUseCase
}

func NewProductController(
	createUC *product.CreateProductUseCase,
	deleteUC *product.DeleteProductUseCase,
	findUC *product.FindProductUseCase,
	updateUC *product.UpdateProductUseCase,
) *ProductController {
	return &ProductController{createUC: createUC, deleteUC: deleteUC, findUC: findUC, updateUC: updateUC}
}

func (ctrl *ProductController) CreateProduct(req CreateProductRequest) (ProductResponse, error) {
	fav, err := ctrl.createUC.Execute(product.CreateProductInput{
		Name:        req.Name,
		Price:       req.Price,
		Description: req.Description,
	})
	if err != nil {
		return ProductResponse{}, err
	}

	return ProductResponse{
		ID:          fav.ID,
		Name:        fav.Name,
		Price:       fav.Price,
		Description: fav.Description,
	}, nil
}

func (ctrl *ProductController) GetProduct(id uint) (ProductResponse, error) {
	fav, err := ctrl.findUC.Execute(id)
	if err != nil {
		return ProductResponse{}, err
	}

	return ProductResponse{
		ID:          fav.ID,
		Name:        fav.Name,
		Price:       fav.Price,
		Description: fav.Description,
	}, nil
}

func (ctrl *ProductController) DeleteProduct(id uint) error {
	err := ctrl.deleteUC.Execute(id)
	if err != nil {
		return err
	}

	return nil
}

func (ctrl *ProductController) UpdateProduct(id uint, req UpdateProductRequest) (ProductResponse, error) {
	input, err := ctrl.updateUC.Execute(product.UpdateProductInput{
		ID:          id,
		Name:        req.Name,
		Price:       req.Price,
		Description: req.Description,
	})
	if err != nil {
		return ProductResponse{}, err
	}

	return ProductResponse{
		ID:          input.ID,
		Name:        input.Name,
		Price:       input.Price,
		Description: input.Description,
	}, nil
}

func (ctrl *ProductController) GetAllProducts(page, size int) ([]ProductResponse, error) {
	products, err := ctrl.findUC.ExecuteAll(page, size)
	if err != nil {
		return nil, err
	}

	var responses []ProductResponse
	for _, c := range products {
		responses = append(responses, ProductResponse{
			ID:          c.ID,
			Name:        c.Name,
			Price:       c.Price,
			Description: c.Description,
		})
	}

	return responses, nil
}
