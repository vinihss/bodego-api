package product

import (
	"errors"
	"reflect"
	"testing"
	uc "github.com/vinihss/bodego-api/internal/usecases/product"
	"github.com/vinihss/bodego-api/internal/domain/product"
)

type mockRepo struct {
	created product.Product
	createErr error
}

func (m *mockRepo) Create(entity product.Product) (product.Product, error) {
	m.created = entity
	return entity, m.createErr
}
func (m *mockRepo) Delete(id uint) error { return nil }
func (m *mockRepo) FindByID(id uint) (product.Product, error) { return product.Product{}, nil }
func (m *mockRepo) Update(entity product.Product) (product.Product, error) { return entity, nil }
func (m *mockRepo) FindAll(int, int) ([]product.Product, error) { return nil, nil }

func TestCreateProductUseCase_Execute(t *testing.T) {
	repo := &mockRepo{}
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
	repo := &mockRepo{createErr: errors.New("erro ao criar")}
	usecase := uc.NewCreateProductUseCase(repo)
	input := uc.CreateProductInput{Name: "Produto", Price: 10.0, Description: "desc"}
	_, err := usecase.Execute(input)
	if err == nil {
		t.Error("esperado erro, recebeu nil")
	}
}
