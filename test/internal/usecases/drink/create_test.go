package drink

import (
	"errors"
	"reflect"
	"testing"
	uc "github.com/vinihss/bodego-api/internal/usecases/drink"
	"github.com/vinihss/bodego-api/internal/domain/drink"
)

type mockDrinkRepoCreate struct {
	created drink.Drink
	createErr error
}

func (m *mockDrinkRepoCreate) Create(d drink.Drink) (drink.Drink, error) {
	m.created = d
	return d, m.createErr
}
func (m *mockDrinkRepoCreate) Delete(id uint) error { return nil }
func (m *mockDrinkRepoCreate) FindByID(id uint) (drink.Drink, error) { return drink.Drink{}, nil }
func (m *mockDrinkRepoCreate) Update(d drink.Drink) (*drink.Drink, error) { return &d, nil }
func (m *mockDrinkRepoCreate) FindAll() ([]drink.Drink, error) { return nil, nil }

func TestCreateDrinkUseCase_Execute(t *testing.T) {
	repo := &mockDrinkRepoCreate{}
	usecase := uc.CreateDrinkUseCase{Repo: repo}
	d := drink.Drink{ID: 1, Name: "Coca"}
	result, err := usecase.Execute(d)
	if err != nil {
		t.Fatalf("esperado erro nulo, recebeu: %v", err)
	}
	if !reflect.DeepEqual(result, repo.created) {
		t.Errorf("esperado: %+v, recebeu: %+v", repo.created, result)
	}
}

func TestCreateDrinkUseCase_Execute_Error(t *testing.T) {
	repo := &mockDrinkRepoCreate{createErr: errors.New("erro ao criar")}
	usecase := uc.CreateDrinkUseCase{Repo: repo}
	d := drink.Drink{ID: 2, Name: "Pepsi"}
	_, err := usecase.Execute(d)
	if err == nil {
		t.Error("esperado erro, recebeu nil")
	}
}
