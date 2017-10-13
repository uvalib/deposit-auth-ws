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
// AuthorizationImport -- authorization import request handler
//
func AuthorizationImport(w http.ResponseWriter, r *http.Request) {

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

   // get the details ready to be imported
   imports, err := sis.Exchanger.Import()
   if err != nil {
      logger.Log(fmt.Sprintf("ERROR: %s\n", err.Error()))
      status := http.StatusInternalServerError
      encodeImportExportResponse(w, status,
         fmt.Sprintf("%s (%s)", http.StatusText(status), err),
         0)
      return
   }

   // if we have nothing to import, bail out
   if len(imports) == 0 {
      status := http.StatusOK
      encodeImportExportResponse(w, status, http.StatusText(status), 0)
      return
   }

   // import each record and keep track of progress
   okCount := 0
   duplicateCount := 0
   errorCount := 0
   for _, e := range imports {

      // check to see if this record already exists
      exists, err := dao.DB.DepositAuthorizationExists(*e)
      if err != nil {
         errorCount++
      } else {
         if exists == true {
            duplicateCount++
            logger.Log(fmt.Sprintf("record already exists, ignoring (%s/%s/%s/%s)", e.ComputingID, e.Degree, e.Plan, e.Title))
         } else {
            _, err = dao.DB.CreateDepositAuthorization(*e)
            if err != nil {
               logger.Log(fmt.Sprintf("Error inserting record; ignoring %s for (%s/%s/%s/%s)", err, e.ComputingID, e.Degree, e.Plan, e.Title))
               errorCount++
            } else {
               logger.Log(fmt.Sprintf("Success inserting (%s/%s/%s/%s)", e.ComputingID, e.Degree, e.Plan, e.Title))
               okCount++
            }
         }
      }
   }

   // log summary
   logger.Log(fmt.Sprintf("Import summary: %d success(es), %d duplicate(s), %d error(s)", okCount, duplicateCount, errorCount))

   // did we encounter any errors
   if errorCount != 0 {
      status := http.StatusCreated
      encodeImportExportResponse(w, status,
         fmt.Sprintf("%s (%d errors encountered)", http.StatusText(status), errorCount),
         okCount)
      return
   }

   // its all over
   status := http.StatusOK
   encodeImportExportResponse(w, status, http.StatusText(status), okCount)
}

//
// end of file
//
