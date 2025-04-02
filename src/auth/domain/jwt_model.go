package domain

type User struct {
	ID         int    `json:"id"`
	Usuario    string `json:"usuario"`
	Contrasena string `json:"-"`
	Role       string `json:"role"`
}

type UserRepository interface {
	ValidateUser(usuario, contrasena string) (*User, error)
	CreateUser(user *User) error
}
