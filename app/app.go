package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jorgeluizjava/banking/app/domain"
	"github.com/jorgeluizjava/banking/app/service"
)

func Start() {
	router := mux.NewRouter()

	customerRepository := domain.NewCustomerRepositoryStub()
	customerService := service.NewCustomerService(customerRepository)

	ch := CustomerHandlers{
		customerService: customerService,
	}

	router.HandleFunc("/customers", ch.GetAllCustomers).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
