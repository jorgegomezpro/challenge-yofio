package contracts

import "main.go/models"

type TransactionDal interface {
	AddTransaction(successful bool) error
	GetStatistics() (error, models.Transactions)
}
