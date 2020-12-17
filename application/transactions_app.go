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
}

func (t *transactionsApp) GetTransactions() ([]entity.Transaction, error) {
	return t.tr.GetTransactions()
}

func (t *transactionsApp) SaveTransactions(trx *entity.Transaction) (*entity.Transaction, map[string]string) {
	return t.tr.SaveTransactions(trx)
}
