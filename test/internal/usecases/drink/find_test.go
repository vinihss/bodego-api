package drink

import (
	"errors"
	"reflect"
	"testing"
	uc "github.com/vinihss/bodego-api/internal/usecases/drink"
	"github.com/vinihss/bodego-api/internal/domain/drink"
)

type mockRepo struct {
	findByID    drink.Drink
	findByIDErr error
	findAll     []drink.Drink
	findAllErr  error
}

func (m *mockRepo) Create(d drink.Drink) (drink.Drink, error) { return d, nil }
func (m *mockRepo) Delete(id uint) error { return nil }
func (m *mockRepo) FindByID(id uint) (drink.Drink, error) { return m.findByID, m.findByIDErr }
func (m *mockRepo) Update(d drink.Drink) (*drink.Drink, error) { return &d, nil }
func (m *mockRepo) FindAll() ([]drink.Drink, error) { return m.findAll, m.findAllErr }

func TestFindDrinkUseCase_Execute(t *testing.T) {
	repo := &mockRepo{findByID: drink.Drink{ID: 1, Name: "Coca"}}
	usecase := uc.FindDrinkUseCase{Repo: repo}
	result, err := usecase.Execute(1)
	if err != nil {
		t.Fatalf("esperado erro nulo, recebeu: %v", err)
	}
	if result == nil || !reflect.DeepEqual(*result, repo.findByID) {
		t.Errorf("esperado: %+v, recebeu: %+v", repo.findByID, result)
	}
}

func TestFindDrinkUseCase_Execute_Error(t *testing.T) {
	repo := &mockRepo{findByIDErr: errors.New("erro ao buscar")}
	usecase := uc.FindDrinkUseCase{Repo: repo}
	_, err := usecase.Execute(2)
	if err == nil {
		t.Error("esperado erro, recebeu nil")
	}
}

func TestFindDrinkUseCase_ExecuteAll(t *testing.T) {
	repo := &mockRepo{findAll: []drink.Drink{{ID: 1}, {ID: 2}}}
	usecase := uc.FindDrinkUseCase{Repo: repo}
	result, err := usecase.ExecuteAll()
	if err != nil {
		t.Fatalf("esperado erro nulo, recebeu: %v", err)
	}
	if !reflect.DeepEqual(result, repo.findAll) {
		t.Errorf("esperado: %+v, recebeu: %+v", repo.findAll, result)
	}
}

func TestFindDrinkUseCase_ExecuteAll_Error(t *testing.T) {
	repo := &mockRepo{findAllErr: errors.New("erro ao buscar todos")}
	usecase := uc.FindDrinkUseCase{Repo: repo}
	_, err := usecase.ExecuteAll()
	if err == nil {
		t.Error("esperado erro, recebeu nil")
	}
}
