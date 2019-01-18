package handlers

import (
	"fmt"
	"github.com/uvalib/deposit-auth-ws/depositauthws/authtoken"
	"github.com/uvalib/deposit-auth-ws/depositauthws/config"
	"github.com/uvalib/deposit-auth-ws/depositauthws/dao"
	"github.com/uvalib/deposit-auth-ws/depositauthws/logger"
	"net/http"
)

//
// InboundHandler -- search authorization request handler
//
func InboundHandler(w http.ResponseWriter, r *http.Request) {

	token := r.URL.Query().Get("auth")
	after := r.URL.Query().Get("after")

	// parameters OK ?
	if isEmpty(token) || isEmpty(after) {
		status := http.StatusBadRequest
		encodeStandardResponse(w, status, http.StatusText(status), nil)
		return
	}

	// validate the token
	if authtoken.Validate(config.Configuration.AuthTokenEndpoint, token, config.Configuration.ServiceTimeout) == false {
		status := http.StatusForbidden
		encodeStandardResponse(w, status, http.StatusText(status), nil)
		return
	}

	// get from the inbound queue
	reqs, err := dao.Store.GetInbound(after)

	if err != nil {
		logger.Log(fmt.Sprintf("ERROR: %s\n", err.Error()))
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
