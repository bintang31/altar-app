package repository

import (
	"altar-app/domain/entity"
)

//PelangganRepository : Pelanggan collection of methods that the infrastructure
type PelangganRepository interface {
	GetPelanggans() ([]entity.Pelanggan, error)
	GetTagihanAirPelanggansByNosamb(string) ([]entity.Drd, error)
	GetTagihanNonAirPelanggansByNosamb(*entity.PeriodeNonair) ([]entity.Nonair, error)
	InquiryLoketTagihanAirByNosamb(*entity.InputInquiryPelanggan) ([]entity.RekairDetail, error)
	UpdateDrdByNosamb(*entity.Drd) (*entity.Drd, map[string]string)
	UpdateNonAirByNosamb(*entity.Nonair) (*entity.Nonair, map[string]string)
	InquiryLoketTagihanNonAirByNosamb(*entity.InputInquiryPelanggan) ([]entity.NonAirDetail, error)
	InquiryLoketAngsuranByNosamb(*entity.InputInquiryPelanggan) ([]entity.AngsuranDetail, error)
	GetTagihanNonAirPelanggansByPeriode(string) ([]map[string]interface{}, error)
	InsertAngsuranByNosamb(*entity.Nonair) (*entity.Nonair, map[string]string)
	GetRiwayatTagihanByNosamb(string) ([]map[string]interface{}, error)
}
