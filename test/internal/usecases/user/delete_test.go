package user

import (
	"errors"
	"testing"
	uc "github.com/vinihss/bodego-api/internal/usecases/user"
	"github.com/vinihss/bodego-api/internal/domain/user"
)

type mockUserRepoDelete struct {
	findByID    user.User
	findByIDErr error
	deleteErr   error
}

func (m *mockUserRepoDelete) Create(entity user.User) (user.User, error) { return entity, nil }
func (m *mockUserRepoDelete) Delete(id uint) error { return m.deleteErr }
func (m *mockUserRepoDelete) FindByID(id uint) (user.User, error) { return m.findByID, m.findByIDErr }
func (m *mockUserRepoDelete) Update(entity user.User) (user.User, error) { return entity, nil }
func (m *mockUserRepoDelete) FindAll(offset, size int) ([]user.User, error) { return nil, nil }

func TestDeleteUserUseCase_Execute(t *testing.T) {
    repo := &mockUserRepoDelete{findByID: user.User{ID: 1}}
	usecase := uc.NewDeleteUserUseCase(repo)
	err := usecase.Execute(1)
	if err != nil {
		t.Fatalf("esperado erro nulo, recebeu: %v", err)
	}
}

func TestDeleteUserUseCase_Execute_FindByIDError(t *testing.T) {
	repo := &mockUserRepoDelete{findByIDErr: errors.New("não encontrado")}
	usecase := uc.NewDeleteUserUseCase(repo)
	err := usecase.Execute(2)
	if err == nil {
		t.Error("esperado erro, recebeu nil")
	}
}

func TestDeleteUserUseCase_Execute_DeleteError(t *testing.T) {
	repo := &mockUserRepoDelete{findByID: user.User{ID: 3}, deleteErr: errors.New("erro ao deletar")}
	usecase := uc.NewDeleteUserUseCase(repo)
	err := usecase.Execute(3)
	if err == nil {
		t.Error("esperado erro, recebeu nil")
	}
}
