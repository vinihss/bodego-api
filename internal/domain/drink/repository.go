package drink

type Repository interface {
	Create(drink Drink) (Drink, error)
	FindByID(id uint) (*Drink, error)
	FindAll() ([]Drink, error)
	Update(drink Drink) (*Drink, error)
	Delete(id uint) error
}
