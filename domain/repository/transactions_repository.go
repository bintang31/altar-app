package repository

import (
	"altar-app/domain/entity"
)

//TransactionRepository : Transaction collection of methods that the infrastructure
type TransactionRepository interface {
	GetTransactions() ([]entity.Transaction, error)
	SaveTransactions(*entity.Transaction) (*entity.Transaction, map[string]string)
	SaveTransactionsKolektif(*entity.TransactionsKolektif) (*entity.TransactionsKolektif, map[string]string)
	GetTransactionsByID(int) (*entity.TransactionPelanggan, map[string]string)
	GetDetailTransactionsByID(int) ([]map[string]interface{}, error)
}
