package user

import (
	"errors"
	"reflect"
	"testing"
	uc "github.com/vinihss/bodego-api/internal/usecases/user"
	"github.com/vinihss/bodego-api/internal/domain/user"
)

type mockRepo struct {
	created user.User
	createErr error
}

func (m *mockRepo) Create(entity user.User) (user.User, error) {
	m.created = entity
	return entity, m.createErr
}
func (m *mockRepo) Delete(id uint) error { return nil }
func (m *mockRepo) FindByID(id uint) (user.User, error) { return user.User{}, nil }
func (m *mockRepo) Update(entity user.User) (user.User, error) { return entity, nil }
func (m *mockRepo) FindAll(int, int) ([]user.User, error) { return nil, nil }

func TestCreateUserUseCase_Execute(t *testing.T) {
	repo := &mockRepo{}
	usecase := uc.NewCreateUserUseCase(repo)
	input := uc.CreateUserInput{Name: "User", Email: "user@email.com"}
	result, err := usecase.Execute(input)
	if err != nil {
		t.Fatalf("esperado erro nulo, recebeu: %v", err)
	}
	if !reflect.DeepEqual(result, repo.created) {
		t.Errorf("esperado: %+v, recebeu: %+v", repo.created, result)
	}
}

func TestCreateUserUseCase_Execute_Error(t *testing.T) {
	repo := &mockRepo{createErr: errors.New("erro ao criar")}
	usecase := uc.NewCreateUserUseCase(repo)
	input := uc.CreateUserInput{Name: "User", Email: "user@email.com"}
	_, err := usecase.Execute(input)
	if err == nil {
		t.Error("esperado erro, recebeu nil")
	}
}
