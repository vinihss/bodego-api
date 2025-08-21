package product

type Product struct {
	Entity
	ID          uint
	Name        string
	Price       float64
	Description string
}
