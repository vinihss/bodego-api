package customer

import (
	"github.com/vinihss/bodego-api/internal/domain/customer"
)

type FindCustomer struct {
	repo customer.Repository
}

func NewFindCustomer(repo customer.Repository) *FindCustomer {
	return &FindCustomer{repo: repo}
}

func (uc *FindCustomer) Execute(id uint) (customer.Customer, error) {
	entity, err := uc.repo.FindByID(id)
	if err != nil {
		return entity, err
	}
	return entity, nil
}

func (uc *FindCustomer) ExecuteAll(page, size int) ([]customer.Customer, error) {
	offset := (page - 1) * size
	customers, err := uc.repo.FindAll(offset, size)
	if err != nil {
		return nil, err
	}
	return customers, nil
}
