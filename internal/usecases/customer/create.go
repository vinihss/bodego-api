package customer

import (
	"github.com/vinihss/bodego-api/internal/domain/customer"
)

type CreateCustomer struct {
	repo customer.Repository
}
type CreateCustomerInput struct {
	Name  string
	Email string
}

func NewCreateCustomer(repo customer.Repository) *CreateCustomer {
	return &CreateCustomer{repo: repo}
}

func (uc *CreateCustomer) Execute(input CreateCustomerInput) (customer.Customer, error) {
	fav := customer.Customer{
		Name:  input.Name,
		Email: input.Email,
	}
	return uc.repo.Create(fav)
}
