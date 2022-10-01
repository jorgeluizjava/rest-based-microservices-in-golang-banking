package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jorgeluizjava/banking/app/dto"
	"github.com/jorgeluizjava/banking/app/service"
)

type AccountHandlers struct {
	accountService service.AccountService
}

func NewAccountHandlers(accountService service.AccountService) AccountHandlers {
	return AccountHandlers{
		accountService: accountService,
	}
}

func (a AccountHandlers) NewAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		request.CustomerId = customerId
		account, appError := a.accountService.NewAccount(request)
		if appError != nil {
			writeResponse(w, appError.Code, appError.AsMessage())
		} else {
			writeResponse(w, http.StatusCreated, account)
		}
	}
}
