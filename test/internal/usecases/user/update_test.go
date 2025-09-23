package user

import (
	"errors"
	"reflect"
	"testing"
	uc "github.com/vinihss/bodego-api/internal/usecases/user"
	"github.com/vinihss/bodego-api/internal/domain/user"
)

type mockUserRepoUpdate struct {
	updated user.User
	updateErr error
}

func (m *mockUserRepoUpdate) Create(entity user.User) (user.User, error) { return entity, nil }
func (m *mockUserRepoUpdate) Delete(id uint) error { return nil }
func (m *mockUserRepoUpdate) FindByID(id uint) (user.User, error) { return user.User{}, nil }
func (m *mockUserRepoUpdate) Update(entity user.User) (user.User, error) {
	m.updated = entity
	return entity, m.updateErr
}
func (m *mockUserRepoUpdate) FindAll(int, int) ([]user.User, error) { return nil, nil }

func TestUpdateUserUseCase_Execute(t *testing.T) {
	repo := &mockUserRepoUpdate{}
	usecase := uc.NewUpdateUserUseCase(repo)
	input := uc.UpdateUserInput{ID: 1, Name: "User", Email: "user@email.com"}
	result, err := usecase.Execute(input)
	if err != nil {
		t.Fatalf("esperado erro nulo, recebeu: %v", err)
	}
	if !reflect.DeepEqual(result, repo.updated) {
		t.Errorf("esperado: %+v, recebeu: %+v", repo.updated, result)
	}
}

func TestUpdateUserUseCase_Execute_Error(t *testing.T) {
	repo := &mockUserRepoUpdate{updateErr: errors.New("erro ao atualizar")}
	usecase := uc.NewUpdateUserUseCase(repo)
	input := uc.UpdateUserInput{ID: 2, Name: "User", Email: "user@email.com"}
	_, err := usecase.Execute(input)
	if err == nil {
		t.Error("esperado erro, recebeu nil")
	}
}
