package application

import "apiusersafe/src/comprador/domain"

type UpdateCompradorService struct {
	Repo domain.CompradorRepository
}

func NewUpdateCompradorService(repo domain.CompradorRepository) *UpdateCompradorService {
	return &UpdateCompradorService{Repo: repo}
}

func (s *UpdateCompradorService) Execute(comprador *domain.Comprador) error {
	return s.Repo.UpdateComprador(comprador)
}
