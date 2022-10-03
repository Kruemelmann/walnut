package role

import (
	"context"
	"log"

	"github.com/kruemelmann/walnut/src/data/dbutils"
)

type Role struct {
	ID    string
	Title string
}

type RoleRepository interface {
	GetRoles() ([]Role, error)
	CreateRole(role *Role) error
	UpdateRole(role *Role) error
	DeleteRole(title string) error
	FindRoleByTitle(title string) (Role, error)
}

func NewRoleRepository(ctx context.Context) (RoleRepository, error) {
	db, dbk, err := dbutils.ExtractDatabaseSession(ctx)
	if err != nil {
		log.Fatalf("Unable to extract db session from context %s\n", err)
	}

	switch dbk {
	case "postgres":
		return &postgresRoleRepository{
			db: db,
		}, nil
	default:
		panic("")
	}
}
