package service

import (
	"github.com/jorgeluizjava/banking/app/domain"
	"github.com/jorgeluizjava/banking/app/dto"
	"github.com/jorgeluizjava/banking/app/errs"
)

type AccountService interface {
	NewAccount(request dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
}

type DefaultAccountService struct {
	accountRepository domain.AccountRepository
}

func NewAccountService(accountRepository domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{
		accountRepository: accountRepository,
	}
}

func (s DefaultAccountService) NewAccount(request dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {

	if err := request.Validate(); err != nil {
		return nil, err
	}
	account := domain.NewAccount(request.CustomerId, request.AccountType, request.Amount)
	newAccount, err := s.accountRepository.Save(account)
	if err != nil {
		return nil, err
	}
	return newAccount.ToNewAccountResponseDto(), nil
}
