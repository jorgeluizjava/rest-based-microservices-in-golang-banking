package app

import (
	"encoding/json"
	"net/http"

	"github.com/jorgeluizjava/banking/app/service"
)

type CustomerHandlers struct {
	customerService service.CustomerService
}

func (ch *CustomerHandlers) GetAllCustomers(w http.ResponseWriter, r *http.Request) {

	customers, _ := ch.customerService.GetAllCustomers()
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}
