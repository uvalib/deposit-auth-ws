package handlers

import (
   "depositauthws/authtoken"
   "depositauthws/config"
   "depositauthws/dao"
   "depositauthws/logger"
   "fmt"
   "github.com/gorilla/mux"
   "net/http"
)

//
// AuthorizationGet -- get authorization request handler
//
func AuthorizationGet(w http.ResponseWriter, r *http.Request) {

   vars := mux.Vars(r)
   id := vars["id"]
   token := r.URL.Query().Get("auth")

   // parameters OK ?
   if notEmpty(id) == false || notEmpty(token) == false {
      status := http.StatusBadRequest
      encodeStandardResponse(w, status, http.StatusText(status), nil)
      return
   }

   // validate the token
   if authtoken.Validate(config.Configuration.AuthTokenEndpoint, token, config.Configuration.Timeout) == false {
      status := http.StatusForbidden
      encodeStandardResponse(w, status, http.StatusText(status), nil)
      return
   }

   // get the authorization details
   reqs, err := dao.DB.GetDepositAuthorizationByID(id)
   if err != nil {
      logger.Log(fmt.Sprintf("ERROR: %s\n", err.Error()))
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

   // do necessary field mappings
   mapResultsFieldValues(reqs)

   status := http.StatusOK
   encodeStandardResponse(w, status, http.StatusText(status), reqs)
}

//
// end of file
//
