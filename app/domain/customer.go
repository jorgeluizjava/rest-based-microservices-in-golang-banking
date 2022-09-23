package domain

import (
	"github.com/jorgeluizjava/banking/app/dto"
	"github.com/jorgeluizjava/banking/app/errs"
)

type Customer struct {
	Id          string
	Name        string
	City        string
	Zipcode     string
	DateofBirth string
	Status      string
}

func (c Customer) StatusAsText() string {
	statusAsString := "active"
	if c.Status == "0" {
		statusAsString = "inactive"
	}
	return statusAsString
}

func (c Customer) ToDto() dto.CustomerResponse {
	response := dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateofBirth: c.DateofBirth,
		Status:      c.StatusAsText(),
	}

	return response
}

type CustomerRepository interface {
	FindAll(status string) ([]Customer, *errs.AppError)
	ById(id string) (*Customer, *errs.AppError)
	FindAllByStatus(status string) ([]Customer, *errs.AppError)
}
