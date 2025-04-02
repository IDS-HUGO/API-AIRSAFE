package application

import (
	"apiusersafe/src/admin/domain"
)

type UpdateAdminService struct {
	Repo domain.AdminRepository
}

func (s *UpdateAdminService) Execute(admin *domain.Admin) error {
	return s.Repo.Update(admin)
}
