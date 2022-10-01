package domain

import (
	"time"

	"github.com/jorgeluizjava/banking/app/dto"
	"github.com/jorgeluizjava/banking/app/errs"
)

const dbTSLayout = "2006-01-02 15:04:05"

type Account struct {
	AccountId   string  `db:"account_id"`
	CustomerId  string  `db:"customer_id"`
	OpeningDate string  `db:"opening_date"`
	AccountType string  `db:"account_type"`
	Amount      float64 `db:"amount"`
	Status      string  `db:"status"`
}

type AccountRepository interface {
	Save(account Account) (*Account, *errs.AppError)
}

func NewAccount(customerId string, accountType string, amount float64) Account {
	now := time.Now()
	return Account{
		CustomerId:  customerId,
		AccountType: accountType,
		Amount:      amount,
		OpeningDate: now.Format(dbTSLayout),
		Status:      "1",
	}
}

func (a Account) ToNewAccountResponseDto() *dto.NewAccountResponse {
	return &dto.NewAccountResponse{
		AccountId: a.AccountId,
	}
}
