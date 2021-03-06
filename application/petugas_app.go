package application

import (
	"altar-app/domain/entity"
	"altar-app/domain/repository"
)

type petugasApp struct {
	pt repository.PetugasRepository
}

//petugasApp implements the PelangganAppInterface
var _ PetugasAppInterface = &petugasApp{}

//PetugasAppInterface : App Interface pelanggan repo
type PetugasAppInterface interface {
	GetPetugas() ([]entity.Petugas, error)
	GetTagihanAirByPetugas(uint64) ([]entity.DrdbyPetugas, error)
	GetTagihanAirKolektifByPetugas(uint64) ([]entity.DrdbyPetugas, error)
	GetTagihanNonAirByPetugas(uint64) ([]entity.Nonair, error)
	GetProfilePetugas(uint64) (*entity.Petugas, error)
}

func (pe *petugasApp) GetPetugas() ([]entity.Petugas, error) {
	return pe.pt.GetPetugas()
}

func (pe *petugasApp) GetTagihanAirByPetugas(userID uint64) ([]entity.DrdbyPetugas, error) {
	return pe.pt.GetTagihanAirByPetugas(userID)
}

func (pe *petugasApp) GetTagihanAirKolektifByPetugas(userID uint64) ([]entity.DrdbyPetugas, error) {
	return pe.pt.GetTagihanAirKolektifByPetugas(userID)
}
func (pe *petugasApp) GetTagihanNonAirByPetugas(userID uint64) ([]entity.Nonair, error) {
	return pe.pt.GetTagihanNonAirByPetugas(userID)
}

func (pe *petugasApp) GetProfilePetugas(userID uint64) (*entity.Petugas, error) {
	return pe.pt.GetProfilePetugas(userID)
}
