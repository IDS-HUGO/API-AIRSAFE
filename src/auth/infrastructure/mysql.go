package infrastructure

import (
	"apiusersafe/src/auth/domain"
	"database/sql"
)

type MySQLUserRepository struct {
	DB *sql.DB
}

func NewMySQLUserRepository(db *sql.DB) *MySQLUserRepository {
	return &MySQLUserRepository{DB: db}
}

func (r *MySQLUserRepository) ValidateUser(usuario, contrasena string) (*domain.User, error) {
	user := &domain.User{}
	query := `SELECT id, usuario, role FROM usuarios WHERE usuario = ? AND contrasena = ?`
	err := r.DB.QueryRow(query, usuario, contrasena).Scan(&user.ID, &user.Usuario, &user.Role)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Add this method to MySQLUserRepository
func (r *MySQLUserRepository) CreateUser(user *domain.User) error {
	query := "INSERT INTO usuarios (usuario, contrasena, role) VALUES (?, ?, ?)"
	_, err := r.DB.Exec(query, user.Usuario, user.Contrasena, user.Role)
	return err
}
