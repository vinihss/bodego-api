package domain

type RepositoryInterface interface {
	Create(f Entity) (Entity, error)
	FindByID(id string) (*Entity, error)
	Delete(id string) (*Entity, error)
	Update(f Entity) (*Entity, error)
}

type Repository struct {
	RepositoryInterface
}
