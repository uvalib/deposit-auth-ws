package tests

import (
   "depositauthws/client"
   "net/http"
   "testing"
)

//
// get inbound tests
//

func TestGetHappyDay(t *testing.T) {

   existing := getExistingAuthorization()
   if existing == nil {
      t.Fatalf("Unable to get existing authorization\n")
   }

   expected := http.StatusOK
   status, details := client.GetDepositAuthorization(cfg.Endpoint, existing.ID, goodToken)
   if status != expected {
      t.Fatalf("Expected %v, got %v\n", expected, status)
   }

   if len(details) != 1 {
      t.Fatalf("Expected 1 item, got %v\n", len(details))
   }

   ensureValidAuthorizations(t, details)
}

func TestGetEmptyId(t *testing.T) {
   expected := http.StatusBadRequest
   status, _ := client.GetDepositAuthorization(cfg.Endpoint, empty, goodToken)
   if status != expected {
      t.Fatalf("Expected %v, got %v\n", expected, status)
   }
}

func TestGetNotFoundId(t *testing.T) {
   expected := http.StatusNotFound
   status, _ := client.GetDepositAuthorization(cfg.Endpoint, notFoundID, goodToken)
   if status != expected {
      t.Fatalf("Expected %v, got %v\n", expected, status)
   }
}

func TestGetBadToken(t *testing.T) {
   expected := http.StatusForbidden
   status, _ := client.GetDepositAuthorization(cfg.Endpoint, goodID, badToken)
   if status != expected {
      t.Fatalf("Expected %v, got %v\n", expected, status)
   }
}

//
// end of file
//
