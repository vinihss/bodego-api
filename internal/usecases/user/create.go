package user

import (
	"github.com/vinihss/bodego-api/internal/domain/user"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserInput struct {
	Name     string
	Email    string
	Password string
	Role     user.Role
}

type CreateUserOutput struct {
	ID    uint
	Email string
	Role  user.Role
}

type CreateUser struct {
	Repo user.Repository
}

func NewCreateUser(repo user.Repository) *CreateUser {
	return &CreateUser{Repo: repo}
}

func (uc *CreateUser) Execute(input CreateUserInput) (*CreateUserOutput, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user := &user.User{
		Email:    input.Email,
		Password: string(hash),
		Role:     input.Role,
	}
	err = uc.Repo.Create(user)
	if err != nil {
		return nil, err
	}
	return &CreateUserOutput{
		ID:    user.ID,
		Email: user.Email,
		Role:  user.Role,
	}, nil
}
