package customer

import (
	"errors"
	"reflect"
	"testing"
	uc "github.com/vinihss/bodego-api/internal/usecases/customer"
	"github.com/vinihss/bodego-api/internal/domain/customer"
)

type mockCustomerRepoCreate struct {
	created customer.Customer
	createErr error
}

func (m *mockCustomerRepoCreate) Create(entity customer.Customer) (customer.Customer, error) {
	m.created = entity
	return entity, m.createErr
}
func (m *mockCustomerRepoCreate) Delete(id uint) error { return nil }
func (m *mockCustomerRepoCreate) FindByID(id uint) (customer.Customer, error) { return customer.Customer{}, nil }
func (m *mockCustomerRepoCreate) Update(entity customer.Customer) (customer.Customer, error) { return entity, nil }
func (m *mockCustomerRepoCreate) FindAll(int, int) ([]customer.Customer, error) { return nil, nil }

func TestCreateCustomerUseCase_Execute(t *testing.T) {
	repo := &mockCustomerRepoCreate{}
	usecase := uc.NewCreateCustomerUseCase(repo)
	input := uc.CreateCustomerInput{Name: "João", Email: "joao@email.com"}

	result, err := usecase.Execute(input)
	if err != nil {
		t.Fatalf("esperado erro nulo, recebeu: %v", err)
	}
	if !reflect.DeepEqual(result, repo.created) {
		t.Errorf("esperado: %+v, recebeu: %+v", repo.created, result)
	}
}

func TestCreateCustomerUseCase_Execute_Error(t *testing.T) {
	repo := &mockCustomerRepoCreate{createErr: errors.New("erro ao criar")}
	usecase := uc.NewCreateCustomerUseCase(repo)
	input := uc.CreateCustomerInput{Name: "Maria", Email: "maria@email.com"}

	_, err := usecase.Execute(input)
	if err == nil {
		t.Error("esperado erro, recebeu nil")
	}
}
