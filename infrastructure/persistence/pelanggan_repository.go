package persistence

import (
	"altar-app/application/config"
	"altar-app/domain/entity"
	"altar-app/domain/repository"
	"bytes"
	"encoding/json"
	"errors"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

//PelangganRepo : Call DB
type PelangganRepo struct {
	db *gorm.DB
}

//NewPelangganRepository : Pelanggan Repository
func NewPelangganRepository(db *gorm.DB) *PelangganRepo {
	return &PelangganRepo{db}
}

//PelangganRepo implements the repository.PelangganRepo interface
var _ repository.PelangganRepository = &PelangganRepo{}

//GetPelanggans : Get All Data Pelanggan
func (r *PelangganRepo) GetPelanggans() ([]entity.Pelanggan, error) {
	var pelanggans []entity.Pelanggan
	err := r.db.Debug().Find(&pelanggans).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("pelanggan not found")
	}
	return pelanggans, nil
}

//GetTagihanAirPelanggansByNosamb : Get Data Tagihan Air by Nosamb
func (r *PelangganRepo) GetTagihanAirPelanggansByNosamb(nosamb string) ([]entity.Drd, error) {
	var tagihanair []entity.Drd
	err := r.db.Debug().Where("nosamb = ? AND lunas = ?", nosamb, "0").Find(&tagihanair).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("pelanggan not found")
	}
	return tagihanair, nil
}

//GetTagihanNonAirPelanggansByNosamb : Get Data Tagihan Nonair by Nosamb
func (r *PelangganRepo) GetTagihanNonAirPelanggansByNosamb(u *entity.PeriodeNonair) ([]entity.Nonair, error) {
	var tagihannonair []entity.Nonair
	err := r.db.Debug().Table("nonairs").Select("nonairs.periode,nonairs_jenis.name as jenis_tagihan,nonairs.administrasi,nonairs.biayapasang as biayapasang,"+
		"nonairs.total,nonairs.denda_tunggakan,nonairs.lainnya").Joins("left join nonairs_jenis "+
		"ON nonairs_jenis.code = nonairs.jenis").Joins("left join angsuran ON angsuran.nomor = nonairs.nomor").Where("nonairs.nomor = ? "+
		"and nonairs.lunas = ? and nonairs.angsur = ?", u.Nosamb, "0", "0").Order("nonairs.periode asc,nonairs.termin asc").Find(&tagihannonair).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("pelanggan not found")
	}
	return tagihannonair, nil
}

//GetTagihanNonAirPelanggansByPeriode : Get Data Tagihan Nonair by Periode
func (r *PelangganRepo) GetTagihanNonAirPelanggansByPeriode(nosamb string) ([]map[string]interface{}, error) {
	var periodenonair []entity.PeriodeNonair
	var nonair []entity.Nonair
	var angsuran []entity.AngsuranNonAir
	err := r.db.Debug().Table("nonairs").Select("periode,bulan,sum(total) as total_per_periode,"+
		"'BELUM TERBAYAR' as status").Where("lunas = ?", "0").Group("periode,bulan").Order("periode desc").Where("nomor = ?", nosamb).Find(&periodenonair).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("pelanggan not found")
	}
	list := make([]map[string]interface{}, len(periodenonair))

	x := 0
	for _, p := range periodenonair {

		//Non Angsuran Berkala
		err := r.db.Table("nonairs").Select("nonairs.periode,nonairs_jenis.name as jenis_tagihan,nonairs.administrasi,nonairs.biayapasang as biayapasang,"+
			"nonairs.total,nonairs.denda_tunggakan,nonairs.lainnya").Joins("left join nonairs_jenis "+
			"ON nonairs_jenis.code = nonairs.jenis").Joins("left join angsuran ON angsuran.nomor = nonairs.nomor").Where("nonairs.nomor = ? "+
			"and nonairs.lunas = ? and nonairs.periode = ? and nonairs.angsur = ?", nosamb, "0", p.Periode, "0").Order("nonairs.periode asc,nonairs.termin asc").Find(&nonair).Error

		//Angsuran
		err = r.db.Table("angsuran").Select("angsuran.nomor,nonairs_jenis.name as jenis,nonairs.administrasi,nonairs.denda_tunggakan,angsuran."+
			"jumlahangsuranpokok as total_tagihan,nonairs.total,"+
			"concat(nonairs.termin,' dari ' ,angsuran.jumlah_termin) as jumlah_termin,concat('DRD',' ', nonairs.bulan) as keterangan,"+
			"penagihans_angsuran.sisa_tagihan").Joins("join nonairs ON nonairs.nomor = angsuran.nomor join nonairs_jenis "+
			"ON nonairs.jenis = nonairs_jenis.code join penagihans_angsuran ON penagihans_angsuran.noangsuran = nonairs.no_angsuran").Where(""+
			"angsuran.nomor = ? and nonairs.periode = ? and nonairs.angsur = ?", nosamb, p.Periode, "1").Order("nonairs.periode asc,nonairs.termin asc").Find(&angsuran).Error

		if err != nil {
			return nil, err
		}
		if gorm.IsRecordNotFoundError(err) {
			return nil, errors.New("pelanggan not found")
		}

		list[x] = map[string]interface{}{
			"periode":           p.Bulan,
			"nonair":            nonair,
			"angsuran":          angsuran,
			"total_per_periode": p.TotalPerPeriode,
			"status":            p.Status,
		}
		x = x + 1
	}

	return list, nil
}

//InquiryLoketTagihanAirByNosamb : Inquiry Get Data Tagihan Air by Nosamb
func (r *PelangganRepo) InquiryLoketTagihanAirByNosamb(u *entity.InputInquiryPelanggan) ([]entity.RekairDetail, error) {
	var rekairdetail []entity.RekairDetail

	conf := config.LoadModuleLoketConfig("moduleloket")
	values := map[string]string{"pdamcode": u.Pdam, "nosamb": u.Nosamb, "userloket": conf.UserLoket}
	jsonValue, _ := json.Marshal(values)

	client := &http.Client{}
	req, err := http.NewRequest("GET", conf.EndPoint+"bill/inquery", bytes.NewBuffer(jsonValue))
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(conf.User, conf.Password)
	resp, err := client.Do(req)
	bodyText, err := ioutil.ReadAll(resp.Body)

	var inquiryData entity.InquiryData
	json.Unmarshal(bodyText, &inquiryData)

	bodyBytes, _ := json.Marshal(inquiryData.Data)

	var collect entity.InquiryCollection
	json.Unmarshal(bodyBytes, &collect)

	bodyBytes2, _ := json.Marshal(collect.Rekair)

	_ = json.Unmarshal([]byte(bodyBytes2), &rekairdetail)

	return rekairdetail, nil
}

//UpdateDrdByNosamb : Update Data DRD
func (r *PelangganRepo) UpdateDrdByNosamb(rd *entity.Drd) (*entity.Drd, map[string]string) {
	var drds entity.Drd
	currentTimeBayar := time.Now().Format("2006-01-02 15:04:05")
	dbErr := map[string]string{}
	if rd.TransactionsID > 0 {
		err := r.db.Debug().Model(&drds).Where("nosamb = ? ", rd.Nosamb).Updates(map[string]interface{}{
			"lunas":           rd.Lunas,
			"transactions_id": rd.TransactionsID,
			"updated_at":      currentTimeBayar,
			"tglbayar":        currentTimeBayar,
		}).Error
		if err != nil {
			//If the email is already taken
			if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
				dbErr["email_taken"] = "email already taken"
				return nil, dbErr
			}
			//any other db error
			dbErr["db_error"] = "database error"
			return nil, dbErr
		}
	} else {
		err := r.db.Debug().Model(&drds).Where("nosamb = ? and periode = ?", rd.Nosamb, rd.Periode).Updates(map[string]interface{}{
			"total": rd.Total,
		}).Error
		if err != nil {
			//If the email is already taken
			if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
				dbErr["email_taken"] = "email already taken"
				return nil, dbErr
			}
			//any other db error
			dbErr["db_error"] = "database error"
			return nil, dbErr
		}
	}

	return &drds, nil
}

//UpdateNonAirByNosamb : Update Data NonAir
func (r *PelangganRepo) UpdateNonAirByNosamb(rd *entity.Nonair) (*entity.Nonair, map[string]string) {
	var nonairs entity.Nonair
	currentTimeBayar := time.Now().Format("2006-01-02 15:04:05")
	dbErr := map[string]string{}
	if rd.TransactionsID > 0 {
		err := r.db.Debug().Model(&nonairs).Where("nomor = ? and lunas = ?", rd.Nomor, "0").Updates(map[string]interface{}{
			"lunas":           rd.Lunas,
			"transactions_id": rd.TransactionsID,
			"updated_at":      currentTimeBayar,
			"waktu_bayar":     currentTimeBayar,
		}).Error
		if err != nil {
			//If the email is already taken
			if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
				dbErr["email_taken"] = "email already taken"
				return nil, dbErr
			}
			//any other db error
			dbErr["db_error"] = "database error"
			return nil, dbErr
		}
	}

	return &nonairs, nil
}

//InquiryLoketTagihanNonAirByNosamb : Inquiry Get Data Tagihan Nonair by Nosamb
func (r *PelangganRepo) InquiryLoketTagihanNonAirByNosamb(u *entity.InputInquiryPelanggan) ([]entity.NonAirDetail, error) {
	var nonairdetail []entity.NonAirDetail

	conf := config.LoadModuleLoketConfig("moduleloket")
	values := map[string]string{"pdamcode": u.Pdam, "nosamb": u.Nosamb, "userloket": conf.UserLoket}
	jsonValue, _ := json.Marshal(values)

	client := &http.Client{}
	req, err := http.NewRequest("GET", conf.EndPoint+"bill/inquery", bytes.NewBuffer(jsonValue))
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(conf.User, conf.Password)
	resp, err := client.Do(req)
	bodyText, err := ioutil.ReadAll(resp.Body)

	var inquiryData entity.InquiryData
	json.Unmarshal(bodyText, &inquiryData)

	bodyBytes, _ := json.Marshal(inquiryData.Data)

	var collect entity.InquiryCollection
	json.Unmarshal(bodyBytes, &collect)

	bodyBytes2, _ := json.Marshal(collect.Nonair)

	_ = json.Unmarshal([]byte(bodyBytes2), &nonairdetail)

	return nonairdetail, nil
}

//InquiryLoketAngsuranByNosamb : Inquiry Get Data Tagihan Angsuran by Nosamb
func (r *PelangganRepo) InquiryLoketAngsuranByNosamb(u *entity.InputInquiryPelanggan) ([]entity.AngsuranDetail, error) {
	var angsuran []entity.AngsuranDetail

	conf := config.LoadModuleLoketConfig("moduleloket")
	values := map[string]string{"pdamcode": u.Pdam, "nosamb": u.Nosamb, "userloket": conf.UserLoket}
	jsonValue, _ := json.Marshal(values)

	client := &http.Client{}
	req, err := http.NewRequest("GET", conf.EndPoint+"bill/inquery", bytes.NewBuffer(jsonValue))
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(conf.User, conf.Password)
	resp, err := client.Do(req)
	bodyText, err := ioutil.ReadAll(resp.Body)

	var inquiryData entity.InquiryData
	json.Unmarshal(bodyText, &inquiryData)

	bodyBytes, _ := json.Marshal(inquiryData.Data)

	var collect entity.InquiryCollection
	json.Unmarshal(bodyBytes, &collect)

	bodyBytes2, _ := json.Marshal(collect.Angsuran)

	_ = json.Unmarshal([]byte(bodyBytes2), &angsuran)

	return angsuran, nil
}

//InsertAngsuranByNosamb : Insert Tagihan Angsuran
func (r *PelangganRepo) InsertAngsuranByNosamb(rd *entity.Nonair) (*entity.Nonair, map[string]string) {
	var nonairs entity.Nonair
	dbErr := map[string]string{}
	var count int64
	err := r.db.Debug().First(&nonairs).Where("urutan = ?", rd.Urutan).Count(&count).Error
	if err != nil {
		return nil, dbErr
	}
	if count < 1 {
		rd.UpdatedAt = time.Now()
		err := r.db.Debug().Create(&rd).Error
		if err != nil {
			//If the email is already taken
			if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
				dbErr["email_taken"] = "email already taken"
				return nil, dbErr
			}
			//any other db error
			dbErr["db_error"] = "database error"
			return nil, dbErr
		}
	}

	return &nonairs, nil
}

//GetRiwayatTagihanByNosamb : Get Data Riwayat Tagihan by Nosamb
func (r *PelangganRepo) GetRiwayatTagihanByNosamb(nosamb string) ([]map[string]interface{}, error) {
	var perioderiwayat []entity.PeriodeRiwayat
	var tagihanair []entity.Drd
	var nonair []entity.Nonair
	var angsuran []entity.AngsuranNonAir
	err := r.db.Debug().Table("riwayat_billing").Select("periode,bulan,status,sum(total) as total_per_periode,"+
		"status as status,sum(pakai) as total_pemakaian").Where("nosamb = ?", nosamb).Group("periode,bulan,status").Order("periode desc").Find(&perioderiwayat).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("pelanggan not found")
	}
	list := make([]map[string]interface{}, len(perioderiwayat))

	x := 0
	for _, p := range perioderiwayat {

		//Tagihan Air
		err := r.db.Debug().Where("nosamb = ? AND periode = ? AND lunas = ?", nosamb, p.Periode, "1").Find(&tagihanair).Error

		//Tagihan Non Air
		err = r.db.Table("nonairs").Select("nonairs.periode,nonairs_jenis.name as jenis_tagihan,nonairs.administrasi,nonairs.biayapasang as biayapasang,"+
			"nonairs.total,nonairs.denda_tunggakan,nonairs.lainnya").Joins("left join nonairs_jenis "+
			"ON nonairs_jenis.code = nonairs.jenis").Joins("left join angsuran ON angsuran.nomor = nonairs.nomor").Where("nonairs.nomor = ? "+
			"and nonairs.lunas = ? and nonairs.periode = ? and nonairs.angsur = ?", nosamb, "1", p.Periode, "0").Order("nonairs.periode asc,nonairs.termin asc").Find(&nonair).Error

		//Angsuran
		err = r.db.Table("angsuran").Select("angsuran.nomor,nonairs_jenis.name as jenis,nonairs.administrasi,nonairs.denda_tunggakan,angsuran."+
			"jumlahangsuranpokok as total_tagihan,nonairs.total,"+
			"concat(nonairs.termin,' dari ' ,angsuran.jumlah_termin) as jumlah_termin,concat('DRD',' ', nonairs.bulan) as keterangan,"+
			"penagihans_angsuran.sisa_tagihan").Joins("join nonairs ON nonairs.nomor = angsuran.nomor join nonairs_jenis "+
			"ON nonairs.jenis = nonairs_jenis.code join penagihans_angsuran ON penagihans_angsuran.noangsuran = nonairs.no_angsuran").Where(""+
			"angsuran.nomor = ? and nonairs.periode = ? and nonairs.angsur = ?", nosamb, p.Periode, "1").Order("nonairs.periode asc,nonairs.termin asc").Find(&angsuran).Error

		if err != nil {
			return nil, err
		}
		if gorm.IsRecordNotFoundError(err) {
			return nil, errors.New("pelanggan not found")
		}

		list[x] = map[string]interface{}{
			"periode":           p.Bulan,
			"air":               tagihanair,
			"nonair":            nonair,
			"angsuran":          angsuran,
			"total_per_periode": p.TotalPerPeriode,
			"total_pemakaian":   p.TotalPemakaian,
			"status":            p.Status,
		}
		x = x + 1
	}

	return list, nil
}
