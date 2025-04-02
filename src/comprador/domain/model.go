package domain

import "time"

type Comprador struct {
	ID          int       `json:"id"`
	Usuario     string    `json:"usuario"`
	Contrasena  string    `json:"-"`
	Telefono    string    `json:"telefono"`
	Email       string    `json:"email"`
	FechaCreado time.Time `json:"fecha_creado"`
}
