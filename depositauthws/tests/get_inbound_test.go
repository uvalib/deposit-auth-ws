package tests

import (
	"github.com/uvalib/deposit-auth-ws/depositauthws/api"
	"github.com/uvalib/deposit-auth-ws/depositauthws/client"
	"net/http"
	"strconv"
	"testing"
)

//
// get inbound tests
//

func TestGetInboundAll(t *testing.T) {

	expected := http.StatusOK
	status, details := client.GetInboundDepositAuthorization(cfg.Endpoint, firstID, goodToken)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}

	if len(details) == 0 {
		t.Fatalf("Expected more than 0 items, got 0\n")
	}

	ensureValidInbound(t, details, firstID)
	ensureValidAuthorizations(t, details)
}

func TestGetInboundSome(t *testing.T) {

	expected := http.StatusOK
	status, details := client.GetInboundDepositAuthorization(cfg.Endpoint, firstID, goodToken)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}

	if len(details) == 0 {
		t.Fatalf("Expected more than 0 items, got 0\n")
	}

	lastID, _ := strconv.Atoi(details[len(details)-1].InboundID)
	testID := strconv.Itoa(lastID - 1)

	status, details = client.GetInboundDepositAuthorization(cfg.Endpoint, testID, goodToken)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}

	if len(details) != 1 {
		t.Fatalf("Expected 1 item, got %d\n", len(details))
	}

	ensureValidInbound(t, details, testID)
	ensureValidAuthorizations(t, details)
}

func TestGetInboundNone(t *testing.T) {

	expected := http.StatusNotFound
	status, _ := client.GetInboundDepositAuthorization(cfg.Endpoint, lastID, goodToken)
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

func ensureValidInbound(t *testing.T, details []*api.Authorization, afterID string) {

	after, _ := strconv.Atoi(afterID)
	for _, e := range details {

		ID, _ := strconv.Atoi(e.InboundID)
		if ID <= after {
			t.Fatalf("Expected ID > %d, got ID of %d\n", after, ID)
		}
	}
}

//
// end of file
//
