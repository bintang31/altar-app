package repository

import (
	"altar-app/domain/entity"
)

//RoleRepository : Role collection of methods that the infrastructure
type RoleRepository interface {
	GetRoles() ([]entity.Role, error)
}
