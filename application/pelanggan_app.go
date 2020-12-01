package application

import (
	"altar-app/domain/entity"
	"altar-app/domain/repository"
)

type pelangganApp struct {
	pl repository.PelangganRepository
}

//pelangganApp implements the PelangganAppInterface
var _ PelangganAppInterface = &pelangganApp{}

//PelangganAppInterface : App Interface pelanggan repo
type PelangganAppInterface interface {
	GetPelanggans() ([]entity.Pelanggan, error)
}

func (p *pelangganApp) GetPelanggans() ([]entity.Pelanggan, error) {
	return p.pl.GetPelanggans()
}
