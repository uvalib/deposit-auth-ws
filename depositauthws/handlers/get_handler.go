package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/uvalib/deposit-auth-ws/depositauthws/authtoken"
	"github.com/uvalib/deposit-auth-ws/depositauthws/config"
	"github.com/uvalib/deposit-auth-ws/depositauthws/dao"
	"github.com/uvalib/deposit-auth-ws/depositauthws/logger"
	"net/http"
)

//
// GetHandler -- get authorization request handler
//
func GetHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]
	token := r.URL.Query().Get("auth")

	// parameters OK ?
	if isEmpty(id) || isEmpty(token) {
		status := http.StatusBadRequest
		encodeStandardResponse(w, status, http.StatusText(status), nil)
		return
	}

	// validate the token
	if authtoken.Validate(config.Configuration.SharedSecret, token) == false {
		status := http.StatusForbidden
		encodeStandardResponse(w, status, http.StatusText(status), nil)
		return
	}

	// get the authorization details
	reqs, err := dao.Store.GetDepositAuthorizationByID(id)
	if err != nil {
		logger.Log(fmt.Sprintf("ERROR: %s", err.Error()))
		status := http.StatusInternalServerError
		encodeStandardResponse(w, status,
			fmt.Sprintf("%s (%s)", http.StatusText(status), err),
			nil)
		return
	}

	// we did not find the item, return 404
	if reqs == nil || len(reqs) == 0 {
		status := http.StatusNotFound
		encodeStandardResponse(w, status, http.StatusText(status), nil)
		return
	}

	// do necessary field mappings
	mapResultsFieldValues(reqs)

	status := http.StatusOK
	encodeStandardResponse(w, status, http.StatusText(status), reqs)
}

//
// end of file
//
