package product

import "github.com/vinihss/bodego-api/internal/domain"

type Product struct {
	domain.EntityInterface
	ID          uint
	Name        string
	Price       float64
	Description string
}
