package tab

type Repository interface {
	Create(f Tab) (Tab, error)
	FindByID(id int) (*Tab, error)
	Delete(id string) (*Tab, error)
	Update(f Tab) (*Tab, error)
}
