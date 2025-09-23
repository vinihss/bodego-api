package user

import (
	"errors"
	"reflect"
	"testing"
	uc "github.com/vinihss/bodego-api/internal/usecases/user"
	"github.com/vinihss/bodego-api/internal/domain/user"
)

type mockRepo struct {
	findByID    user.User
	findByIDErr error
	findAll     []user.User
	findAllErr  error
}

func (m *mockRepo) Create(entity user.User) (user.User, error) { return entity, nil }
func (m *mockRepo) Delete(id uint) error { return nil }
func (m *mockRepo) FindByID(id uint) (user.User, error) { return m.findByID, m.findByIDErr }
func (m *mockRepo) Update(entity user.User) (user.User, error) { return entity, nil }
func (m *mockRepo) FindAll(offset, size int) ([]user.User, error) { return m.findAll, m.findAllErr }

func TestFindUserUseCase_Execute(t *testing.T) {
	repo := &mockRepo{findByID: user.User{ID: 1, Name: "User"}}
	usecase := uc.NewFindUserUseCase(repo)
	result, err := usecase.Execute(1)
	if err != nil {
		t.Fatalf("esperado erro nulo, recebeu: %v", err)
	}
	if !reflect.DeepEqual(result, repo.findByID) {
		t.Errorf("esperado: %+v, recebeu: %+v", repo.findByID, result)
	}
}

func TestFindUserUseCase_Execute_Error(t *testing.T) {
	repo := &mockRepo{findByIDErr: errors.New("erro ao buscar")}
	usecase := uc.NewFindUserUseCase(repo)
	_, err := usecase.Execute(2)
	if err == nil {
		t.Error("esperado erro, recebeu nil")
	}
}

func TestFindUserUseCase_ExecuteAll(t *testing.T) {
	repo := &mockRepo{findAll: []user.User{{ID: 1}, {ID: 2}}}
	usecase := uc.NewFindUserUseCase(repo)
	result, err := usecase.ExecuteAll(1, 2)
	if err != nil {
		t.Fatalf("esperado erro nulo, recebeu: %v", err)
	}
	if !reflect.DeepEqual(result, repo.findAll) {
		t.Errorf("esperado: %+v, recebeu: %+v", repo.findAll, result)
	}
}

func TestFindUserUseCase_ExecuteAll_Error(t *testing.T) {
	repo := &mockRepo{findAllErr: errors.New("erro ao buscar todos")}
	usecase := uc.NewFindUserUseCase(repo)
	_, err := usecase.ExecuteAll(1, 2)
	if err == nil {
		t.Error("esperado erro, recebeu nil")
	}
}
