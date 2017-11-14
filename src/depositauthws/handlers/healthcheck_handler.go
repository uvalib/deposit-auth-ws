package handlers

import (
   "depositauthws/dao"
   "depositauthws/sis"
   "net/http"
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
      }

      if importErr != nil {
         importMsg = importErr.Error()
      }

      if exportErr != nil {
         exportMsg = exportErr.Error()
      }
   }

   encodeHealthCheckResponse(w, status, dbMsg, importMsg, exportMsg)
}

//
// end of file
//
