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
	err := r.db.Debug().Where("nosamb = ?", nosamb).Find(&tagihanair).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("pelanggan not found")
	}
	return tagihanair, nil
}

//GetTagihanNonAirPelanggansByNosamb : Get Data Tagihan Nonair by Nosamb
func (r *PelangganRepo) GetTagihanNonAirPelanggansByNosamb(nosamb string) ([]entity.Nonair, error) {
	var tagihannonair []entity.Nonair
	err := r.db.Debug().Where("nomor = ?", nosamb).Find(&tagihannonair).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("pelanggan not found")
	}
	return tagihannonair, nil
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
		err := r.db.Debug().Model(&drds).Where("nosamb = ? and periode = ?", rd.Nosamb, rd.Periode).Updates(map[string]interface{}{
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
		err := r.db.Debug().Model(&nonairs).Where("nomor = ? and periode = ?", rd.Nomor, rd.Periode).Updates(map[string]interface{}{
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
		err := r.db.Debug().Model(&nonairs).Where("nomor = ? and periode = ?", rd.Nomor, rd.Periode).Updates(map[string]interface{}{
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
