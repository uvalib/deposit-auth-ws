package tests

import (
   "depositauthws/client"
   "net/http"
   "testing"
)

//
// search by ID tests
//

func TestSearchByIdHappyDay(t *testing.T) {
   expected := http.StatusOK
   status, details := client.SearchDepositAuthorizationByID(cfg.Endpoint, "0", goodToken)
   if status != expected {
      t.Fatalf("Expected %v, got %v\n", expected, status)
   }
   ensureValidAuthorizations(t, details)
}

func TestSearchByIdEmptyId(t *testing.T) {
   expected := http.StatusBadRequest
   status, _ := client.SearchDepositAuthorizationByID(cfg.Endpoint, empty, goodToken)
   if status != expected {
      t.Fatalf("Expected %v, got %v\n", expected, status)
   }
}

func TestSearchByIdBadToken(t *testing.T) {
   expected := http.StatusForbidden
   status, _ := client.SearchDepositAuthorizationByID(cfg.Endpoint, goodID, badToken)
   if status != expected {
      t.Fatalf("Expected %v, got %v\n", expected, status)
   }
}

//
// search by computing ID tests
//

func TestSearchByCidHappyDay(t *testing.T) {

   existing := getExistingAuthorization()
   if existing == nil {
      t.Fatalf("Unable to get existing authorization\n")
   }

   expected := http.StatusOK
   status, details := client.SearchDepositAuthorizationByCid(cfg.Endpoint, existing.ComputingID, goodToken)
   if status != expected {
      t.Fatalf("Expected %v, got %v\n", expected, status)
   }
   ensureValidAuthorizations(t, details)
}

func TestSearchByCidEmptyCid(t *testing.T) {
   expected := http.StatusBadRequest
   status, _ := client.SearchDepositAuthorizationByCid(cfg.Endpoint, empty, goodToken)
   if status != expected {
      t.Fatalf("Expected %v, got %v\n", expected, status)
   }
}

func TestSearchByCidBadToken(t *testing.T) {
   expected := http.StatusForbidden
   status, _ := client.SearchDepositAuthorizationByCid(cfg.Endpoint, goodID, badToken)
   if status != expected {
      t.Fatalf("Expected %v, got %v\n", expected, status)
   }
}

//
// search by create date tests
//

func TestSearchByCreatedHappyDay(t *testing.T) {
   expected := http.StatusOK
   status, details := client.SearchDepositAuthorizationByCreated(cfg.Endpoint, goodDate, goodToken)
   if status != expected {
      t.Fatalf("Expected %v, got %v\n", expected, status)
   }
   ensureValidAuthorizations(t, details)
}

func TestSearchByCreatedEmptyCreated(t *testing.T) {
   expected := http.StatusBadRequest
   status, _ := client.SearchDepositAuthorizationByCreated(cfg.Endpoint, empty, goodToken)
   if status != expected {
      t.Fatalf("Expected %v, got %v\n", expected, status)
   }
}

func TestSearchByCreatedBadToken(t *testing.T) {
   expected := http.StatusForbidden
   status, _ := client.SearchDepositAuthorizationByCreated(cfg.Endpoint, goodDate, badToken)
   if status != expected {
      t.Fatalf("Expected %v, got %v\n", expected, status)
   }
}

//
// search by export date tests
//

func TestSearchByExportedHappyDay(t *testing.T) {
   expected := http.StatusOK
   status, details := client.SearchDepositAuthorizationByExported(cfg.Endpoint, goodDate, goodToken)
   if status != expected {
      t.Fatalf("Expected %v, got %v\n", expected, status)
   }
   ensureValidAuthorizations(t, details)
}

func TestSearchByExportedEmptyExported(t *testing.T) {
   expected := http.StatusBadRequest
   status, _ := client.SearchDepositAuthorizationByExported(cfg.Endpoint, empty, goodToken)
   if status != expected {
      t.Fatalf("Expected %v, got %v\n", expected, status)
   }
}

func TestSearchByExportedBadToken(t *testing.T) {
   expected := http.StatusForbidden
   status, _ := client.SearchDepositAuthorizationByExported(cfg.Endpoint, goodDate, badToken)
   if status != expected {
      t.Fatalf("Expected %v, got %v\n", expected, status)
   }
}

//
// end of file
//
