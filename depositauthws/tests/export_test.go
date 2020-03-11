package tests

import (
	"github.com/uvalib/deposit-auth-ws/depositauthws/client"
	"net/http"
	"testing"
)

//
// export tests
//

func TestExportHappyDay(t *testing.T) {
	expected := http.StatusOK
	status, _, errCount := client.ExportDepositAuthorization(cfg.Endpoint, goodToken(cfg.Secret))
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
	if errCount != 0 {
		t.Fatalf("Unexpected error count, got %v\n", errCount)
	}
}

func TestExportBadToken(t *testing.T) {
	expected := http.StatusForbidden
	status, _, _ := client.ExportDepositAuthorization(cfg.Endpoint, badToken(cfg.Secret))
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

//
// end of file
//
