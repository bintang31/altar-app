package persistence

import (
	"altar-app/domain/entity"
	"altar-app/domain/repository"
	"errors"
	"github.com/jinzhu/gorm"
)

//RoleRepo : Call DB
type RoleRepo struct {
	db *gorm.DB
}

//NewRoleRepository : Role Repository
func NewRoleRepository(db *gorm.DB) *RoleRepo {
	return &RoleRepo{db}
}

//UserRepo implements the repository.UserRepository interface
var _ repository.RoleRepository = &RoleRepo{}

//GetRoles : Get All User from DB
func (r *RoleRepo) GetRoles() ([]entity.Role, error) {
	var roles []entity.Role
	err := r.db.Debug().Find(&roles).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("role not found")
	}
	return roles, nil
}
