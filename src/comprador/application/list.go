package application

import "apiusersafe/src/comprador/domain"

type ListCompradoresService struct {
	Repo domain.CompradorRepository
}

func NewListCompradoresService(repo domain.CompradorRepository) *ListCompradoresService {
	return &ListCompradoresService{Repo: repo}
}

func (s *ListCompradoresService) Execute() ([]domain.Comprador, error) {
	return s.Repo.ListCompradores()
}
