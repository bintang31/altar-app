package persistence

import (
	"fmt"
	"github.com/jinzhu/gorm"
	//Gorm POSTGRES
	"altar-app/domain/entity"
	"altar-app/domain/repository"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//Repositories : Assign Repository
type Repositories struct {
	User      repository.UserRepository
	Role      repository.RoleRepository
	Pelanggan repository.PelangganRepository
	Penagihan repository.PenagihanRepository
	db        *gorm.DB
}

//NewRepositories : Register Repository
func NewRepositories(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) (*Repositories, error) {
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	db, err := gorm.Open(Dbdriver, DBURL)
	if err != nil {
		return nil, err
	}
	db.LogMode(true)

	return &Repositories{
		User:      NewUserRepository(db),
		Role:      NewRoleRepository(db),
		Pelanggan: NewPelangganRepository(db),
		Penagihan: NewPenagihanRepository(db),
		db:        db,
	}, nil
}

//Close : closes the  database connection
func (s *Repositories) Close() error {
	return s.db.Close()
}

//Automigrate : This migrate all tables
func (s *Repositories) Automigrate() error {
	return s.db.AutoMigrate(&entity.User{}, &entity.Role{}).Error
}
