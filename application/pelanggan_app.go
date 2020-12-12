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
	GetTagihanAirPelanggansByNosamb(string) ([]entity.Drd, error)
	GetTagihanNonAirPelanggansByNosamb(string) ([]entity.Nonair, error)
}

func (p *pelangganApp) GetPelanggans() ([]entity.Pelanggan, error) {
	return p.pl.GetPelanggans()
}

func (p *pelangganApp) GetTagihanAirPelanggansByNosamb(nosamb string) ([]entity.Drd, error) {
	return p.pl.GetTagihanAirPelanggansByNosamb(nosamb)
}

func (p *pelangganApp) GetTagihanNonAirPelanggansByNosamb(nosamb string) ([]entity.Nonair, error) {
	return p.pl.GetTagihanNonAirPelanggansByNosamb(nosamb)
}
