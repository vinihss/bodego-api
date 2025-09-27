package user

import (
	"github.com/vinihss/bodego-api/internal/domain/user"
)

type Repository interface {
	user.Repository
}
type Entity struct {
	user.User
}
type FindCustomer struct {
	repo user.Repository
}

func NewFindCustomer(repo Repository) *FindCustomer {
	return &FindCustomer{repo: repo}
}

func (uc *FindCustomer) Execute(id uint) (user.User, error) {
	entity, err := uc.repo.FindByID(id)
	if err != nil {
		return entity, err
	}
	return entity, nil
}

func (uc *FindCustomer) ExecuteAll(page, size int) ([]user.User, error) {
	offset := (page - 1) * size
	customers, err := uc.repo.List(offset, size)
	if err != nil {
		return nil, err
	}
	return customers, nil
}
