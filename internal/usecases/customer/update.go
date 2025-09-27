package customer

import (
	"github.com/vinihss/bodego-api/internal/domain/customer"
)

type UpdateCustomer struct {
	repo customer.Repository
}
type UpdateCustomerInput struct {
	ID    uint
	Name  string
	Email string
}

func NewUpdateCustomer(repo customer.Repository) *UpdateCustomer {
	return &UpdateCustomer{repo: repo}
}

func (uc *UpdateCustomer) Execute(input UpdateCustomerInput) (customer.Customer, error) {
	fav := customer.Customer{
		ID:    input.ID,
		Name:  input.Name,
		Email: input.Email,
	}
	return uc.repo.Update(fav)
}
