package application

import (
	"altar-app/domain/entity"
	"altar-app/domain/repository"
)

type transactionsApp struct {
	tr repository.TransactionRepository
}

//transactionsApp implements the PelangganAppInterface
var _ TransactionsAppInterface = &transactionsApp{}

//TransactionsAppInterface : App Interface transaksi repo
type TransactionsAppInterface interface {
	GetTransactions() ([]entity.Transaction, error)
	SaveTransactions(*entity.Transaction) (*entity.Transaction, map[string]string)
	SaveTransactionsKolektif(*entity.TransactionsKolektif) (*entity.TransactionsKolektif, map[string]string)
	GetTransactionsByID(int) (*entity.TransactionPelanggan, map[string]string)
	GetDetailTransactionsByID(int) ([]map[string]interface{}, error)
}

func (t *transactionsApp) GetTransactions() ([]entity.Transaction, error) {
	return t.tr.GetTransactions()
}

func (t *transactionsApp) GetTransactionsByID(id int) (*entity.TransactionPelanggan, map[string]string) {
	return t.tr.GetTransactionsByID(id)
}

func (t *transactionsApp) SaveTransactions(trx *entity.Transaction) (*entity.Transaction, map[string]string) {
	return t.tr.SaveTransactions(trx)
}

func (t *transactionsApp) SaveTransactionsKolektif(trx *entity.TransactionsKolektif) (*entity.TransactionsKolektif, map[string]string) {
	return t.tr.SaveTransactionsKolektif(trx)
}

func (t *transactionsApp) GetDetailTransactionsByID(id int) ([]map[string]interface{}, error) {
	return t.tr.GetDetailTransactionsByID(id)
}
