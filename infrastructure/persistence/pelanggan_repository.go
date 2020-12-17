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

//InquiryLoketTagihanAirByNosamb : Inquiry Get Data Tagihan Nonair by Nosamb
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
