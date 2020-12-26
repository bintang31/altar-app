package persistence

import (
	"altar-app/domain/entity"
	"altar-app/domain/repository"
	"altar-app/infrastructure/security"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
)

//UserRepo : Call DB
type UserRepo struct {
	db *gorm.DB
}

//NewUserRepository : User Repository
func NewUserRepository(db *gorm.DB) *UserRepo {
	return &UserRepo{db}
}

//UserRepo implements the repository.UserRepository interface
var _ repository.UserRepository = &UserRepo{}

//SaveUser : Save User to DB
func (r *UserRepo) SaveUser(user *entity.User) (*entity.User, map[string]string) {
	dbErr := map[string]string{}
	err := r.db.Debug().Create(&user).Error
	if err != nil {
		//If the email is already taken
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			dbErr["email_taken"] = "email already taken"
			return nil, dbErr
		}
		//any other db error
		dbErr["db_error"] = "database error"
		return nil, dbErr
	}
	return user, nil
}

//UpdateUser : Update User to DB
func (r *UserRepo) UpdateUser(u *entity.User) (*entity.User, map[string]string) {
	dbErr := map[string]string{}
	var user entity.User
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	err := r.db.Debug().Model(&user).Where("id = ?", u.ID).Updates(map[string]interface{}{
		"pin":        u.Pin,
		"limit":      u.Limit,
		"updated_at": currentTime,
	}).Error
	fmt.Printf("userID :%+v\n", user)
	if err != nil {
		//any other db error
		dbErr["db_error"] = "database error"
		return nil, dbErr
	}
	return u, nil
}

//GetUser : Get User Detail from DB
func (r *UserRepo) GetUser(id uint64) (*entity.User, error) {
	var user entity.User
	err := r.db.Debug().Where("id = ?", id).Take(&user).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

//GetUsers : Get All User from DB
func (r *UserRepo) GetUsers() ([]entity.User, error) {
	var users []entity.User
	err := r.db.Debug().Find(&users).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("user not found")
	}
	return users, nil
}

//GetUserByEmailAndPassword :  Get User Profile by email and password
func (r *UserRepo) GetUserByEmailAndPassword(u *entity.User) (*entity.User, map[string]string) {
	var user entity.User
	dbErr := map[string]string{}
	err := r.db.Debug().Where("email = ?", u.Email).Take(&user).Error
	if gorm.IsRecordNotFoundError(err) {
		dbErr["no_user"] = "user not found"
		return nil, dbErr
	}
	if err != nil {
		dbErr["db_error"] = "database error"
		return nil, dbErr
	}
	//Verify the password
	err = security.VerifyPassword(user.Password, u.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		dbErr["incorrect_password"] = "incorrect password"
		return nil, dbErr
	}
	return &user, nil
}

//GetUserByUserNamelAndPassword : Get User Profile by Username
func (r *UserRepo) GetUserByUserNamelAndPassword(u *entity.User) (*entity.User, map[string]string) {
	var user entity.User
	dbErr := map[string]string{}
	err := r.db.Debug().Where("username = ?", u.Username).Take(&user).Error
	if gorm.IsRecordNotFoundError(err) {
		dbErr["no_user"] = "user not found"
		return nil, dbErr
	}
	if err != nil {
		dbErr["db_error"] = "database error"
		return nil, dbErr
	}
	//Verify the password
	err = security.VerifyPassword(user.Password, u.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		dbErr["incorrect_password"] = "incorrect password"
		return nil, dbErr
	}
	return &user, nil
}
