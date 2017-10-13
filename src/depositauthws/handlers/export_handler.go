package handlers

import (
   "depositauthws/authtoken"
   "depositauthws/config"
   "depositauthws/dao"
   "depositauthws/logger"
   "depositauthws/sis"
   "fmt"
   "net/http"
)

//
// AuthorizationExport -- export authorizations request handler
//
func AuthorizationExport(w http.ResponseWriter, r *http.Request) {

   token := r.URL.Query().Get("auth")

   // parameters OK ?
   if notEmpty(token) == false {
      status := http.StatusBadRequest
      encodeImportExportResponse(w, status, http.StatusText(status), 0)
      return
   }

   // validate the token
   if authtoken.Validate(config.Configuration.AuthTokenEndpoint, token, config.Configuration.Timeout) == false {
      status := http.StatusForbidden
      encodeImportExportResponse(w, status, http.StatusText(status), 0)
      return
   }

   // get the details ready to be exported
   exports, err := dao.DB.GetDepositAuthorizationForExport()
   if err != nil {
      logger.Log(fmt.Sprintf("ERROR: %s\n", err.Error()))
      status := http.StatusInternalServerError
      encodeImportExportResponse(w, status,
         fmt.Sprintf("%s (%s)", http.StatusText(status), err),
         0)
      return
   }

   // do the export
   err = sis.Exchanger.Export(exports)
   if err != nil {
      logger.Log(fmt.Sprintf("ERROR: %s\n", err.Error()))
      status := http.StatusInternalServerError
      encodeImportExportResponse(w, status,
         fmt.Sprintf("%s (%s)", http.StatusText(status), err),
         0)
      return
   }

   // update the status so we do not export them again
   err = dao.DB.UpdateExportedDepositAuthorization(exports)
   if err != nil {
      logger.Log(fmt.Sprintf("ERROR: %s\n", err.Error()))
      status := http.StatusInternalServerError
      encodeImportExportResponse(w, status,
         fmt.Sprintf("%s (%s)", http.StatusText(status), err),
         0)
      return
   }

   // log summary
   logger.Log(fmt.Sprintf("Export summary: %d record(s) exported", len(exports)))

   // its all over
   status := http.StatusOK
   encodeImportExportResponse(w, status, http.StatusText(status), len(exports))
}

//
// end of file
//
