package tab

import (
	"errors"
	"reflect"
	"testing"
	uc "github.com/vinihss/bodego-api/internal/usecases/tab"
	"github.com/vinihss/bodego-api/internal/domain/tab"
)

type mockRepo struct {
	findByID    tab.Tab
	findByIDErr error
	findAll     []tab.Tab
	findAllErr  error
}

func (m *mockRepo) Create(entity tab.Tab) (tab.Tab, error) { return entity, nil }
func (m *mockRepo) Delete(id uint) error { return nil }
func (m *mockRepo) FindByID(id uint) (tab.Tab, error) { return m.findByID, m.findByIDErr }
func (m *mockRepo) Update(entity tab.Tab) (tab.Tab, error) { return entity, nil }
func (m *mockRepo) FindAll(offset, size int) ([]tab.Tab, error) { return m.findAll, m.findAllErr }

func TestFindTabUseCase_Execute(t *testing.T) {
	repo := &mockRepo{findByID: tab.Tab{ID: 1, Name: "Tab"}}
	usecase := uc.NewFindTabUseCase(repo)
	result, err := usecase.Execute(1)
	if err != nil {
		t.Fatalf("esperado erro nulo, recebeu: %v", err)
	}
	if !reflect.DeepEqual(result, repo.findByID) {
		t.Errorf("esperado: %+v, recebeu: %+v", repo.findByID, result)
	}
}

func TestFindTabUseCase_Execute_Error(t *testing.T) {
	repo := &mockRepo{findByIDErr: errors.New("erro ao buscar")}
	usecase := uc.NewFindTabUseCase(repo)
	_, err := usecase.Execute(2)
	if err == nil {
		t.Error("esperado erro, recebeu nil")
	}
}

func TestFindTabUseCase_ExecuteAll(t *testing.T) {
	repo := &mockRepo{findAll: []tab.Tab{{ID: 1}, {ID: 2}}}
	usecase := uc.NewFindTabUseCase(repo)
	result, err := usecase.ExecuteAll(1, 2)
	if err != nil {
		t.Fatalf("esperado erro nulo, recebeu: %v", err)
	}
	if !reflect.DeepEqual(result, repo.findAll) {
		t.Errorf("esperado: %+v, recebeu: %+v", repo.findAll, result)
	}
}

func TestFindTabUseCase_ExecuteAll_Error(t *testing.T) {
	repo := &mockRepo{findAllErr: errors.New("erro ao buscar todos")}
	usecase := uc.NewFindTabUseCase(repo)
	_, err := usecase.ExecuteAll(1, 2)
	if err == nil {
		t.Error("esperado erro, recebeu nil")
	}
}
