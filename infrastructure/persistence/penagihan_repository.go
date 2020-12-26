package persistence

import (
	"altar-app/application/config"
	"altar-app/domain/entity"
	"altar-app/domain/repository"
	"encoding/json"
	"errors"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	//BELUMTERBAYAR : status
	BELUMTERBAYAR = "BELUM TERBAYAR"
)

//PenagihanRepo : Call DB
type PenagihanRepo struct {
	db *gorm.DB
}

//NewPenagihanRepository : Pelanggan Repository
func NewPenagihanRepository(db *gorm.DB) *PenagihanRepo {
	return &PenagihanRepo{db}
}

//PenagihanRepo implements the repository.PelangganRepo interface
var _ repository.PelangganRepository = &PelangganRepo{}

//GetPenagihans : Get All Data Penagihans
func (r *PenagihanRepo) GetPenagihans() ([]entity.Penagihan, error) {
	var penagihans []entity.Penagihan
	err := r.db.Debug().Find(&penagihans).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("pelanggan not found")
	}
	return penagihans, nil
}

//GetPenagihansByUserPDAM : Penagihan By User Login PDAM
func (r *PenagihanRepo) GetPenagihansByUserPDAM(id uint64) ([]entity.PenagihansSrKolektif, error) {
	var penagihans []entity.PenagihansSrKolektif

	err := r.db.Debug().Table("petugas_rayons").Select("penagihans_sr_kolektifs.nosamb, penagihans_sr_kolektifs.nama,penagihans_sr_kolektifs.pelanggan,penagihans_sr_kolektifs.notelp,penagihans_sr_kolektifs.golongan,"+
		"penagihans_sr_kolektifs.kode_pdam,penagihans_sr_kolektifs.alamat,penagihans_sr_kolektifs.pdam,penagihans_sr_kolektifs.rayon_name,"+
		"penagihans_sr_kolektifs.status_kolektif,penagihans_sr_kolektifs.status_pelanggan,penagihans_sr_kolektifs.tagihan_air,penagihans_sr_kolektifs.total_tagihan_air,penagihans_sr_kolektifs.tagihan_nonair,"+
		"penagihans_sr_kolektifs.total_tagihan_nonair,penagihans_sr_kolektifs.total_tagihan,penagihans_sr_kolektifs.status_billing,penagihans_sr_kolektifs.periode_tagihan").Joins("join penagihans_sr_kolektifs ON penagihans_sr_kolektifs.kode_rayon = "+
		"petugas_rayons.rayon").Where("petugas_rayons.petugas = ? AND penagihans_sr_kolektifs.status_billing = ?", id, "BELUM TERBAYAR").Order("penagihans_sr_kolektifs.total_tagihan desc").Find(&penagihans).Error

	//err := r.db.Debug().Find(&penagihans).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("penagihan not found")
	}
	return penagihans, nil
}

//GetPenagihanByNosamb : Get Penagihan by Nosamb
func (r *PenagihanRepo) GetPenagihanByNosamb(nosamb string) (*entity.Penagihan, error) {
	var penagihan entity.Penagihan
	err := r.db.Debug().Where("nosamb = ?", nosamb).Find(&penagihan).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("pelanggan not found")
	}
	return &penagihan, nil
}

//BayarTagihanByNosamb : Bayar Tagihan by Nosamb
func (r *PenagihanRepo) BayarTagihanByNosamb(u *entity.Bayar) (*entity.ResponseLoket, map[string]string) {
	var penagihans entity.PenagihansSrKolektif
	var responseloket entity.ResponseLoket
	timeBayar := time.Now().Format("2006-01-02 15:04:05")

	dbErr := map[string]string{}
	err := r.db.Debug().Where("nosamb = ?", u.Nosamb).Find(&penagihans).Error
	if err != nil {
		return nil, dbErr
	}

	conf := config.LoadModuleLoketConfig("moduleloket")
	formloket := url.Values{}
	formloket.Add("nosamb", u.Nosamb)
	formloket.Add("pdamcode", u.Pdam)
	formloket.Add("userloket", u.UserLoket)
	formloket.Add("totaltagihan", strconv.FormatFloat(penagihans.TotalTagihan, 'f', 2, 64))
	formloket.Add("tglbayar", timeBayar)

	client := &http.Client{}
	req, err := http.NewRequest("POST", conf.EndPoint+"payment", strings.NewReader(formloket.Encode()))
	req.SetBasicAuth(conf.User, conf.Password)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		return nil, dbErr
	}
	bodyText, err := ioutil.ReadAll(resp.Body)

	var inquiryData entity.InquiryData
	json.Unmarshal(bodyText, &inquiryData)

	bodyBytes2, _ := json.Marshal(inquiryData.Response)

	_ = json.Unmarshal([]byte(bodyBytes2), &responseloket)

	return &responseloket, nil
}

//GetPenagihanByParam : Get Penagihan by Param
func (r *PenagihanRepo) GetPenagihanByParam(t *entity.PenagihansParams) ([]entity.PenagihansSrKolektif, error) {
	var penagihan []entity.PenagihansSrKolektif
	var err error
	switch t.Filter {
	case "SR":
		err = r.db.Debug().Table("petugas_rayons").Select("penagihans_sr_kolektifs.nosamb, penagihans_sr_kolektifs.nama,penagihans_sr_kolektifs.pelanggan,penagihans_sr_kolektifs.notelp,penagihans_sr_kolektifs.golongan,penagihans_sr_kolektifs.kode_pdam,penagihans_sr_kolektifs.alamat,penagihans_sr_kolektifs.pdam,penagihans_sr_kolektifs.rayon_name,"+
			"penagihans_sr_kolektifs.status_kolektif,penagihans_sr_kolektifs.status_pelanggan,penagihans_sr_kolektifs.tagihan_air,penagihans_sr_kolektifs.total_tagihan_air,penagihans_sr_kolektifs.tagihan_nonair,"+
			"penagihans_sr_kolektifs.total_tagihan_nonair,penagihans_sr_kolektifs.total_tagihan,penagihans_sr_kolektifs.status_billing,penagihans_sr_kolektifs.periode_tagihan").Joins("join penagihans_sr_kolektifs ON penagihans_sr_kolektifs.kode_rayon = "+
			"petugas_rayons.rayon").Where("petugas_rayons.petugas = ? AND penagihans_sr_kolektifs.status_billing = ? AND penagihans_sr_kolektifs.status_kolektif = ?", t.UserID, "BELUM TERBAYAR", "sambungan_rumah").Order("penagihans_sr_kolektifs.total_tagihan desc").Find(&penagihan).Error

	default:
		err = r.db.Debug().Table("petugas_rayons").Select("penagihans_sr_kolektifs.nosamb, penagihans_sr_kolektifs.nama,penagihans_sr_kolektifs.pelanggan,penagihans_sr_kolektifs.notelp,penagihans_sr_kolektifs.golongan,penagihans_sr_kolektifs.kode_pdam,penagihans_sr_kolektifs.alamat,penagihans_sr_kolektifs.pdam,penagihans_sr_kolektifs.rayon_name,"+
			"penagihans_sr_kolektifs.status_kolektif,penagihans_sr_kolektifs.status_pelanggan,penagihans_sr_kolektifs.tagihan_air,penagihans_sr_kolektifs.total_tagihan_air,penagihans_sr_kolektifs.tagihan_nonair,"+
			"penagihans_sr_kolektifs.total_tagihan_nonair,penagihans_sr_kolektifs.total_tagihan,penagihans_sr_kolektifs.status_billing,penagihans_sr_kolektifs.periode_tagihan").Joins("join penagihans_sr_kolektifs ON penagihans_sr_kolektifs.kode_rayon = "+
			"petugas_rayons.rayon").Where("petugas_rayons.petugas = ? AND penagihans_sr_kolektifs.status_billing = ?", t.UserID, "BELUM TERBAYAR").Order("penagihans_sr_kolektifs.total_tagihan desc").Find(&penagihan).Error

	}

	config.Paging(&config.Param{
		DB:      r.db,
		Page:    t.Page,
		Limit:   5,
		OrderBy: []string{"penagihans_sr_kolektifs.total_tagihan desc"},
		ShowSQL: true,
	}, &penagihan)

	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("pelanggan not found")
	}
	return penagihan, nil
}

//GetPenagihansInquiryProcess : Get All Data Penagihans for Inquiry Process
func (r *PenagihanRepo) GetPenagihansInquiryProcess() ([]entity.PenagihansSrKolektif, error) {
	var penagihans []entity.PenagihansSrKolektif
	today := time.Now()
	backFiveHours := time.Hour * time.Duration(-600)
	past := today.Add(backFiveHours)

	backBillUpload := time.Hour * time.Duration(-120)
	pastbillupload := today.Add(backBillUpload)

	err := r.db.Debug().Table("petugas_rayons").Select("penagihans.nosamb, penagihans.nama,penagihans.kode_pdam,penagihans.alamat,penagihans.pdam,penagihans.rayon_name,"+
		"penagihans.status_kolektif,penagihans.status_pelanggan,penagihans.tagihan_air,penagihans.total_tagihan_air,penagihans.tagihan_nonair,"+
		"penagihans.total_tagihan_nonair,penagihans.total_tagihan").Joins("join penagihans ON penagihans.kode_rayon = "+
		"petugas_rayons.rayon").Where("penagihans.kode_pdam = ? AND (date(penagihans.last_inquiry) = ? OR date(penagihans.last_bill_upload) = ?) "+
		"AND penagihans.status_billing = ?", "MJI", past.Format("2006-01-02"), pastbillupload.Format("2006-01-02"), BELUMTERBAYAR).Find(&penagihans).Error

	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("pelanggan not found")
	}
	return penagihans, nil
}
