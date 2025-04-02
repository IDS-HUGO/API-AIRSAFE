package application

import (
	"apiusersafe/src/admin/domain"
)

type ListAdminsService struct {
	Repo domain.AdminRepository
}

func (s *ListAdminsService) Execute() ([]domain.Admin, error) {
	return s.Repo.List()
}
