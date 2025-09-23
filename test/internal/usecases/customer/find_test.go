package customer

import (
	"errors"
	"reflect"
	"testing"
	uc "github.com/vinihss/bodego-api/internal/usecases/customer"
	"github.com/vinihss/bodego-api/internal/domain/customer"
)

type mockCustomerRepoFind struct {
	findByID    customer.Customer
	findByIDErr error
	findAll     []customer.Customer
	findAllErr  error
}

func (m *mockCustomerRepoFind) Create(entity customer.Customer) (customer.Customer, error) { return entity, nil }
func (m *mockCustomerRepoFind) Delete(id uint) error { return nil }
func (m *mockCustomerRepoFind) FindByID(id uint) (customer.Customer, error) { return m.findByID, m.findByIDErr }
func (m *mockCustomerRepoFind) Update(entity customer.Customer) (customer.Customer, error) { return entity, nil }
func (m *mockCustomerRepoFind) FindAll(offset, size int) ([]customer.Customer, error) { return m.findAll, m.findAllErr }

func TestFindCustomerUseCase_Execute(t *testing.T) {
    repo := &mockCustomerRepoFind{findByID: customer.Customer{ID: 1, Name: "Teste", Email: "teste@email.com"}}
	usecase := uc.NewFindCustomerUseCase(repo)
	result, err := usecase.Execute(1)
	if err != nil {
		t.Fatalf("esperado erro nulo, recebeu: %v", err)
	}
	if !reflect.DeepEqual(result, repo.findByID) {
		t.Errorf("esperado: %+v, recebeu: %+v", repo.findByID, result)
	}
}

func TestFindCustomerUseCase_Execute_Error(t *testing.T) {
	repo := &mockCustomerRepoFind{findByIDErr: errors.New("erro ao buscar")}
	usecase := uc.NewFindCustomerUseCase(repo)
	_, err := usecase.Execute(2)
	if err == nil {
		t.Error("esperado erro, recebeu nil")
	}
}

func TestFindCustomerUseCase_ExecuteAll(t *testing.T) {
	repo := &mockCustomerRepoFind{findAll: []customer.Customer{{ID: 1}, {ID: 2}}}
	usecase := uc.NewFindCustomerUseCase(repo)
	result, err := usecase.ExecuteAll(1, 2)
	if err != nil {
		t.Fatalf("esperado erro nulo, recebeu: %v", err)
	}
	if !reflect.DeepEqual(result, repo.findAll) {
		t.Errorf("esperado: %+v, recebeu: %+v", repo.findAll, result)
	}
}

func TestFindCustomerUseCase_ExecuteAll_Error(t *testing.T) {
	repo := &mockCustomerRepoFind{findAllErr: errors.New("erro ao buscar todos")}
	usecase := uc.NewFindCustomerUseCase(repo)
	_, err := usecase.ExecuteAll(1, 2)
	if err == nil {
		t.Error("esperado erro, recebeu nil")
	}
}
