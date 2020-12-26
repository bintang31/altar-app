package persistence

import (
	"altar-app/domain/entity"
	"altar-app/domain/repository"
	"bufio"
	"encoding/base64"
	"errors"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"os"
)

//PetugasRepo : Call DB
type PetugasRepo struct {
	db *gorm.DB
}

//NewPetugasRepository : Petugas Repository
func NewPetugasRepository(db *gorm.DB) *PetugasRepo {
	return &PetugasRepo{db}
}

//PelangganRepo implements the repository.PelangganRepo interface
var _ repository.PetugasRepository = &PetugasRepo{}

//GetPetugas : Get All Data Petugas
func (p *PetugasRepo) GetPetugas() ([]entity.Petugas, error) {
	var petugas []entity.Petugas
	err := p.db.Debug().Find(&petugas).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("petugas not found")
	}
	return petugas, nil
}

//GetProfilePetugas : Get Data Profile Petugas
func (p *PetugasRepo) GetProfilePetugas(id uint64) (*entity.Petugas, error) {
	var petugas entity.Petugas
	err := p.db.Debug().Select("id,nama_petugas,pdam,kode_pdam,alamat_pdam,notelp_pdam,photo_uri,photo_base64,online,transaction,total_transaction,lembar_tagihan_diterima,total_diterima,sambungan_rumah_diterima,kolektif_diterima,"+
		"sambungan_rumah_setor,kolektif_setor,setoran,total_setoran,sisa_limit,limit_petugas,FLOOR(EXTRACT(epoch FROM created))::int as created,TO_CHAR(created, 'Mon yyyy') as tgl_gabung").Where("id = ?", id).Find(&petugas).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("petugas not found")
	}
	// Open file on disk.
	f, _ := os.Open(petugas.PhotoBase64)

	// Read entire JPG into byte slice.
	reader := bufio.NewReader(f)
	content, _ := ioutil.ReadAll(reader)

	// Encode as base64.
	encoded := base64.StdEncoding.EncodeToString(content)
	petugas.PhotoBase64 = encoded
	return &petugas, nil
}

//GetTagihanAirByPetugas : Get Data Tagihan Air per Petugas
func (p *PetugasRepo) GetTagihanAirByPetugas(id uint64) ([]entity.DrdbyPetugas, error) {
	var tagihanair []entity.DrdbyPetugas
	err := p.db.Debug().Table("petugas_rayons").Select("penagihans_sr_kolektifs.nosamb,"+
		"drds.periode,drds.pdam,drds.pakai,drds.rekair, 0 as danameter,0 as airlimbah,drds.administrasi,drds.pemeliharaan,drds.dendatunggakan,drds.bulan,"+
		"drds.angsur,drds.batal,drds.lunas,drds.total,floor(extract(epoch FROM drds.tglupload))::int as tglupload,floor(extract(epoch FROM drds.tglbayar))::int as tglbayar").Joins("join penagihans_sr_kolektifs ON penagihans_sr_kolektifs.kode_rayon = "+
		"petugas_rayons.rayon join drds ON (drds.nosamb = penagihans_sr_kolektifs.nosamb and drds.pdam = penagihans_sr_kolektifs.kode_pdam)").Where("petugas_rayons.petugas = ? AND penagihans_sr_kolektifs.status_billing = ?", id, "BELUM TERBAYAR").Order("drds.periode asc").Find(&tagihanair).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("Tagihan not found")
	}
	return tagihanair, nil
}

//GetTagihanAirKolektifByPetugas : Get Data Tagihan Air Kolektif per Petugas
func (p *PetugasRepo) GetTagihanAirKolektifByPetugas(id uint64) ([]entity.DrdbyPetugas, error) {
	var tagihanair []entity.DrdbyPetugas
	err := p.db.Debug().Table("petugas_rayons").Select("penagihans.nosamb,"+
		"drds.periode,drds.pdam,drds.pakai,drds.rekair, 0 as danameter,0 as airlimbah,drds.administrasi,drds.pemeliharaan,drds.dendatunggakan,drds.bulan,"+
		"drds.angsur,drds.batal,drds.lunas,drds.total,floor(extract(epoch FROM drds.tglupload))::int as tglupload,floor(extract(epoch FROM drds.tglbayar))::int as tglbayar").Joins("join penagihans ON penagihans.kode_rayon = "+
		"petugas_rayons.rayon join drds ON (drds.nosamb = penagihans.nosamb and drds.pdam = penagihans.kode_pdam)").Where("petugas_rayons.petugas = ? AND penagihans.status_billing = ? AND penagihans.status_kolektif = ? AND drds.lunas = ?", id, "BELUM TERBAYAR", "kolektif", "0").Order("drds.periode asc").Find(&tagihanair).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("Tagihan not found")
	}
	return tagihanair, nil
}

//GetTagihanNonAirByPetugas : Get Data Tagihan Non Air per Petugas
func (p *PetugasRepo) GetTagihanNonAirByPetugas(id uint64) ([]entity.Nonair, error) {
	var tagihannonair []entity.Nonair
	err := p.db.Debug().Table("petugas_rayons").Select("penagihans.nosamb as nomor,"+
		"nonairs.periode,nonairs_jenis.name as jenis_tagihan,"+
		"nonairs.total,nonairs.administrasi,nonairs.biayapasang,nonairs.denda_tunggakan,nonairs.lainnya").Joins("join penagihans ON penagihans.kode_rayon = "+
		"petugas_rayons.rayon join nonairs ON (nonairs.nomor = penagihans.nosamb and nonairs.pdam = "+
		"penagihans.kode_pdam) join nonairs_jenis ON nonairs_jenis.code =nonairs.jenis").Where("petugas_rayons.petugas = ? AND penagihans.status_billing = ?", id, "BELUM TERBAYAR").Order("nonairs.periode asc").Find(&tagihannonair).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("Tagihan not found")
	}
	return tagihannonair, nil
}
