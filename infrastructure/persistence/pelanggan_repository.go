package persistence

import (
	"altar-app/domain/entity"
	"altar-app/domain/repository"
	"errors"
	"github.com/jinzhu/gorm"
)

//PelangganRepo : Call DB
type PelangganRepo struct {
	db *gorm.DB
}

//NewPelangganRepository : Pelanggan Repository
func NewPelangganRepository(db *gorm.DB) *PelangganRepo {
	return &PelangganRepo{db}
}

//PelangganRepo implements the repository.PelangganRepo interface
var _ repository.PelangganRepository = &PelangganRepo{}

//GetPelanggans : Get All Data Pelanggan
func (r *PelangganRepo) GetPelanggans() ([]entity.Pelanggan, error) {
	var pelanggans []entity.Pelanggan
	err := r.db.Debug().Find(&pelanggans).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("pelanggan not found")
	}
	return pelanggans, nil
}
