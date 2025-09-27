package customer

type Infos interface {
	GetName() string
	GetEmail() string
}

type Customer struct {
	ID    uint
	Name  string
	Email string
}

// Implementação dos métodos da interface
func (c Customer) GetName() string {
	return c.Name
}

func (c Customer) GetEmail() string {
	return c.Email
}
