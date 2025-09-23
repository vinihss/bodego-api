package tab

import (
	"errors"
	"reflect"
	"testing"
	uc "github.com/vinihss/bodego-api/internal/usecases/tab"
	"github.com/vinihss/bodego-api/internal/domain/tab"
)

type mockTabRepoCreate struct {
	created tab.Tab
	createErr error
}

func (m *mockTabRepoCreate) Create(entity tab.Tab) (tab.Tab, error) {
	m.created = entity
	return entity, m.createErr
}
func (m *mockTabRepoCreate) Delete(id uint) error { return nil }
func (m *mockTabRepoCreate) FindByID(id uint) (tab.Tab, error) { return tab.Tab{}, nil }
func (m *mockTabRepoCreate) Update(entity tab.Tab) (tab.Tab, error) { return entity, nil }
func (m *mockTabRepoCreate) FindAll(int, int) ([]tab.Tab, error) { return nil, nil }

func TestCreateTabUseCase_Execute(t *testing.T) {
	repo := &mockTabRepoCreate{}
	usecase := uc.NewCreateTabUseCase(repo)
	input := uc.CreateTabInput{Name: "Tab", Email: "tab@email.com"}
	result, err := usecase.Execute(input)
	if err != nil {
		t.Fatalf("esperado erro nulo, recebeu: %v", err)
	}
	if !reflect.DeepEqual(result, repo.created) {
		t.Errorf("esperado: %+v, recebeu: %+v", repo.created, result)
	}
}

func TestCreateTabUseCase_Execute_Error(t *testing.T) {
	repo := &mockTabRepoCreate{createErr: errors.New("erro ao criar")}
	usecase := uc.NewCreateTabUseCase(repo)
	input := uc.CreateTabInput{Name: "Tab", Email: "tab@email.com"}
	_, err := usecase.Execute(input)
	if err == nil {
		t.Error("esperado erro, recebeu nil")
	}
}
