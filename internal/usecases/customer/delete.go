package customer

import (
	"github.com/vinihss/bodego-api/internal/domain/customer"
)

type DeleteCustomer struct {
	repo customer.Repository
}

type DeleteCustomerInput struct {
	ID    uint
	Name  string
	Email string
}

func NewDeleteCustomer(repo customer.Repository) *DeleteCustomer {
	return &DeleteCustomer{repo: repo}
}

func (uc *DeleteCustomer) Execute(input DeleteCustomerInput) error {
	entity, err := uc.repo.FindByID(input.ID)
	if err != nil {
		return err
	}

	if uc.repo.Delete(entity.ID) != nil {
		return err
	}
	return nil
}
