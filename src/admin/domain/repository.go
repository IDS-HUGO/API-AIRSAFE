package domain

type AdminRepository interface {
	Create(admin *Admin) error
	Update(admin *Admin) error
	Delete(id int) error
	List() ([]Admin, error)
}
