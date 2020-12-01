package repository

import (
	"altar-app/domain/entity"
)

//PenagihanRepository : Penagihan collection of methods that the infrastructure
type PenagihanRepository interface {
	GetPenagihans() ([]entity.Penagihan, error)
	GetPenagihansByUserPDAM(uint64) ([]entity.PenagihansSrKolektif, error)
}
