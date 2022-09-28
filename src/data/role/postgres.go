package role

import "database/sql"

type postgresRoleRepository struct {
	db *sql.DB
}

func (r *postgresRoleRepository) GetRoles() ([]Role, error) {
	panic("Unimplemented")
}
func (r *postgresRoleRepository) CreateRole(role *Role) error {
	panic("Unimplemented")
}
func (r *postgresRoleRepository) UpdateRole(role *Role) error {
	panic("Unimplemented")
}
func (r *postgresRoleRepository) DeleteRole(title string) error {
	panic("Unimplemented")
}
func (r *postgresRoleRepository) FindRoleByTitle(title string) (Role, error) {
	panic("Unimplemented")
}
