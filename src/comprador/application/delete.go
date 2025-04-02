package application

import "apiusersafe/src/comprador/domain"

type DeleteCompradorService struct {
	Repo domain.CompradorRepository
}

func NewDeleteCompradorService(repo domain.CompradorRepository) *DeleteCompradorService {
	return &DeleteCompradorService{Repo: repo}
}

func (s *DeleteCompradorService) Execute(id int) error {
	return s.Repo.DeleteComprador(id)
}
