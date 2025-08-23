package product

import (
	"github.com/vinihss/bodego-api/internal/domain/product"
)

type UpdateProductUseCase struct {
	repo ProductRepository
}
type UpdateProductInput struct {
	ID          uint
	Name        string
	Price       float64
	Description string
}

func NewUpdateProductUseCase(repo ProductRepository) *UpdateProductUseCase {
	return &UpdateProductUseCase{repo: repo}
}

func (uc *UpdateProductUseCase) Execute(input UpdateProductInput) (product.Product, error) {
	fav := product.Product{
		ID:          input.ID,
		Name:        input.Name,
		Price:       input.Price,
		Description: input.Description,
	}
	return uc.repo.Update(fav)
}
