package customer

import "github.com/vinihss/bodego-api/domain"

type Customer struct {
	domain.Entity
	ID    uint
	Name  string
	Email string
}
