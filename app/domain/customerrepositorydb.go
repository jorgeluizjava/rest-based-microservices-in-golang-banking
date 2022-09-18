package domain

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jorgeluizjava/banking/app/errs"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (s CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {

	customers := make([]Customer, 0)
	var rows *sql.Rows
	var err error
	if status == "" {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		rows, err = s.client.Query(findAllSql)
		if err != nil {
			log.Println("Error while querying customer table " + err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}
	} else {
		findAllCustomersByStatus := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
		rows, err = s.client.Query(findAllCustomersByStatus, status)
		if err != nil {
			log.Println("Error while querying customer table " + err.Error())
			return []Customer{}, errs.NewUnexpectedError("unexpected database error")
		}
	}

	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
		if err != nil {
			log.Println("Error while scanning customer " + err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}
		customers = append(customers, c)
	}

	return customers, nil
}

func (s CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	sqlCustomer := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"
	row := s.client.QueryRow(sqlCustomer, id)
	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError(fmt.Sprintf("Customer not found for id: %s", id))
		}
		log.Println("Error while scanning customer by id " + id)
		return nil, errs.NewUnexpectedError("unexpected database error")
	}
	return &c, nil
}

func (s CustomerRepositoryDb) FindAllByStatus(status string) ([]Customer, *errs.AppError) {
	findAllCustomersByStatus := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
	rows, err := s.client.Query(findAllCustomersByStatus, status)
	if err != nil {
		log.Println("Error while querying customer table " + err.Error())
		return []Customer{}, errs.NewUnexpectedError("unexpected database error")
	}
	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
		if err != nil {
			log.Println("Error while scanning customer " + err.Error())
			return []Customer{}, errs.NewUnexpectedError("unexpected database error")
		}
		customers = append(customers, c)
	}
	return customers, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {

	client, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return CustomerRepositoryDb{client: client}
}
