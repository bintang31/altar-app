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

//GetTagihanAirPelanggansByNosamb : Get Data Tagihan Air by Nosamb
func (r *PelangganRepo) GetTagihanAirPelanggansByNosamb(nosamb string) ([]entity.Drd, error) {
	var tagihanair []entity.Drd
	err := r.db.Debug().Where("nosamb = ?", nosamb).Find(&tagihanair).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("pelanggan not found")
	}
	return tagihanair, nil
}

//GetTagihanNonAirPelanggansByNosamb : Get Data Tagihan Nonair by Nosamb
func (r *PelangganRepo) GetTagihanNonAirPelanggansByNosamb(nosamb string) ([]entity.Nonair, error) {
	var tagihannonair []entity.Nonair
	err := r.db.Debug().Where("nomor = ?", nosamb).Find(&tagihannonair).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("pelanggan not found")
	}
	return tagihannonair, nil
}
