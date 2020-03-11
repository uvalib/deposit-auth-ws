package tests

import (
	"github.com/uvalib/deposit-auth-ws/depositauthws/client"
	"net/http"
	"testing"
)

//
// get auth tests
//

func TestGetAuthHappyDay(t *testing.T) {

	status, existing := getExistingAuthorization()
	if status != http.StatusOK {
		t.Fatalf("Unable to get existing authorization: status %d\n", status)
	}

	expected := http.StatusOK
	status, details := client.GetDepositAuthorization(cfg.Endpoint, existing.ID, goodToken(cfg.Secret))
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}

	if len(details) != 1 {
		t.Fatalf("Expected 1 item, got %v\n", len(details))
	}

	ensureValidAuthorizations(t, details)
}

func TestGetAuthEmptyId(t *testing.T) {
	expected := http.StatusBadRequest
	status, _ := client.GetDepositAuthorization(cfg.Endpoint, empty, goodToken(cfg.Secret))
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

func TestGetAuthNotFoundId(t *testing.T) {
	expected := http.StatusNotFound
	status, _ := client.GetDepositAuthorization(cfg.Endpoint, notFoundID, goodToken(cfg.Secret))
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

func TestGetAuthBadToken(t *testing.T) {
	expected := http.StatusForbidden
	status, _ := client.GetDepositAuthorization(cfg.Endpoint, goodID, badToken(cfg.Secret))
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

//
// end of file
//
