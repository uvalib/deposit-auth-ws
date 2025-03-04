package handlers

import (
	"fmt"
	"github.com/uvalib/deposit-auth-ws/depositauthws/api"
	"github.com/uvalib/deposit-auth-ws/depositauthws/authtoken"
	"github.com/uvalib/deposit-auth-ws/depositauthws/config"
	"github.com/uvalib/deposit-auth-ws/depositauthws/dao"
	"github.com/uvalib/deposit-auth-ws/depositauthws/logger"
	"net/http"
)

// SearchHandler -- search authorization request handler
func SearchHandler(w http.ResponseWriter, r *http.Request) {

	token := r.URL.Query().Get("auth")
	cid := r.URL.Query().Get("cid")
	createdAt := r.URL.Query().Get("created")
	exportedAt := r.URL.Query().Get("exported")

	// parameters OK ?
	if isEmpty(token) || (isEmpty(cid) && isEmpty(createdAt) && isEmpty(exportedAt)) {
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

	var reqs []*api.Authorization
	var err error

	if isEmpty(cid) == false {
		// doing a search by computing ID
		reqs, err = dao.Store.SearchDepositAuthorizationByCid(cid)
	} else if isEmpty(createdAt) == false {
		// doing a search by create date
		reqs, err = dao.Store.SearchDepositAuthorizationByCreateDate(createdAt)
	} else if isEmpty(exportedAt) == false {
		// doing a search by export date
		reqs, err = dao.Store.SearchDepositAuthorizationByExportDate(exportedAt)
	}

	if err != nil {
		logger.Log(fmt.Sprintf("ERROR: %s", err.Error()))
		status := http.StatusInternalServerError
		encodeStandardResponse(w, status,
			fmt.Sprintf("%s (%s)", http.StatusText(status), err),
			nil)
		return
	}

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
