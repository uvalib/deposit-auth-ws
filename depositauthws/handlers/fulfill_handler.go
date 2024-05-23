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

// FulfillHandler -- fulfill authorization request handler
func FulfillHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]
	token := r.URL.Query().Get("auth")
	did := r.URL.Query().Get("deposit")

	// parameters OK ?
	if isEmpty(id) || isEmpty(token) || isEmpty(did) {
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

	// handle the fulfill
	err = dao.Store.UpdateDepositAuthorizationByIDSetFulfilled(id, did)
	if err != nil {
		logger.Log(fmt.Sprintf("ERROR: %s", err.Error()))
		status := http.StatusInternalServerError
		encodeStandardResponse(w, status,
			fmt.Sprintf("%s (%s)", http.StatusText(status), err),
			nil)
		return
	}

	// its all over
	status := http.StatusOK
	encodeStandardResponse(w, status, http.StatusText(status), nil)
}

//
// end of file
//
