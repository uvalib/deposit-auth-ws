package tests

import (
	"github.com/uvalib/deposit-auth-ws/depositauthws/api"
	"github.com/uvalib/deposit-auth-ws/depositauthws/client"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"
)

type testConfig struct {
	Endpoint string
	Token    string
}

var cfg = loadConfig()

var firstID = "0"
var lastID = "9999999"
var goodID = "1"
var notFoundID = "x"
var goodToken = cfg.Token
var badToken = "badness"
var goodDepositID = "libra:12345"
var badDepositID = " "
var empty = " "
var goodDate = "2016-01-01"

func getExistingAuthorization() (int, *api.Authorization) {

	status, details := client.SearchDepositAuthorizationByCreated(cfg.Endpoint, goodDate, goodToken)
	if status == http.StatusOK {
		if len(details) != 0 {
			return status, details[0]
		}
	}

	return status, nil
}

func ensureValidAuthorizations(t *testing.T, details []*api.Authorization) {

	for _, e := range details {
		if emptyField(e.ID) ||
			emptyField(e.EmployeeID) ||
			emptyField(e.ComputingID) ||
			emptyField(e.FirstName) ||
			//emptyField( e.MiddleName ) ||
			emptyField(e.LastName) ||
			emptyField(e.Career) ||
			emptyField(e.Program) ||
			emptyField(e.Plan) ||
			emptyField(e.Degree) ||
			//emptyField( e.Department ) ||
			//emptyField( e.Title ) ||
			emptyField(e.DocType) ||
			//emptyField( e.LibraID ) ||
			emptyField(e.Status) ||
			emptyField(e.CreatedAt) ||
			//emptyField( e.ApprovedAt ) ||
			//emptyField( e.AcceptedAt ) ||
			//emptyField( e.ExportedAt ) ||
			//emptyField( e.UpdatedAt ) ||
			emptyField(e.Status) {
			log.Printf("%v", e)
			t.Fatalf("Expected non-empty field but one is empty\n")
		}
	}
}

func emptyField(field string) bool {
	return len(strings.TrimSpace(field)) == 0
}

func loadConfig() testConfig {

	data, err := ioutil.ReadFile("service_test.yml")
	if err != nil {
		log.Fatal(err)
	}

	var c testConfig
	if err := yaml.Unmarshal(data, &c); err != nil {
		log.Fatal(err)
	}

	log.Printf("Test config; endpoint   [%s]\n", c.Endpoint)
	log.Printf("Test config; auth token [%s]\n", c.Token)

	return c
}

//
// end of file
//
