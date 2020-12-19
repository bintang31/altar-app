package repository

import (
	"altar-app/domain/entity"
)

//TransactionRepository : Transaction collection of methods that the infrastructure
type TransactionRepository interface {
	GetTransactions() ([]entity.Transaction, error)
	SaveTransactions(*entity.Transaction) (*entity.Transaction, map[string]string)
	SaveTransactionsKolektif(*entity.TransactionsKolektif) (*entity.TransactionsKolektif, map[string]string)
}
