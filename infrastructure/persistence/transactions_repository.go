package persistence

import (
	"altar-app/domain/entity"
	"altar-app/domain/repository"
	"errors"
	"github.com/jinzhu/gorm"
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
