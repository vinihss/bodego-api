package product

import (
	"github.com/vinihss/bodego-api/internal/domain/product"
)

type FindProductUseCase struct {
	repo ProductRepository
}

func NewFindProductUseCase(repo ProductRepository) *FindProductUseCase {
	return &FindProductUseCase{repo: repo}
}

func (uc *FindProductUseCase) Execute(id uint) (product.Product, error) {
	entity, err := uc.repo.FindByID(id)
	if err != nil {
		return entity, err
	}
	return entity, nil
}

func (uc *FindProductUseCase) ExecuteAll(page, size int) ([]product.Product, error) {
	offset := (page - 1) * size
	products, err := uc.repo.FindAll(offset, size)
	if err != nil {
		return nil, err
	}
	return products, nil
}
