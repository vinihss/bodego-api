package database

import "reflect"

type RepositoryInterface interface {
	Create(model Model) (Model, error)
	FindByID(id uint) (Model, error)
	Delete(id uint) error
	Update(model Model) (Model, error)
}

type Repository struct {
}

func (r *Repository) Create(m *Model) error {
	var model = NewModelByTypeName(reflect.TypeOf(m).Name(), Models{})

	err := HydrateModel(entity, &product)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if err := r.db.Create(&model).Error; err != nil {

		err
	}

	return customer.Customer{
		ID:    model.ID,
		Name:  model.Name,
		Email: model.Email,
	}, nil
}
