package repository

import (
	"altar-app/domain/entity"
)

//UserRepository : User collection of methods that the infrastructure
type UserRepository interface {
	SaveUser(*entity.User) (*entity.User, map[string]string)
	GetUser(uint64) (*entity.User, error)
	GetUsers() ([]entity.User, error)
	GetUserByEmailAndPassword(*entity.User) (*entity.User, map[string]string)
	GetUserByUserNamelAndPassword(*entity.User) (*entity.User, map[string]string)
}
