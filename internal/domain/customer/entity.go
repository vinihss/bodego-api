package customer

import "github.com/vinihss/bodego-api/internal/domain"

type Customer struct {
	domain.Entity
	ID    uint
	Name  string
	Email string
}
