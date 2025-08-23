package user

type Repository interface {
	Create(f User) (User, error)
	FindByID(id string) (*User, error)
	Delete(id string) (*User, error)
	Update(f User) (*User, error)
}
