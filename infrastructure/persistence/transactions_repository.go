package persistence

import (
	"altar-app/domain/entity"
	"altar-app/domain/repository"
	"errors"
	"github.com/jinzhu/gorm"
	"strings"
)

//TransactionRepo : Call DB
type TransactionRepo struct {
	db *gorm.DB
}

//NewTransactionRepository : Transaction Repository
func NewTransactionRepository(db *gorm.DB) *TransactionRepo {
	return &TransactionRepo{db}
}

//PelangganRepo implements the repository.PelangganRepo interface
var _ repository.TransactionRepository = &TransactionRepo{}

//GetTransactions : Get All Data Transaksi
func (r *TransactionRepo) GetTransactions() ([]entity.Transaction, error) {
	var transactions []entity.Transaction
	err := r.db.Debug().Find(&transactions).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("transaction not found")
	}
	return transactions, nil
}

//GetTransactionsByID : Get Data Transaksi By ID
func (r *TransactionRepo) GetTransactionsByID(id int) (*entity.TransactionPelanggan, map[string]string) {
	var transactions entity.TransactionPelanggan
	dbErr := map[string]string{}
	err := r.db.Debug().Table("transactions").Select("id,transactions.pelanggan as nosamb, transactions.notes as notes,penagihans.nama,penagihans.alamat,penagihans.pdam,penagihans.rayon_name,penagihans.golongan,penagihans.notelp,penagihans.kode_pdam,"+
		"penagihans.status_kolektif,penagihans.status_pelanggan,transactions.total as total_tagihan,transactions.periode_billing as periode_billing,TO_CHAR(transactions.created_at::DATE, 'dd Mon yyyy') as tgl_bayar,penagihans.status_billing").Joins("left join penagihans ON penagihans.nosamb = "+
		"transactions.pelanggan").Where("transactions.id = ? ", id).Find(&transactions).Where("id = ?", id).Error

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
	return &transactions, nil
}

//SaveTransactions : Save Data Transaksi
func (r *TransactionRepo) SaveTransactions(trx *entity.Transaction) (*entity.Transaction, map[string]string) {
	var transactions entity.Transaction
	dbErr := map[string]string{}
	err := r.db.Debug().Create(&trx).Error
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
	return &transactions, nil
}

//SaveTransactionsKolektif : Save Data Transaksi Kolektif
func (r *TransactionRepo) SaveTransactionsKolektif(trx *entity.TransactionsKolektif) (*entity.TransactionsKolektif, map[string]string) {
	var transactions entity.TransactionsKolektif
	dbErr := map[string]string{}
	err := r.db.Debug().Create(&trx).Error
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
	return &transactions, nil
}

//GetDetailTransactionsByID : Get Detail Rincian Transactions
func (r *TransactionRepo) GetDetailTransactionsByID(id int) ([]map[string]interface{}, error) {
	var perioderiwayat []entity.RiwayatTransactionPelanggan
	var tagihanair []entity.Drd
	var nonair []entity.Nonair
	var angsuran []entity.AngsuranNonAir
	err := r.db.Debug().Table("riwayat_billing").Select("periode,bulan,status,sum(total) as total_per_periode,"+
		"status as status,sum(pakai) as total_pemakaian").Where("transactions_id = ?", id).Group("periode,bulan,status").Order("periode desc").Find(&perioderiwayat).Error
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
		err := r.db.Debug().Where("transactions_id = ? AND periode = ? AND lunas = ?", id, p.Periode, "1").Find(&tagihanair).Error

		//Tagihan Non Air
		err = r.db.Table("nonairs").Select("nonairs.periode,nonairs_jenis.name as jenis_tagihan,nonairs.administrasi,nonairs.biayapasang as biayapasang,"+
			"nonairs.total,nonairs.denda_tunggakan,nonairs.lainnya").Joins("left join nonairs_jenis "+
			"ON nonairs_jenis.code = nonairs.jenis").Joins("left join angsuran ON angsuran.nomor = nonairs.nomor").Where("nonairs.transactions_id = ? "+
			"and nonairs.lunas = ? and nonairs.periode = ? and nonairs.angsur = ?", id, "1", p.Periode, "0").Order("nonairs.periode asc,nonairs.termin asc").Find(&nonair).Error

		//Angsuran
		err = r.db.Table("angsuran").Select("angsuran.nomor,nonairs_jenis.name as jenis,nonairs.administrasi,nonairs.denda_tunggakan,angsuran."+
			"jumlahangsuranpokok as total_tagihan,nonairs.total,"+
			"concat(nonairs.termin,' dari ' ,angsuran.jumlah_termin) as jumlah_termin,concat('DRD',' ', nonairs.bulan) as keterangan,"+
			"penagihans_angsuran.sisa_tagihan").Joins("join nonairs ON nonairs.nomor = angsuran.nomor join nonairs_jenis "+
			"ON nonairs.jenis = nonairs_jenis.code join penagihans_angsuran ON penagihans_angsuran.noangsuran = nonairs.no_angsuran").Where(""+
			"nonairs.transactions_id = ? and nonairs.periode = ? and nonairs.angsur = ?", id, p.Periode, "1").Order("nonairs.periode asc,nonairs.termin asc").Find(&angsuran).Error

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
