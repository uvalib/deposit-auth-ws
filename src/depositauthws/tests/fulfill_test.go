package tests

import (
   "depositauthws/client"
   "net/http"
   "testing"
)

//
// fulfill tests
//

func TestFulfillHappyDay(t *testing.T) {
   existing := getExistingAuthorization()
   if existing == nil {
      t.Fatalf("Unable to get existing authorization\n")
   }

   expected := http.StatusOK
   status := client.FulfillDepositAuthorization(cfg.Endpoint, existing.ID, goodDepositID, goodToken)
   if status != expected {
      t.Fatalf("Expected %v, got %v\n", expected, status)
   }
}

func TestFulfillEmptyId(t *testing.T) {
   expected := http.StatusBadRequest
   status := client.FulfillDepositAuthorization(cfg.Endpoint, empty, goodDepositID, goodToken)
   if status != expected {
      t.Fatalf("Expected %v, got %v\n", expected, status)
   }
}

func TestFulfillNotFoundId(t *testing.T) {
   expected := http.StatusNotFound
   status := client.FulfillDepositAuthorization(cfg.Endpoint, notFoundID, goodDepositID, goodToken)
   if status != expected {
      t.Fatalf("Expected %v, got %v\n", expected, status)
   }
}

func TestFulfillBadToken(t *testing.T) {
   expected := http.StatusForbidden
   status := client.FulfillDepositAuthorization(cfg.Endpoint, goodID, goodDepositID, badToken)
   if status != expected {
      t.Fatalf("Expected %v, got %v\n", expected, status)
   }
}

func TestFulfillBadDepositId(t *testing.T) {
   expected := http.StatusBadRequest
   status := client.FulfillDepositAuthorization(cfg.Endpoint, goodID, badDepositID, goodToken)
   if status != expected {
      t.Fatalf("Expected %v, got %v\n", expected, status)
   }
}

//
// end of file
//
