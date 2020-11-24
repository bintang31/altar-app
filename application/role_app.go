package application

import (
	"altar-app/domain/entity"
	"altar-app/domain/repository"
)

type roleApp struct {
	rl repository.RoleRepository
}

//RoleApp implements the RoleAppInterface
var _ RoleAppInterface = &roleApp{}

//RoleAppInterface : Interface Role
type RoleAppInterface interface {
	GetRoles() ([]entity.Role, error)
}

func (r *roleApp) GetRoles() ([]entity.Role, error) {
	return r.rl.GetRoles()
}
