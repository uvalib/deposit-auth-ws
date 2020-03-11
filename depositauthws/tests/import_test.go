package tests

import (
	"github.com/uvalib/deposit-auth-ws/depositauthws/client"
	"net/http"
	"testing"
)

//
// import tests
//

func TestImportHappyDay(t *testing.T) {
	expected := http.StatusOK
	status, _, _, _, errCount := client.ImportDepositAuthorization(cfg.Endpoint, goodToken(cfg.Secret))
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
	if errCount != 0 {
		t.Fatalf("Unexpected error count, got %v\n", errCount)
	}
}

func TestImportBadToken(t *testing.T) {
	expected := http.StatusForbidden
	status, _, _, _, _ := client.ImportDepositAuthorization(cfg.Endpoint, badToken(cfg.Secret))
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

//
// end of file
//
