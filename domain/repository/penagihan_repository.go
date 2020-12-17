package repository

import (
	"altar-app/domain/entity"
)

//PenagihanRepository : Penagihan collection of methods that the infrastructure
type PenagihanRepository interface {
	GetPenagihans() ([]entity.Penagihan, error)
	GetPenagihansByUserPDAM(uint64) ([]entity.PenagihansSrKolektif, error)
	GetPenagihanByNosamb(string) (*entity.Penagihan, error)
	BayarTagihanByNosamb(*entity.Bayar) (*entity.ResponseLoket, map[string]string)
	GetPenagihanByParam(*entity.PenagihansParams) ([]entity.PenagihansSrKolektif, error)
}
