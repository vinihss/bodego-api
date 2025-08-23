package customer

import (
	"github.com/vinihss/bodego-api/internal/domain"
)

type FindCustomerUseCase struct {
	repo CustomerRepository
}

func NewFindCustomerUseCase(repo CustomerRepository) *FindCustomerUseCase {
	return &FindCustomerUseCase{repo: repo}
}

func (uc *FindCustomerUseCase) Execute(id uint) (domain.EntityInterface, error) {
	entity, err := uc.repo.FindByID(id)
	if err != nil {
		return entity, err
	}
	return entity, nil
}

func (uc *FindCustomerUseCase) ExecuteAll(page, size int) ([]domain.EntityInterface, error) {
	offset := (page - 1) * size
	customers, err := uc.repo.FindAll(offset, size)
	if err != nil {
		return nil, err
	}
	return customers, nil
}
