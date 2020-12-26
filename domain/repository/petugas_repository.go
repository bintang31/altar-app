package repository

import (
	"altar-app/domain/entity"
)

//PetugasRepository : Petugas collection of methods that the infrastructure
type PetugasRepository interface {
	GetPetugas() ([]entity.Petugas, error)
	GetTagihanAirByPetugas(uint64) ([]entity.DrdbyPetugas, error)
	GetTagihanAirKolektifByPetugas(uint64) ([]entity.DrdbyPetugas, error)
	GetTagihanNonAirByPetugas(uint64) ([]entity.Nonair, error)
	GetProfilePetugas(uint64) (*entity.Petugas, error)
}
