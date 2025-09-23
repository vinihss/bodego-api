package product

import (
	"errors"
	"testing"
	uc "github.com/vinihss/bodego-api/internal/usecases/product"
	"github.com/vinihss/bodego-api/internal/domain/product"
)

type mockRepo struct {
	findByID    product.Product
	findByIDErr error
	deleteErr   error
}

func (m *mockRepo) Create(entity product.Product) (product.Product, error) { return entity, nil }
func (m *mockRepo) Delete(id uint) error { return m.deleteErr }
func (m *mockRepo) FindByID(id uint) (product.Product, error) { return m.findByID, m.findByIDErr }
func (m *mockRepo) Update(entity product.Product) (product.Product, error) { return entity, nil }
func (m *mockRepo) FindAll(offset, size int) ([]product.Product, error) { return nil, nil }

func TestDeleteProductUseCase_Execute(t *testing.T) {
	repo := &mockRepo{findByID: product.Product{ID: 1}}
	usecase := uc.NewDeleteProductUseCase(repo)
	err := usecase.Execute(1)
	if err != nil {
		t.Fatalf("esperado erro nulo, recebeu: %v", err)
	}
}

func TestDeleteProductUseCase_Execute_FindByIDError(t *testing.T) {
	repo := &mockRepo{findByIDErr: errors.New("não encontrado")}
	usecase := uc.NewDeleteProductUseCase(repo)
	err := usecase.Execute(2)
	if err == nil {
		t.Error("esperado erro, recebeu nil")
	}
}

func TestDeleteProductUseCase_Execute_DeleteError(t *testing.T) {
	repo := &mockRepo{findByID: product.Product{ID: 3}, deleteErr: errors.New("erro ao deletar")}
	usecase := uc.NewDeleteProductUseCase(repo)
	err := usecase.Execute(3)
	if err == nil {
		t.Error("esperado erro, recebeu nil")
	}
}
