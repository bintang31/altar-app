package repository

import (
	"altar-app/domain/entity"
)

//PelangganRepository : Pelanggan collection of methods that the infrastructure
type PelangganRepository interface {
	GetPelanggans() ([]entity.Pelanggan, error)
	GetTagihanAirPelanggansByNosamb(string) ([]entity.Drd, error)
	GetTagihanNonAirPelanggansByNosamb(string) ([]entity.Nonair, error)
	InquiryLoketTagihanAirByNosamb(*entity.InputInquiryPelanggan) ([]entity.RekairDetail, error)
}
