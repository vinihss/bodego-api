package product

import (
	"errors"
	"reflect"
	"testing"
	uc "github.com/vinihss/bodego-api/internal/usecases/product"
	"github.com/vinihss/bodego-api/internal/domain/product"
)

type mockProductRepoFind struct {
	findByID    product.Product
	findByIDErr error
	findAll     []product.Product
	findAllErr  error
}

func (m *mockProductRepoFind) Create(entity product.Product) (product.Product, error) { return entity, nil }
func (m *mockProductRepoFind) Delete(id uint) error { return nil }
func (m *mockProductRepoFind) FindByID(id uint) (product.Product, error) { return m.findByID, m.findByIDErr }
func (m *mockProductRepoFind) Update(entity product.Product) (product.Product, error) { return entity, nil }
func (m *mockProductRepoFind) FindAll(offset, size int) ([]product.Product, error) { return m.findAll, m.findAllErr }

func TestFindProductUseCase_Execute(t *testing.T) {
    repo := &mockProductRepoFind{findByID: product.Product{ID: 1, Name: "Produto"}}
	usecase := uc.NewFindProductUseCase(repo)
	result, err := usecase.Execute(1)
	if err != nil {
		t.Fatalf("esperado erro nulo, recebeu: %v", err)
	}
	if !reflect.DeepEqual(result, repo.findByID) {
		t.Errorf("esperado: %+v, recebeu: %+v", repo.findByID, result)
	}
}

func TestFindProductUseCase_Execute_Error(t *testing.T) {
	repo := &mockProductRepoFind{findByIDErr: errors.New("erro ao buscar")}
	usecase := uc.NewFindProductUseCase(repo)
	_, err := usecase.Execute(2)
	if err == nil {
		t.Error("esperado erro, recebeu nil")
	}
}

func TestFindProductUseCase_ExecuteAll(t *testing.T) {
	repo := &mockProductRepoFind{findAll: []product.Product{{ID: 1}, {ID: 2}}}
	usecase := uc.NewFindProductUseCase(repo)
	result, err := usecase.ExecuteAll(1, 2)
	if err != nil {
		t.Fatalf("esperado erro nulo, recebeu: %v", err)
	}
	if !reflect.DeepEqual(result, repo.findAll) {
		t.Errorf("esperado: %+v, recebeu: %+v", repo.findAll, result)
	}
}

func TestFindProductUseCase_ExecuteAll_Error(t *testing.T) {
	repo := &mockProductRepoFind{findAllErr: errors.New("erro ao buscar todos")}
	usecase := uc.NewFindProductUseCase(repo)
	_, err := usecase.ExecuteAll(1, 2)
	if err == nil {
		t.Error("esperado erro, recebeu nil")
	}
}
