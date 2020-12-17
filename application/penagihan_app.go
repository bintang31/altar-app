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
	GetPenagihanByNosamb(string) (*entity.Penagihan, error)
	BayarTagihanByNosamb(*entity.Bayar) (*entity.ResponseLoket, map[string]string)
	GetPenagihanByParam(*entity.PenagihansParams) ([]entity.PenagihansSrKolektif, error)
}

func (p *penagihanApp) GetPenagihans() ([]entity.Penagihan, error) {
	return p.pn.GetPenagihans()
}

func (p *penagihanApp) GetPenagihansByUserPDAM(userID uint64) ([]entity.PenagihansSrKolektif, error) {
	return p.pn.GetPenagihansByUserPDAM(userID)
}

func (p *penagihanApp) GetPenagihanByNosamb(nosamb string) (*entity.Penagihan, error) {
	return p.pn.GetPenagihanByNosamb(nosamb)
}

func (p *penagihanApp) BayarTagihanByNosamb(u *entity.Bayar) (*entity.ResponseLoket, map[string]string) {
	return p.pn.BayarTagihanByNosamb(u)
}

func (p *penagihanApp) GetPenagihanByParam(t *entity.PenagihansParams) ([]entity.PenagihansSrKolektif, error) {
	return p.pn.GetPenagihanByParam(t)
}
