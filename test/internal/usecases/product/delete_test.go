package product

import (
	"errors"
	"testing"
	uc "github.com/vinihss/bodego-api/internal/usecases/product"
	"github.com/vinihss/bodego-api/internal/domain/product"
)

type mockProductRepoDelete struct {
	findByID    product.Product
	findByIDErr error
	deleteErr   error
}

func (m *mockProductRepoDelete) Create(entity product.Product) (product.Product, error) { return entity, nil }
func (m *mockProductRepoDelete) Delete(id uint) error { return m.deleteErr }
func (m *mockProductRepoDelete) FindByID(id uint) (product.Product, error) { return m.findByID, m.findByIDErr }
func (m *mockProductRepoDelete) Update(entity product.Product) (product.Product, error) { return entity, nil }
func (m *mockProductRepoDelete) FindAll(offset, size int) ([]product.Product, error) { return nil, nil }

func TestDeleteProductUseCase_Execute(t *testing.T) {
    repo := &mockProductRepoDelete{findByID: product.Product{ID: 1}}
	usecase := uc.NewDeleteProductUseCase(repo)
	err := usecase.Execute(1)
	if err != nil {
		t.Fatalf("esperado erro nulo, recebeu: %v", err)
	}
}

func TestDeleteProductUseCase_Execute_FindByIDError(t *testing.T) {
	repo := &mockProductRepoDelete{findByIDErr: errors.New("não encontrado")}
	usecase := uc.NewDeleteProductUseCase(repo)
	err := usecase.Execute(2)
	if err == nil {
		t.Error("esperado erro, recebeu nil")
	}
}

func TestDeleteProductUseCase_Execute_DeleteError(t *testing.T) {
	repo := &mockProductRepoDelete{findByID: product.Product{ID: 3}, deleteErr: errors.New("erro ao deletar")}
	usecase := uc.NewDeleteProductUseCase(repo)
	err := usecase.Execute(3)
	if err == nil {
		t.Error("esperado erro, recebeu nil")
	}
}
