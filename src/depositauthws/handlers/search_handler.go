package handlers

import (
	"depositauthws/api"
	"depositauthws/authtoken"
	"depositauthws/config"
	"depositauthws/dao"
	"depositauthws/logger"
	"fmt"
	"net/http"
	//"golang.org/x/net/http2/hpack"
)

func AuthorizationSearch(w http.ResponseWriter, r *http.Request) {

	token := r.URL.Query().Get("auth")
	id := r.URL.Query().Get("later")
	cid := r.URL.Query().Get("cid")
	created_at := r.URL.Query().Get("created")
	exported_at := r.URL.Query().Get("exported")

	// parameters OK ?
	if NotEmpty(token) == false || (NotEmpty(id) == false && NotEmpty(cid) == false && NotEmpty(created_at) == false && NotEmpty(exported_at) == false) {
		status := http.StatusBadRequest
		EncodeStandardResponse(w, status, http.StatusText(status), nil)
		return
	}

	// validate the token
	if authtoken.Validate(config.Configuration.AuthTokenEndpoint, token, config.Configuration.Timeout) == false {
		status := http.StatusForbidden
		EncodeStandardResponse(w, status, http.StatusText(status), nil)
		return
	}

	var reqs []*api.Authorization
	var err error

	if NotEmpty(id) {
		// doing a search by ID
		reqs, err = dao.Database.SearchDepositAuthorizationById(id)
	} else if NotEmpty(cid) {
		// doing a search by computing ID
		reqs, err = dao.Database.SearchDepositAuthorizationByCid(cid)
	} else if NotEmpty(created_at) {
		// doing a search by create date
		reqs, err = dao.Database.SearchDepositAuthorizationByCreateDate(created_at)
	} else if NotEmpty(exported_at) {
		// doing a search by export date
		reqs, err = dao.Database.SearchDepositAuthorizationByExportDate(exported_at)
	}

	if err != nil {
		logger.Log(fmt.Sprintf("ERROR: %s\n", err.Error()))
		status := http.StatusInternalServerError
		EncodeStandardResponse(w, status,
			fmt.Sprintf("%s (%s)", http.StatusText(status), err),
			nil)
		return
	}

	if reqs == nil || len(reqs) == 0 {
		status := http.StatusNotFound
		EncodeStandardResponse(w, status, http.StatusText(status), nil)
		return
	}

	// do necessary field mappings
	MapResultsFieldValues(reqs)

	status := http.StatusOK
	EncodeStandardResponse(w, status, http.StatusText(status), reqs)
}
