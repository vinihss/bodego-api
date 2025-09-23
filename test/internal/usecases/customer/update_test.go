package customer

import (
	"errors"
	"reflect"
	"testing"
	uc "github.com/vinihss/bodego-api/internal/usecases/customer"
	"github.com/vinihss/bodego-api/internal/domain/customer"
)

type mockRepo struct {
	updated customer.Customer
	updateErr error
}

func (m *mockRepo) Create(entity customer.Customer) (customer.Customer, error) { return entity, nil }
func (m *mockRepo) Delete(id uint) error { return nil }
func (m *mockRepo) FindByID(id uint) (customer.Customer, error) { return customer.Customer{}, nil }
func (m *mockRepo) Update(entity customer.Customer) (customer.Customer, error) {
	m.updated = entity
	return entity, m.updateErr
}
func (m *mockRepo) FindAll(int, int) ([]customer.Customer, error) { return nil, nil }

func TestUpdateCustomerUseCase_Execute(t *testing.T) {
	repo := &mockRepo{}
	usecase := uc.NewUpdateCustomerUseCase(repo)
	input := uc.UpdateCustomerInput{ID: 1, Name: "Novo", Email: "novo@email.com"}
	result, err := usecase.Execute(input)
	if err != nil {
		t.Fatalf("esperado erro nulo, recebeu: %v", err)
	}
	if !reflect.DeepEqual(result, repo.updated) {
		t.Errorf("esperado: %+v, recebeu: %+v", repo.updated, result)
	}
}

func TestUpdateCustomerUseCase_Execute_Error(t *testing.T) {
	repo := &mockRepo{updateErr: errors.New("erro ao atualizar")}
	usecase := uc.NewUpdateCustomerUseCase(repo)
	input := uc.UpdateCustomerInput{ID: 2, Name: "Erro", Email: "erro@email.com"}
	_, err := usecase.Execute(input)
	if err == nil {
		t.Error("esperado erro, recebeu nil")
	}
}
