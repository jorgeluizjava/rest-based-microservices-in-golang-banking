package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/jorgeluizjava/banking/app/domain"
	"github.com/jorgeluizjava/banking/app/service"
)

func Start() {
	router := mux.NewRouter()

	dbClient := getDbClient()

	customerRepository := domain.NewCustomerRepositoryDb(dbClient)
	accountRepositoryDb := domain.NewAccountRepositoryDb(dbClient)

	customerService := service.NewCustomerService(customerRepository)
	accountService := service.NewAccountService(accountRepositoryDb)

	ch := CustomerHandlers{
		customerService: customerService,
	}

	ah := AccountHandlers{
		accountService: accountService,
	}

	router.HandleFunc("/customers", ch.GetAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.GetCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/accounts", ah.NewAccount).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe("localhost:8000", router))
}

func getDbClient() *sqlx.DB {
	dbUser := os.Getenv("DB_USER")
	dbPasswd := os.Getenv("DB_PASSWD")
	dbAddr := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPasswd, dbAddr, dbPort, dbName)

	//client, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/banking")
	client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return client
}
