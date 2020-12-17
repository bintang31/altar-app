package persistence

import (
	"altar-app/domain/entity"
	"altar-app/domain/repository"
	"errors"
	"github.com/jinzhu/gorm"
)

//PetugasRepo : Call DB
type PetugasRepo struct {
	db *gorm.DB
}

//NewPetugasRepository : Petugas Repository
func NewPetugasRepository(db *gorm.DB) *PetugasRepo {
	return &PetugasRepo{db}
}

//PelangganRepo implements the repository.PelangganRepo interface
var _ repository.PetugasRepository = &PetugasRepo{}

//GetPetugas : Get All Data Petugas
func (p *PetugasRepo) GetPetugas() ([]entity.Petugas, error) {
	var petugas []entity.Petugas
	err := p.db.Debug().Find(&petugas).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("petugas not found")
	}
	return petugas, nil
}

//GetProfilePetugas : Get Data Profile Petugas
func (p *PetugasRepo) GetProfilePetugas(id uint64) (*entity.Petugas, error) {
	var petugas entity.Petugas
	err := p.db.Debug().Where("id = ?", id).Find(&petugas).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("petugas not found")
	}
	return &petugas, nil
}

//GetTagihanAirByPetugas : Get Data Tagihan Air per Petugas
func (p *PetugasRepo) GetTagihanAirByPetugas(id uint64) ([]entity.Drd, error) {
	var tagihanair []entity.Drd
	err := p.db.Debug().Table("petugas_rayons").Select("penagihans_sr_kolektifs.nosamb,"+
		"drds.periode,drds.pdam,drds.pakai,"+
		"drds.total").Joins("join penagihans_sr_kolektifs ON penagihans_sr_kolektifs.kode_rayon = "+
		"petugas_rayons.rayon join drds ON (drds.nosamb = penagihans_sr_kolektifs.nosamb and drds.pdam = penagihans_sr_kolektifs.kode_pdam)").Where("petugas_rayons.petugas = ? AND penagihans_sr_kolektifs.status_billing = ?", id, "BELUM TERBAYAR").Order("drds.periode asc").Find(&tagihanair).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("Tagihan not found")
	}
	return tagihanair, nil
}

//GetTagihanNonAirByPetugas : Get Data Tagihan Non Air per Petugas
func (p *PetugasRepo) GetTagihanNonAirByPetugas(id uint64) ([]entity.Drd, error) {
	var tagihanair []entity.Drd
	err := p.db.Debug().Table("petugas_rayons").Select("penagihans_sr_kolektifs.nosamb,"+
		"drds.periode,"+
		"drds.total").Joins("join penagihans_sr_kolektifs ON penagihans_sr_kolektifs.kode_rayon = "+
		"petugas_rayons.rayon join drds ON (drds.nosamb = penagihans_sr_kolektifs.nosamb and drds.pdam = penagihans_sr_kolektifs.kode_pdam)").Where("petugas_rayons.petugas = ? AND penagihans_sr_kolektifs.status_billing = ?", id, "BELUM TERBAYAR").Order("drds.periode asc").Find(&tagihanair).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("Tagihan not found")
	}
	return tagihanair, nil
}
