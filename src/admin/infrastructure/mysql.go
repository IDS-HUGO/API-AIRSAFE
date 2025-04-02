package infrastructure

import (
	"apiusersafe/src/admin/domain"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type AdminRepositoryMysql struct {
	DB *sql.DB
}

func NewAdminRepositoryMysql(db *sql.DB) *AdminRepositoryMysql {
	return &AdminRepositoryMysql{DB: db}
}

func (r *AdminRepositoryMysql) Create(admin *domain.Admin) error {
	query := "INSERT INTO administradores (usuario, contrasena) VALUES (?, ?)"
	_, err := r.DB.Exec(query, admin.Usuario, admin.Contrasena)
	return err
}

func (r *AdminRepositoryMysql) Update(admin *domain.Admin) error {
	query := "UPDATE administradores SET usuario = ?, contrasena = ? WHERE id = ?"
	_, err := r.DB.Exec(query, admin.Usuario, admin.Contrasena, admin.ID)
	return err
}

func (r *AdminRepositoryMysql) Delete(id int) error {
	query := "DELETE FROM administradores WHERE id = ?"
	_, err := r.DB.Exec(query, id)
	return err
}

func (r *AdminRepositoryMysql) List() ([]domain.Admin, error) {
	query := "SELECT id, usuario, contrasena FROM administradores"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var admins []domain.Admin
	for rows.Next() {
		var admin domain.Admin
		if err := rows.Scan(&admin.ID, &admin.Usuario, &admin.Contrasena); err != nil {
			return nil, err
		}
		admins = append(admins, admin)
	}
	return admins, nil
}
