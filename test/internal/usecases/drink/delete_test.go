package drink

import (
	"errors"
	"testing"
	uc "github.com/vinihss/bodego-api/internal/usecases/drink"
	"github.com/vinihss/bodego-api/internal/domain/drink"
)

type mockRepo struct {
	deleteErr error
}

func (m *mockRepo) Create(d drink.Drink) (drink.Drink, error) { return d, nil }
func (m *mockRepo) Delete(id uint) error { return m.deleteErr }
func (m *mockRepo) FindByID(id uint) (drink.Drink, error) { return drink.Drink{}, nil }
func (m *mockRepo) Update(d drink.Drink) (*drink.Drink, error) { return &d, nil }
func (m *mockRepo) FindAll() ([]drink.Drink, error) { return nil, nil }

func TestDeleteDrinkUseCase_Execute(t *testing.T) {
	repo := &mockRepo{}
	usecase := uc.DeleteDrinkUseCase{Repo: repo}
	err := usecase.Execute(1)
	if err != nil {
		t.Fatalf("esperado erro nulo, recebeu: %v", err)
	}
}

func TestDeleteDrinkUseCase_Execute_Error(t *testing.T) {
	repo := &mockRepo{deleteErr: errors.New("erro ao deletar")}
	usecase := uc.DeleteDrinkUseCase{Repo: repo}
	err := usecase.Execute(2)
	if err == nil {
		t.Error("esperado erro, recebeu nil")
	}
}
