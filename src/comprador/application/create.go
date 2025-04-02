package application

import "apiusersafe/src/comprador/domain"

type CreateCompradorService struct {
	Repo domain.CompradorRepository
}

func NewCreateCompradorService(repo domain.CompradorRepository) *CreateCompradorService {
	return &CreateCompradorService{Repo: repo}
}

func (s *CreateCompradorService) Execute(comprador *domain.Comprador) error {
	return s.Repo.CreateComprador(comprador)
}
