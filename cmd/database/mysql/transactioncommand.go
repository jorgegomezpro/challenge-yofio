package mysql

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"main.go/contracts"
	"main.go/models"
)

type TransactionCommand struct {
	connectionString string
	tableName        string
}

func NewTransactionDal() contracts.TransactionDal {
	// Format to connection string -> "root:password1@tcp(127.0.0.1:3306)/test"
	return &TransactionCommand{
		os.Getenv("CONNECTION_STRING"),
		"transaction_yf",
	}
}

func (t *TransactionCommand) AddTransaction(successful bool) error {
	db, err := sql.Open("mysql", t.connectionString)
	if err != nil {
		return err
	}

	defer db.Close()

	insert, err := db.Query(fmt.Sprintf("INSERT INTO %v(success) VALUES (?)", t.tableName), successful)
	if err != nil {
		return err
	}

	defer insert.Close()
	return nil
}
func (t *TransactionCommand) GetStatistics() (error, models.Transactions) {
	transactions := models.Transactions{}
	db, err := sql.Open("mysql", t.connectionString)
	if err != nil {
		return err, transactions
	}

	defer db.Close()

	// Execute the query
	results, err := db.Query(fmt.Sprintf(
		`select
		*
		,(succesfull/total) as SuccessfulAverage
		,(unsuccesfull/total) as UnSuccessfulAverage
		from(
	select
		count(id) as total,
		count(case when success=1 then 1 else null end) as succesfull,
		count(case when success=0 then 1 else null end) as unsuccesfull
	from precioyviajes.transaction_yf
	)q`,
		t.tableName))

	if err != nil {
		return err, transactions
	}

	for results.Next() {
		err = results.Scan(
			&transactions.Total,
			&transactions.Successful,
			&transactions.UnSuccessful,
			&transactions.SuccessfulAverage,
			&transactions.UnSuccessfulAverage,
		)
		if err != nil {
			return err, transactions
		}

		return nil, transactions
	}

	return nil, transactions
}
