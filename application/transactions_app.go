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
}

func (t *transactionsApp) GetTransactions() ([]entity.Transaction, error) {
	return t.tr.GetTransactions()
}
