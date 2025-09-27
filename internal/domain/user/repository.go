package user

type Repository interface {
	Create(user User) (User, error)
	FindByID(id uint) (User, error)
	FindByEmail(email string) ([]User, error)
	Update(user User) (User, error)
	Delete(id uint) error
}
