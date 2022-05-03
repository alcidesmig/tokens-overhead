package repository

type RolesRepository interface {
	LoadRoles(path string) error
	GetRoles(offset, limit int) ([]string, error)
}
