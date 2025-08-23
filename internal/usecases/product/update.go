package customer

import (
	"github.com/vinihss/bodego-api/internal/domain"
	"github.com/vinihss/bodego-api/internal/domain/customer"
)

type UpdateCustomerUseCase struct {
	repo CustomerRepository
}
type UpdateCustomerInput struct {
	ID    uint
	Name  string
	Email string
}

func NewUpdateCustomerUseCase(repo CustomerRepository) *UpdateCustomerUseCase {
	return &UpdateCustomerUseCase{repo: repo}
}

func (uc *UpdateCustomerUseCase) Execute(input UpdateCustomerInput) (domain.EntityInterface, error) {
	fav := customer.Customer{
		ID:    input.ID,
		Name:  input.Name,
		Email: input.Email,
	}
	return uc.repo.Update(fav)
}
