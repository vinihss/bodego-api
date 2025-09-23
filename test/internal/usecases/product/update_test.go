package product

import (
	"errors"
	"reflect"
	"testing"
	uc "github.com/vinihss/bodego-api/internal/usecases/product"
	"github.com/vinihss/bodego-api/internal/domain/product"
)

type mockProductRepoUpdate struct {
	updated product.Product
	updateErr error
}

func (m *mockProductRepoUpdate) Create(entity product.Product) (product.Product, error) { return entity, nil }
func (m *mockProductRepoUpdate) Delete(id uint) error { return nil }
func (m *mockProductRepoUpdate) FindByID(id uint) (product.Product, error) { return product.Product{}, nil }
func (m *mockProductRepoUpdate) Update(entity product.Product) (product.Product, error) {
	m.updated = entity
	return entity, m.updateErr
}
func (m *mockProductRepoUpdate) FindAll(int, int) ([]product.Product, error) { return nil, nil }

func TestUpdateProductUseCase_Execute(t *testing.T) {
	repo := &mockProductRepoUpdate{}
	usecase := uc.NewUpdateProductUseCase(repo)
	input := uc.UpdateProductInput{ID: 1, Name: "Produto", Price: 10.0, Description: "desc"}
	result, err := usecase.Execute(input)
	if err != nil {
		t.Fatalf("esperado erro nulo, recebeu: %v", err)
	}
	if !reflect.DeepEqual(result, repo.updated) {
		t.Errorf("esperado: %+v, recebeu: %+v", repo.updated, result)
	}
}

func TestUpdateProductUseCase_Execute_Error(t *testing.T) {
	repo := &mockProductRepoUpdate{updateErr: errors.New("erro ao atualizar")}
	usecase := uc.NewUpdateProductUseCase(repo)
	input := uc.UpdateProductInput{ID: 2, Name: "Produto", Price: 10.0, Description: "desc"}
	_, err := usecase.Execute(input)
	if err == nil {
		t.Error("esperado erro, recebeu nil")
	}
}
