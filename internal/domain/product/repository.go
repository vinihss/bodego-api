package product

type Repository interface {
	Create(f Product) (Product, error)
	FindByID(id string) (*Product, error)
	Delete(id string) (*Product, error)
	Update(f Product) (*Product, error)
}
