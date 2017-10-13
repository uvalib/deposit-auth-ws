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

//
// AuthorizationSearch -- search authorization request handler
//
func AuthorizationSearch(w http.ResponseWriter, r *http.Request) {

   token := r.URL.Query().Get("auth")
   id := r.URL.Query().Get("later")
   cid := r.URL.Query().Get("cid")
   createdAt := r.URL.Query().Get("created")
   exportedAt := r.URL.Query().Get("exported")

   // parameters OK ?
   if notEmpty(token) == false || (notEmpty(id) == false && notEmpty(cid) == false && notEmpty(createdAt) == false && notEmpty(exportedAt) == false) {
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

   if notEmpty(id) {
      // doing a search by ID
      reqs, err = dao.DB.SearchDepositAuthorizationByID(id)
   } else if notEmpty(cid) {
      // doing a search by computing ID
      reqs, err = dao.DB.SearchDepositAuthorizationByCid(cid)
   } else if notEmpty(createdAt) {
      // doing a search by create date
      reqs, err = dao.DB.SearchDepositAuthorizationByCreateDate(createdAt)
   } else if notEmpty(exportedAt) {
      // doing a search by export date
      reqs, err = dao.DB.SearchDepositAuthorizationByExportDate(exportedAt)
   }

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
