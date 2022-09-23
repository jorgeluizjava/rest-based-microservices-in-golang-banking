package service

import (
	"github.com/jorgeluizjava/banking/app/domain"
	"github.com/jorgeluizjava/banking/app/dto"
	"github.com/jorgeluizjava/banking/app/errs"
)

type CustomerService interface {
	GetAllCustomers(status string) ([]dto.CustomerResponse, *errs.AppError)
	GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError)
	GetAllCustomersByStatus(status string) ([]dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}

func (s DefaultCustomerService) GetAllCustomers(status string) ([]dto.CustomerResponse, *errs.AppError) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}

	customers, err := s.repo.FindAll(status)
	if err != nil {
		return []dto.CustomerResponse{}, err
	}

	customersResponse := make([]dto.CustomerResponse, len(customers))

	for i, customer := range customers {
		customersResponse[i] = customer.ToDto()
	}

	return customersResponse, nil
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {

	customer, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}

	response := customer.ToDto()
	return &response, nil
}

func (s DefaultCustomerService) GetAllCustomersByStatus(status string) ([]dto.CustomerResponse, *errs.AppError) {

	customers, err := s.repo.FindAllByStatus(status)
	if err != nil {
		return []dto.CustomerResponse{}, err
	}

	customersResponse := make([]dto.CustomerResponse, len(customers))

	for i, customer := range customers {
		customersResponse[i] = customer.ToDto()
	}

	return customersResponse, nil
}
