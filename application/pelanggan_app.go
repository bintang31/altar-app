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
	GetTagihanNonAirPelanggansByNosamb(*entity.PeriodeNonair) ([]entity.Nonair, error)
	InquiryLoketTagihanAirByNosamb(*entity.InputInquiryPelanggan) ([]entity.RekairDetail, error)
	UpdateDrdByNosamb(*entity.Drd) (*entity.Drd, map[string]string)
	UpdateNonAirByNosamb(*entity.Nonair) (*entity.Nonair, map[string]string)
	InquiryLoketTagihanNonAirByNosamb(*entity.InputInquiryPelanggan) ([]entity.NonAirDetail, error)
	InquiryLoketAngsuranByNosamb(*entity.InputInquiryPelanggan) ([]entity.AngsuranDetail, error)
	GetTagihanNonAirPelanggansByPeriode(string) ([]map[string]interface{}, error)
}

func (p *pelangganApp) GetPelanggans() ([]entity.Pelanggan, error) {
	return p.pl.GetPelanggans()
}

func (p *pelangganApp) GetTagihanAirPelanggansByNosamb(nosamb string) ([]entity.Drd, error) {
	return p.pl.GetTagihanAirPelanggansByNosamb(nosamb)
}

func (p *pelangganApp) GetTagihanNonAirPelanggansByNosamb(u *entity.PeriodeNonair) ([]entity.Nonair, error) {
	return p.pl.GetTagihanNonAirPelanggansByNosamb(u)
}

func (p *pelangganApp) GetTagihanNonAirPelanggansByPeriode(nosamb string) ([]map[string]interface{}, error) {
	return p.pl.GetTagihanNonAirPelanggansByPeriode(nosamb)
}

func (p *pelangganApp) InquiryLoketTagihanAirByNosamb(u *entity.InputInquiryPelanggan) ([]entity.RekairDetail, error) {
	return p.pl.InquiryLoketTagihanAirByNosamb(u)
}

func (p *pelangganApp) UpdateDrdByNosamb(rd *entity.Drd) (*entity.Drd, map[string]string) {
	return p.pl.UpdateDrdByNosamb(rd)
}

func (p *pelangganApp) UpdateNonAirByNosamb(rd *entity.Nonair) (*entity.Nonair, map[string]string) {
	return p.pl.UpdateNonAirByNosamb(rd)
}

func (p *pelangganApp) InquiryLoketTagihanNonAirByNosamb(u *entity.InputInquiryPelanggan) ([]entity.NonAirDetail, error) {
	return p.pl.InquiryLoketTagihanNonAirByNosamb(u)
}

func (p *pelangganApp) InquiryLoketAngsuranByNosamb(u *entity.InputInquiryPelanggan) ([]entity.AngsuranDetail, error) {
	return p.pl.InquiryLoketAngsuranByNosamb(u)
}
