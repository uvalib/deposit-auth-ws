package handlers

import (
	"depositauthws/dao"
	"depositauthws/sis"
	"net/http"
	"depositauthws/logger"
	"fmt"
)

//
// HealthCheckHandler -- do the healthcheck
//
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {

	status := http.StatusOK
	dbErr := dao.DB.CheckDB()
	importErr := sis.Exchanger.CheckImport()
	exportErr := sis.Exchanger.CheckExport()

	var dbMsg, importMsg, exportMsg string

	if dbErr != nil || importErr != nil || exportErr != nil {

		status = http.StatusInternalServerError

		if dbErr != nil {
			dbMsg = dbErr.Error()
			logger.Log(fmt.Sprintf( "ERROR: Database reports '%s'", dbMsg ) )
		}

		if importErr != nil {
			importMsg = importErr.Error()
			logger.Log(fmt.Sprintf( "ERROR: Importer reports '%s'", importMsg ) )
		}

		if exportErr != nil {
			exportMsg = exportErr.Error()
			logger.Log(fmt.Sprintf( "ERROR: Exporter reports '%s'", exportMsg ) )
		}
	}

	encodeHealthCheckResponse(w, status, dbMsg, importMsg, exportMsg)
}

//
// end of file
//
