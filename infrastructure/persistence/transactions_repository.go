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
