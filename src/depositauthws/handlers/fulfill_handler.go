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
// FulfillHandler -- fulfill authorization request handler
//
func FulfillHandler(w http.ResponseWriter, r *http.Request) {

   vars := mux.Vars(r)
   id := vars["id"]
   token := r.URL.Query().Get("auth")
   did := r.URL.Query().Get("deposit")

   // parameters OK ?
   if isEmpty(id) || isEmpty(token) || isEmpty(did) {
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

   // handle the fulfill
   err = dao.DB.UpdateDepositAuthorizationByIDSetFulfilled(id, did)
   if err != nil {
      logger.Log(fmt.Sprintf("ERROR: %s\n", err.Error()))
      status := http.StatusInternalServerError
      encodeStandardResponse(w, status,
         fmt.Sprintf("%s (%s)", http.StatusText(status), err),
         nil)
      return
   }

   // its all over
   status := http.StatusOK
   encodeStandardResponse(w, status, http.StatusText(status), nil)
}

//
// end of file
//
