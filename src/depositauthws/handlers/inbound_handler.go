package handlers

import (
   "depositauthws/api"
   "depositauthws/authtoken"
   "depositauthws/config"
   //"depositauthws/dao"
   "depositauthws/logger"
   "fmt"
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
   if authtoken.Validate(config.Configuration.AuthTokenEndpoint, token, config.Configuration.Timeout) == false {
      status := http.StatusForbidden
      encodeStandardResponse(w, status, http.StatusText(status), nil)
      return
   }

   var reqs []*api.Authorization
   var err error

//   if notEmpty(id) {
//      // doing a search by ID
//      reqs, err = dao.DB.SearchDepositAuthorizationByID(id)
//   }

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
