package persistence

import (
	"altar-app/domain/entity"
	"altar-app/domain/repository"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
)

//PenagihanRepo : Call DB
type PenagihanRepo struct {
	db *gorm.DB
}

//NewPenagihanRepository : Pelanggan Repository
func NewPenagihanRepository(db *gorm.DB) *PenagihanRepo {
	return &PenagihanRepo{db}
}

//PenagihanRepo implements the repository.PelangganRepo interface
var _ repository.PelangganRepository = &PelangganRepo{}

//GetPenagihans : Get All Data Penagihans
func (r *PenagihanRepo) GetPenagihans() ([]entity.Penagihan, error) {
	var penagihans []entity.Penagihan
	err := r.db.Debug().Find(&penagihans).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("pelanggan not found")
	}
	return penagihans, nil
}

//GetPenagihansByUserPDAM : Penagihan By User Login PDAM
func (r *PenagihanRepo) GetPenagihansByUserPDAM(id uint64) ([]entity.PenagihansSrKolektif, error) {
	var penagihans []entity.PenagihansSrKolektif
	fmt.Printf("userID :%+v\n", id)
	err := r.db.Debug().Find(&penagihans).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("pelanggan not found")
	}
	return penagihans, nil
}
