
package user

import (
	"context"
	"golang.org/x/crypto/bcrypt"
)


type CreateUserInput struct {
	Email    string
	Password string
	Role     Role
}

type CreateUserOutput struct {
	ID    uint
	Email string
	Role  Role
}

type CreateUserUseCase struct {
	Repo Repository
}

func NewCreateUserUseCase(repo Repository) *CreateUserUseCase {
	return &CreateUserUseCase{Repo: repo}
}

func (uc *CreateUserUseCase) Execute(ctx context.Context, input CreateUserInput) (*CreateUserOutput, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user := &User{
		Email:    input.Email,
		Password: string(hash),
		Role:     input.Role,
	}
	err = uc.Repo.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	return &CreateUserOutput{
		ID:    user.ID,
		Email: user.Email,
		Role:  user.Role,
	}, nil
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

func (uc *CreateCustomerUseCase) Execute(input CreateCustomerInput) (customer.Customer, error) {
	fav := customer.Customer{
		Name:  input.Name,
		Email: input.Email,
	}
	return uc.repo.Create(fav)
}
