package news

import "github.com/vinihss/bodego-api/internal/domain"

type News struct {
	domain.Entity
	ID          uint
	Name        string
	Price       float64
	Description string
}