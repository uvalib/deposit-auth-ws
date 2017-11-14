package tests

import (
   "depositauthws/client"
   "net/http"
   "testing"
)

//
// get inbound tests
//

func TestGetInboundHappyDay(t *testing.T) {

   expected := http.StatusOK
   status, details := client.GetInboundDepositAuthorization(cfg.Endpoint, firstID, goodToken )
   if status != expected {
      t.Fatalf("Expected %v, got %v\n", expected, status)
   }

   if len(details) == 0 {
      t.Fatalf("Expected more than 0 items, got 0\n" )
   }

   ensureValidAuthorizations(t, details)
}

func TestGetInboundNoneLater(t *testing.T) {

   expected := http.StatusNotFound
   status, _ := client.GetInboundDepositAuthorization(cfg.Endpoint, lastID, goodToken )
   if status != expected {
      t.Fatalf("Expected %v, got %v\n", expected, status)
   }
}

func TestGetInboundEmptyId(t *testing.T) {
   expected := http.StatusBadRequest
   status, _ := client.GetInboundDepositAuthorization(cfg.Endpoint, empty, goodToken)
   if status != expected {
      t.Fatalf("Expected %v, got %v\n", expected, status)
   }
}

func TestGetInboundBadToken(t *testing.T) {
   expected := http.StatusForbidden
   status, _ := client.GetInboundDepositAuthorization(cfg.Endpoint, goodID, badToken)
   if status != expected {
      t.Fatalf("Expected %v, got %v\n", expected, status)
   }
}

//
// end of file
//
