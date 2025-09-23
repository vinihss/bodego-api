package customer

import (
	"errors"
	"testing"
	uc "github.com/vinihss/bodego-api/internal/usecases/customer"
	"github.com/vinihss/bodego-api/internal/domain/customer"
)

type mockRepo struct {
	findByID    customer.Customer
	findByIDErr error
	deleteErr   error
}

func (m *mockRepo) Create(entity customer.Customer) (customer.Customer, error) { return entity, nil }
func (m *mockRepo) Delete(id uint) error { return m.deleteErr }
func (m *mockRepo) FindByID(id uint) (customer.Customer, error) { return m.findByID, m.findByIDErr }
func (m *mockRepo) Update(entity customer.Customer) (customer.Customer, error) { return entity, nil }
func (m *mockRepo) FindAll(offset, size int) ([]customer.Customer, error) { return nil, nil }

func TestDeleteCustomerUseCase_Execute(t *testing.T) {
	repo := &mockRepo{findByID: customer.Customer{ID: 1}}
	usecase := uc.NewDeleteCustomerUseCase(repo)
	err := usecase.Execute(1)
	if err != nil {
		t.Fatalf("esperado erro nulo, recebeu: %v", err)
	}
}

func TestDeleteCustomerUseCase_Execute_FindByIDError(t *testing.T) {
	repo := &mockRepo{findByIDErr: errors.New("não encontrado")}
	usecase := uc.NewDeleteCustomerUseCase(repo)
	err := usecase.Execute(2)
	if err == nil {
		t.Error("esperado erro, recebeu nil")
	}
}

func TestDeleteCustomerUseCase_Execute_DeleteError(t *testing.T) {
	repo := &mockRepo{findByID: customer.Customer{ID: 3}, deleteErr: errors.New("erro ao deletar")}
	usecase := uc.NewDeleteCustomerUseCase(repo)
	err := usecase.Execute(3)
	if err == nil {
		t.Error("esperado erro, recebeu nil")
	}
}
