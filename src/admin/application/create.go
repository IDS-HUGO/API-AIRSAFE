package application

import (
	"apiusersafe/src/admin/domain"
)

type CreateAdminService struct {
	Repo domain.AdminRepository
}

func (s *CreateAdminService) Execute(admin *domain.Admin) error {
	return s.Repo.Create(admin)
}
