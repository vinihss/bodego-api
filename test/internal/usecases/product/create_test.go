package product

import (
	"errors"
	"reflect"
	"testing"
	uc "github.com/vinihss/bodego-api/internal/usecases/product"
	"github.com/vinihss/bodego-api/internal/domain/product"
)

type mockProductRepoCreate struct {
	created product.Product
	createErr error
}

func (m *mockProductRepoCreate) Create(entity product.Product) (product.Product, error) {
	m.created = entity
	return entity, m.createErr
}
func (m *mockProductRepoCreate) Delete(id uint) error { return nil }
func (m *mockProductRepoCreate) FindByID(id uint) (product.Product, error) { return product.Product{}, nil }
func (m *mockProductRepoCreate) Update(entity product.Product) (product.Product, error) { return entity, nil }
func (m *mockProductRepoCreate) FindAll(int, int) ([]product.Product, error) { return nil, nil }

func TestCreateProductUseCase_Execute(t *testing.T) {
	repo := &mockProductRepoCreate{}
	usecase := uc.NewCreateProductUseCase(repo)
	input := uc.CreateProductInput{Name: "Produto", Price: 10.0, Description: "desc"}
	result, err := usecase.Execute(input)
	if err != nil {
		t.Fatalf("esperado erro nulo, recebeu: %v", err)
	}
	if !reflect.DeepEqual(result, repo.created) {
		t.Errorf("esperado: %+v, recebeu: %+v", repo.created, result)
	}
}

func TestCreateProductUseCase_Execute_Error(t *testing.T) {
	repo := &mockProductRepoCreate{createErr: errors.New("erro ao criar")}
	usecase := uc.NewCreateProductUseCase(repo)
	input := uc.CreateProductInput{Name: "Produto", Price: 10.0, Description: "desc"}
	_, err := usecase.Execute(input)
	if err == nil {
		t.Error("esperado erro, recebeu nil")
	}
}
