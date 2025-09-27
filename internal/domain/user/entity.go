package user

type User struct {
	ID       uint
	Name     string
	Email    string
	Password string
	Role     Role
}

type Role int

const (
	SysAdmin  Role = 0
	Admin     Role = 1
	Assistent Role = 2
)
