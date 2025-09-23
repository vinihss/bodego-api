package tab

import (
	"errors"
	"testing"
	uc "github.com/vinihss/bodego-api/internal/usecases/tab"
	"github.com/vinihss/bodego-api/internal/domain/tab"
)

type mockRepo struct {
	findByID    tab.Tab
	findByIDErr error
	deleteErr   error
}

func (m *mockRepo) Create(entity tab.Tab) (tab.Tab, error) { return entity, nil }
func (m *mockRepo) Delete(id uint) error { return m.deleteErr }
func (m *mockRepo) FindByID(id uint) (tab.Tab, error) { return m.findByID, m.findByIDErr }
func (m *mockRepo) Update(entity tab.Tab) (tab.Tab, error) { return entity, nil }
func (m *mockRepo) FindAll(offset, size int) ([]tab.Tab, error) { return nil, nil }

func TestDeleteTabUseCase_Execute(t *testing.T) {
	repo := &mockRepo{findByID: tab.Tab{ID: 1}}
	usecase := uc.NewDeleteTabUseCase(repo)
	err := usecase.Execute(1)
	if err != nil {
		t.Fatalf("esperado erro nulo, recebeu: %v", err)
	}
}

func TestDeleteTabUseCase_Execute_FindByIDError(t *testing.T) {
	repo := &mockRepo{findByIDErr: errors.New("não encontrado")}
	usecase := uc.NewDeleteTabUseCase(repo)
	err := usecase.Execute(2)
	if err == nil {
		t.Error("esperado erro, recebeu nil")
	}
}

func TestDeleteTabUseCase_Execute_DeleteError(t *testing.T) {
	repo := &mockRepo{findByID: tab.Tab{ID: 3}, deleteErr: errors.New("erro ao deletar")}
	usecase := uc.NewDeleteTabUseCase(repo)
	err := usecase.Execute(3)
	if err == nil {
		t.Error("esperado erro, recebeu nil")
	}
}
