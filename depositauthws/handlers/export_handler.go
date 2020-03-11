package handlers

import (
	"fmt"
	"github.com/uvalib/deposit-auth-ws/depositauthws/authtoken"
	"github.com/uvalib/deposit-auth-ws/depositauthws/config"
	"github.com/uvalib/deposit-auth-ws/depositauthws/dao"
	"github.com/uvalib/deposit-auth-ws/depositauthws/logger"
	"github.com/uvalib/deposit-auth-ws/depositauthws/sis"
	"net/http"
)

//
// ExportHandler -- export authorizations request handler
//
func ExportHandler(w http.ResponseWriter, r *http.Request) {

	token := r.URL.Query().Get("auth")

	// parameters OK ?
	if isEmpty(token) {
		status := http.StatusBadRequest
		encodeExportResponse(w, status, http.StatusText(status), 0, 0)
		return
	}

	// validate the token
	if authtoken.Validate(config.Configuration.SharedSecret, token) == false {
		status := http.StatusForbidden
		encodeExportResponse(w, status, http.StatusText(status), 0, 0)
		return
	}

	// get the details ready to be exported
	exports, err := dao.Store.GetDepositAuthorizationForExport()
	if err != nil {
		logger.Log(fmt.Sprintf("ERROR: %s\n", err.Error()))
		status := http.StatusInternalServerError
		encodeExportResponse(w, status,
			fmt.Sprintf("%s (%s)", http.StatusText(status), err),
			0, 0)
		return
	}

	// do the export
	err = sis.Exchanger.Export(exports)
	if err != nil {
		logger.Log(fmt.Sprintf("ERROR: %s\n", err.Error()))
		status := http.StatusInternalServerError
		encodeExportResponse(w, status,
			fmt.Sprintf("%s (%s)", http.StatusText(status), err),
			0, 0)
		return
	}

	// update the status so we do not export them again
	err = dao.Store.UpdateExportedDepositAuthorization(exports)
	if err != nil {
		logger.Log(fmt.Sprintf("ERROR: %s\n", err.Error()))
		status := http.StatusInternalServerError
		encodeExportResponse(w, status,
			fmt.Sprintf("%s (%s)", http.StatusText(status), err),
			0, 0)
		return
	}

	// log summary
	logger.Log(fmt.Sprintf("Export summary: %d record(s) exported", len(exports)))

	// its all over
	status := http.StatusOK
	encodeExportResponse(w, status, http.StatusText(status), len(exports), 0)
}

//
// end of file
//
