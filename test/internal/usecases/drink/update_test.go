package drink

import (
	"errors"
	"reflect"
	"testing"
	uc "github.com/vinihss/bodego-api/internal/usecases/drink"
	"github.com/vinihss/bodego-api/internal/domain/drink"
)

type mockDrinkRepoUpdate struct {
	updated drink.Drink
	updateErr error
}

func (m *mockDrinkRepoUpdate) Create(d drink.Drink) (drink.Drink, error) { return d, nil }
func (m *mockDrinkRepoUpdate) Delete(id uint) error { return nil }
func (m *mockDrinkRepoUpdate) FindByID(id uint) (drink.Drink, error) { return drink.Drink{}, nil }
func (m *mockDrinkRepoUpdate) Update(d drink.Drink) (*drink.Drink, error) {
	m.updated = d
	return &d, m.updateErr
}
func (m *mockDrinkRepoUpdate) FindAll() ([]drink.Drink, error) { return nil, nil }

func TestUpdateDrinkUseCase_Execute(t *testing.T) {
	repo := &mockDrinkRepoUpdate{}
	usecase := uc.UpdateDrinkUseCase{Repo: repo}
	d := drink.Drink{ID: 1, Name: "Coca"}
	result, err := usecase.Execute(d)
	if err != nil {
		t.Fatalf("esperado erro nulo, recebeu: %v", err)
	}
	if result == nil || !reflect.DeepEqual(*result, repo.updated) {
		t.Errorf("esperado: %+v, recebeu: %+v", repo.updated, result)
	}
}

func TestUpdateDrinkUseCase_Execute_Error(t *testing.T) {
	repo := &mockDrinkRepoUpdate{updateErr: errors.New("erro ao atualizar")}
	usecase := uc.UpdateDrinkUseCase{Repo: repo}
	d := drink.Drink{ID: 2, Name: "Pepsi"}
	_, err := usecase.Execute(d)
	if err == nil {
		t.Error("esperado erro, recebeu nil")
	}
}
