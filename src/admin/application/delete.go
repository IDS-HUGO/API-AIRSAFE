package application

import (
	"apiusersafe/src/admin/domain"
)

type DeleteAdminService struct {
	Repo domain.AdminRepository
}

func (s *DeleteAdminService) Execute(id int) error {
	return s.Repo.Delete(id)
}
