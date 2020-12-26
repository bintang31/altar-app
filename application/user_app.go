package application

import (
	"altar-app/domain/entity"
	"altar-app/domain/repository"
)

type userApp struct {
	us repository.UserRepository
}

//UserApp implements the UserAppInterface
var _ UserAppInterface = &userApp{}

//UserAppInterface : Interfacing User App to Repository
type UserAppInterface interface {
	SaveUser(*entity.User) (*entity.User, map[string]string)
	GetUsers() ([]entity.User, error)
	GetUser(uint64) (*entity.User, error)
	GetUserByEmailAndPassword(*entity.User) (*entity.User, map[string]string)
	GetUserByUserNamelAndPassword(*entity.User) (*entity.User, map[string]string)
	UpdateUser(*entity.User) (*entity.User, map[string]string)
}

func (u *userApp) SaveUser(user *entity.User) (*entity.User, map[string]string) {
	return u.us.SaveUser(user)
}

func (u *userApp) GetUser(userID uint64) (*entity.User, error) {
	return u.us.GetUser(userID)
}

func (u *userApp) GetUsers() ([]entity.User, error) {
	return u.us.GetUsers()
}

func (u *userApp) GetUserByEmailAndPassword(user *entity.User) (*entity.User, map[string]string) {
	return u.us.GetUserByEmailAndPassword(user)
}

func (u *userApp) GetUserByUserNamelAndPassword(user *entity.User) (*entity.User, map[string]string) {
	return u.us.GetUserByUserNamelAndPassword(user)
}

func (u *userApp) UpdateUser(user *entity.User) (*entity.User, map[string]string) {
	return u.us.UpdateUser(user)
}
