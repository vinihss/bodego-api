package tab

import (
	"errors"
	"reflect"
	"testing"
	uc "github.com/vinihss/bodego-api/internal/usecases/tab"
	"github.com/vinihss/bodego-api/internal/domain/tab"
)

type mockRepo struct {
	updated tab.Tab
	updateErr error
}

func (m *mockRepo) Create(entity tab.Tab) (tab.Tab, error) { return entity, nil }
func (m *mockRepo) Delete(id uint) error { return nil }
func (m *mockRepo) FindByID(id uint) (tab.Tab, error) { return tab.Tab{}, nil }
func (m *mockRepo) Update(entity tab.Tab) (tab.Tab, error) {
	m.updated = entity
	return entity, m.updateErr
}
func (m *mockRepo) FindAll(int, int) ([]tab.Tab, error) { return nil, nil }

func TestUpdateTabUseCase_Execute(t *testing.T) {
	repo := &mockRepo{}
	usecase := uc.NewUpdateTabUseCase(repo)
	input := uc.UpdateTabInput{ID: 1, Name: "Tab", Email: "tab@email.com"}
	result, err := usecase.Execute(input)
	if err != nil {
		t.Fatalf("esperado erro nulo, recebeu: %v", err)
	}
	if !reflect.DeepEqual(result, repo.updated) {
		t.Errorf("esperado: %+v, recebeu: %+v", repo.updated, result)
	}
}

func TestUpdateTabUseCase_Execute_Error(t *testing.T) {
	repo := &mockRepo{updateErr: errors.New("erro ao atualizar")}
	usecase := uc.NewUpdateTabUseCase(repo)
	input := uc.UpdateTabInput{ID: 2, Name: "Tab", Email: "tab@email.com"}
	_, err := usecase.Execute(input)
	if err == nil {
		t.Error("esperado erro, recebeu nil")
	}
}
