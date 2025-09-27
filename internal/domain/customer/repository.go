package customer

type Repository interface {
	Create(f Customer) (Customer, error)
	FindByID(id uint) (Customer, error)
	Delete(id uint) error
	Update(f Customer) (Customer, error)
	FindAll(int, size int) ([]Customer, error)
}
