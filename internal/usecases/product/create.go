package product

import (
	"github.com/vinihss/bodego-api/internal/domain/product"
)

type ProductRepository interface {
	Create(entity product.Product) (product.Product, error)
	Delete(id uint) error
	FindByID(id uint) (product.Product, error)
	Update(entity product.Product) (product.Product, error)
	FindAll(int, size int) ([]product.Product, error)
}

type CreateProductInput struct {
	Name        string
	Price       float64
	Description string
}

type CreateProductUseCase struct {
	repo ProductRepository
}

func NewCreateProductUseCase(repo ProductRepository) *CreateProductUseCase {
	return &CreateProductUseCase{repo: repo}
}

func (uc *CreateProductUseCase) Execute(input CreateProductInput) (product.Product, error) {
	fav := product.Product{
		Name:        input.Name,
		Price:       input.Price,
		Description: input.Description,
	}
	return uc.repo.Create(fav)
}
