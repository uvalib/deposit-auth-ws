package tests

import (
   "depositauthws/client"
   "net/http"
   "testing"
)

//
// import tests
//

func TestImportHappyDay(t *testing.T) {
   expected := http.StatusOK
   status, _ := client.ImportDepositAuthorization(cfg.Endpoint, goodToken)
   if status != expected {
      t.Fatalf("Expected %v, got %v\n", expected, status)
   }
}

func TestImportBadToken(t *testing.T) {
   expected := http.StatusForbidden
   status, _ := client.ImportDepositAuthorization(cfg.Endpoint, badToken)
   if status != expected {
      t.Fatalf("Expected %v, got %v\n", expected, status)
   }
}

//
// end of file
//
