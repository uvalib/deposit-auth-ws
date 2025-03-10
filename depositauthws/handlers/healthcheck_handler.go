package handlers

import (
	"fmt"
	"github.com/uvalib/deposit-auth-ws/depositauthws/dao"
	"github.com/uvalib/deposit-auth-ws/depositauthws/logger"
	"github.com/uvalib/deposit-auth-ws/depositauthws/sis"
	"net/http"
)

// HealthCheckHandler -- do the healthcheck
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {

	status := http.StatusOK
	dbErr := dao.Store.Check()
	importErr := sis.Exchanger.CheckImport()
	exportErr := sis.Exchanger.CheckExport()

	var dbMsg, importMsg, exportMsg string

	if dbErr != nil || importErr != nil || exportErr != nil {

		status = http.StatusInternalServerError

		if dbErr != nil {
			dbMsg = dbErr.Error()
			logger.Log(fmt.Sprintf("ERROR: Datastore reports '%s'", dbMsg))
		}

		if importErr != nil {
			importMsg = importErr.Error()
			logger.Log(fmt.Sprintf("ERROR: Importer reports '%s'", importMsg))
		}

		if exportErr != nil {
			exportMsg = exportErr.Error()
			logger.Log(fmt.Sprintf("ERROR: Exporter reports '%s'", exportMsg))
		}
	}

	encodeHealthCheckResponse(w, status, dbMsg, importMsg, exportMsg)
}

//
// end of file
//
