package application

import (
	"altar-app/domain/entity"
	"altar-app/domain/repository"
)

type penagihanApp struct {
	pn repository.PenagihanRepository
}

//pelangganApp implements the PelangganAppInterface
var _ PenagihanAppInterface = &penagihanApp{}

//PenagihanAppInterface : App Interface pelanggan repo
type PenagihanAppInterface interface {
	GetPenagihans() ([]entity.Penagihan, error)
	GetPenagihansByUserPDAM(uint64) ([]entity.PenagihansSrKolektif, error)
}

func (p *penagihanApp) GetPenagihans() ([]entity.Penagihan, error) {
	return p.pn.GetPenagihans()
}

func (p *penagihanApp) GetPenagihansByUserPDAM(userID uint64) ([]entity.PenagihansSrKolektif, error) {
	return p.pn.GetPenagihansByUserPDAM(userID)
}