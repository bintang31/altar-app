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
	InquiryLoketTagihanAirByNosamb(*entity.InputInquiryPelanggan) ([]entity.RekairDetail, error)
	UpdateDrdByNosamb(*entity.Drd) (*entity.Drd, map[string]string)
	UpdateNonAirByNosamb(*entity.Nonair) (*entity.Nonair, map[string]string)
	InquiryLoketTagihanNonAirByNosamb(*entity.InputInquiryPelanggan) ([]entity.NonAirDetail, error)
	InquiryLoketAngsuranByNosamb(*entity.InputInquiryPelanggan) ([]entity.AngsuranDetail, error)
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
