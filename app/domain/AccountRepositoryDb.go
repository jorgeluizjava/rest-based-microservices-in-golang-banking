package domain

import (
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/jorgeluizjava/banking/app/errs"
	"github.com/jorgeluizjava/banking/app/logger"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func NewAccountRepositoryDb(client *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{client: client}
}

func (d AccountRepositoryDb) Save(account Account) (*Account, *errs.AppError) {
	sqlInsert := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) values (?, ?, ?, ?, ?)"
	result, err := d.client.Exec(sqlInsert, account.CustomerId, account.OpeningDate, account.AccountType, account.Amount, account.Status)
	if err != nil {
		logger.Error("Error while creating new account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}
	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last insert id fromaccount: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}
	account.AccountId = strconv.FormatInt(id, 10)
	return &account, nil
}
