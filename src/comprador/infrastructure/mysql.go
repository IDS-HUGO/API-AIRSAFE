package infrastructure

import (
	"database/sql"
	"errors"
	"time"

	"apiusersafe/src/comprador/domain"
)

type MySQLCompradorRepository struct {
	DB *sql.DB
}

func NewMySQLCompradorRepository(db *sql.DB) domain.CompradorRepository {
	return &MySQLCompradorRepository{DB: db}
}

func (r *MySQLCompradorRepository) CreateComprador(comprador *domain.Comprador) error {
	query := "INSERT INTO compradores (usuario, contrasena, telefono, email, fecha_creado) VALUES (?, ?, ?, ?, ?)"
	_, err := r.DB.Exec(query, comprador.Usuario, comprador.Contrasena, comprador.Telefono, comprador.Email, time.Now())
	return err
}

func (r *MySQLCompradorRepository) GetCompradorByID(id int) (*domain.Comprador, error) {
	query := "SELECT id, usuario, telefono, email, fecha_creado FROM compradores WHERE id = ?"
	row := r.DB.QueryRow(query, id)
	comprador := &domain.Comprador{}
	if err := row.Scan(&comprador.ID, &comprador.Usuario, &comprador.Telefono, &comprador.Email, &comprador.FechaCreado); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return comprador, nil
}

func (r *MySQLCompradorRepository) UpdateComprador(comprador *domain.Comprador) error {
	query := "UPDATE compradores SET usuario=?, telefono=?, email=? WHERE id=?"
	_, err := r.DB.Exec(query, comprador.Usuario, comprador.Telefono, comprador.Email, comprador.ID)
	return err
}

func (r *MySQLCompradorRepository) DeleteComprador(id int) error {
	query := "DELETE FROM compradores WHERE id=?"
	_, err := r.DB.Exec(query, id)
	return err
}

func (r *MySQLCompradorRepository) ListCompradores() ([]domain.Comprador, error) {
	query := "SELECT id, usuario, telefono, email, fecha_creado FROM compradores"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var compradores []domain.Comprador
	for rows.Next() {
		comprador := domain.Comprador{}
		if err := rows.Scan(&comprador.ID, &comprador.Usuario, &comprador.Telefono, &comprador.Email, &comprador.FechaCreado); err != nil {
			return nil, err
		}
		compradores = append(compradores, comprador)
	}
	return compradores, nil
}
