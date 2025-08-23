package customer

import (
	"github.com/vinihss/bodego-api/internal/domain"
	"github.com/vinihss/bodego-api/internal/domain/customer"
)

type CustomerRepository interface {
	Create(entity domain.EntityInterface) (domain.EntityInterface, error)
	Delete(id uint) error
	FindByID(id uint) (domain.EntityInterface, error)
	Update(entity domain.EntityInterface) (domain.EntityInterface, error)
	FindAll(int, size int) ([]domain.EntityInterface, error)
}

type CreateCustomerInput struct {
	Name  string
	Email string
}

type CreateCustomerUseCase struct {
	repo CustomerRepository
}

func NewCreateCustomerUseCase(repo CustomerRepository) *CreateCustomerUseCase {
	return &CreateCustomerUseCase{repo: repo}
}

func (uc *CreateCustomerUseCase) Execute(input CreateCustomerInput) (domain.EntityInterface, error) {
	fav := customer.Customer{
		Name:  input.Name,
		Email: input.Email,
	}
	return uc.repo.Create(fav)
}
