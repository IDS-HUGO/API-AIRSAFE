package domain

type CompradorRepository interface {
	CreateComprador(comprador *Comprador) error
	GetCompradorByID(id int) (*Comprador, error)
	UpdateComprador(comprador *Comprador) error
	DeleteComprador(id int) error
	ListCompradores() ([]Comprador, error)
}
