package dto

import (
	"strings"

	"github.com/jorgeluizjava/banking/app/errs"
)

type NewAccountRequest struct {
	CustomerId  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

func (r NewAccountRequest) allowedTypes() bool {
	return strings.ToLower(r.AccountType) != "saving" || strings.ToLower(r.AccountType) != "checking"
}

func (r NewAccountRequest) Validate() *errs.AppError {
	if r.Amount < 5000 {
		return errs.NewValidationError("To open a new account you eed to deposit at least 5000.00")
	}
	if !r.allowedTypes() {
		return errs.NewValidationError("Account type should be checking or saving")
	}
	return nil
}
