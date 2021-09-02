package main

import (
	"main.go/contracts"
	"main.go/database/mysql"
	"main.go/services"
)

type Services struct {
	creditAssigner contracts.CreditAssigner
	transactionDal contracts.TransactionDal
}

func GetDependencies(creditTypes ...int32) Services {
	return Services{
		creditAssigner: services.NewRecursiveAssign(creditTypes...),
		transactionDal: mysql.NewTransactionDal(),
	}
}
