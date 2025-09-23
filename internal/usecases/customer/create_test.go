package customer

import (
	"errors"
	"reflect"
	"testing"

	"github.com/vinihss/bodego-api/internal/domain/customer"
)

type mockCustomerRepo struct {
	created customer.Customer
	createErr error
}

func (m *mockCustomerRepo) Create(entity customer.Customer) (customer.Customer, error) {
	m.created = entity
	return entity, m.createErr
}
func (m *mockCustomerRepo) Delete(id uint) error { return nil }
func (m *mockCustomerRepo) FindByID(id uint) (customer.Customer, error) { return customer.Customer{}, nil }
func (m *mockCustomerRepo) Update(entity customer.Customer) (customer.Customer, error) { return entity, nil }
func (m *mockCustomerRepo) FindAll(int, int) ([]customer.Customer, error) { return nil, nil }

func TestCreateCustomerUseCase_Execute(t *testing.T) {
	repo := &mockCustomerRepo{}
	uc := NewCreateCustomerUseCase(repo)
	input := CreateCustomerInput{Name: "João", Email: "joao@email.com"}

	result, err := uc.Execute(input)
	if err != nil {
		t.Fatalf("esperado erro nulo, recebeu: %v", err)
	}
	if !reflect.DeepEqual(result, repo.created) {
		t.Errorf("esperado: %+v, recebeu: %+v", repo.created, result)
	}
}

func TestCreateCustomerUseCase_Execute_Error(t *testing.T) {
	repo := &mockCustomerRepo{createErr: errors.New("erro ao criar")}
	uc := NewCreateCustomerUseCase(repo)
	input := CreateCustomerInput{Name: "Maria", Email: "maria@email.com"}

	_, err := uc.Execute(input)
	if err == nil {
		t.Error("esperado erro, recebeu nil")
	}
}
