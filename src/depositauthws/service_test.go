package main

import (
   "depositauthws/api"
   "depositauthws/client"
   "io/ioutil"
   "log"
   "net/http"
   "strings"
   "testing"
	"gopkg.in/yaml.v2"
)

type TestConfig struct {
   Endpoint string
   Token    string
}

var cfg = loadConfig()

var goodId = "1"
var notFoundId = "x"
var goodToken = cfg.Token
var badToken = "badness"
var goodDepositId = "libra:12345"
var badDepositId = " "
var empty = " "
var goodDate = "2016-01-01"

//
// healthcheck tests
//

func TestHealthCheck(t *testing.T) {
   expected := http.StatusOK
   status := client.HealthCheck(cfg.Endpoint)
   if status != expected {
      t.Fatalf("Expected %v, got %v\n", expected, status)
   }
}

//
// version tests
//

func TestVersionCheck(t *testing.T) {
   expected := http.StatusOK
   status, version := client.VersionCheck(cfg.Endpoint)
   if status != expected {
      t.Fatalf("Expected %v, got %v\n", expected, status)
   }

   if len(version) == 0 {
      t.Fatalf("Expected non-zero length version string\n")
   }
}

//
// runtime tests
//

func TestRuntimeCheck(t *testing.T) {
   expected := http.StatusOK
   status, runtime := client.RuntimeCheck(cfg.Endpoint)
   if status != expected {
      t.Fatalf("Expected %v, got %v\n", expected, status)
   }

   if runtime == nil {
      t.Fatalf("Expected non-nil runtime info\n")
   }

   if len( runtime.Version ) == 0 ||
      runtime.AllocatedMemory == 0 ||
      runtime.CpuCount == 0 ||
      runtime.GoRoutineCount == 0 ||
      runtime.ObjectCount == 0 {
      t.Fatalf("Expected non-zero value in runtime info but one is zero\n")
   }
}

//
// get tests
//

func TestGetHappyDay(t *testing.T) {

   existing := getExistingAuthorization()
   if existing == nil {
      t.Fatalf("Unable to get existing authorization\n")
   }

   expected := http.StatusOK
   status, details := client.GetDepositAuthorization(cfg.Endpoint, existing.Id, goodToken)
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
   status, _ := client.GetDepositAuthorization(cfg.Endpoint, notFoundId, goodToken)
   if status != expected {
      t.Fatalf("Expected %v, got %v\n", expected, status)
   }
}

func TestGetBadToken(t *testing.T) {
   expected := http.StatusForbidden
   status, _ := client.GetDepositAuthorization(cfg.Endpoint, goodId, badToken)
   if status != expected {
      t.Fatalf("Expected %v, got %v\n", expected, status)
   }
}

//
// search by ID tests
//

func TestSearchByIdHappyDay(t *testing.T) {
   expected := http.StatusOK
   status, details := client.SearchDepositAuthorizationById(cfg.Endpoint, "0", goodToken)
   if status != expected {
      t.Fatalf("Expected %v, got %v\n", expected, status)
   }
   ensureValidAuthorizations(t, details)
}

func TestSearchByIdEmptyId(t *testing.T) {
   expected := http.StatusBadRequest
   status, _ := client.SearchDepositAuthorizationById(cfg.Endpoint, empty, goodToken)
   if status != expected {
      t.Fatalf("Expected %v, got %v\n", expected, status)
   }
}

func TestSearchByIdBadToken(t *testing.T) {
   expected := http.StatusForbidden
   status, _ := client.SearchDepositAuthorizationById(cfg.Endpoint, goodId, badToken)
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
   status, details := client.SearchDepositAuthorizationByCid(cfg.Endpoint, existing.ComputingId, goodToken)
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
   status, _ := client.SearchDepositAuthorizationByCid(cfg.Endpoint, goodId, badToken)
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
// export tests
//

func TestExportHappyDay(t *testing.T) {
   expected := http.StatusOK
   status, _ := client.ExportDepositAuthorization(cfg.Endpoint, goodToken)
   if status != expected {
      t.Fatalf("Expected %v, got %v\n", expected, status)
   }
}

func TestExportBadToken(t *testing.T) {
   expected := http.StatusForbidden
   status, _ := client.ExportDepositAuthorization(cfg.Endpoint, badToken)
   if status != expected {
      t.Fatalf("Expected %v, got %v\n", expected, status)
   }
}

//
// fulfill tests
//

func TestFulfillHappyDay(t *testing.T) {
   existing := getExistingAuthorization()
   if existing == nil {
      t.Fatalf("Unable to get existing authorization\n")
   }

   expected := http.StatusOK
   status := client.FulfillDepositAuthorization(cfg.Endpoint, existing.Id, goodDepositId, goodToken)
   if status != expected {
      t.Fatalf("Expected %v, got %v\n", expected, status)
   }
}

func TestFulfillEmptyId(t *testing.T) {
   expected := http.StatusBadRequest
   status := client.FulfillDepositAuthorization(cfg.Endpoint, empty, goodDepositId, goodToken)
   if status != expected {
      t.Fatalf("Expected %v, got %v\n", expected, status)
   }
}

func TestFulfillNotFoundId(t *testing.T) {
   expected := http.StatusNotFound
   status := client.FulfillDepositAuthorization(cfg.Endpoint, notFoundId, goodDepositId, goodToken)
   if status != expected {
      t.Fatalf("Expected %v, got %v\n", expected, status)
   }
}

func TestFulfillBadToken(t *testing.T) {
   expected := http.StatusForbidden
   status := client.FulfillDepositAuthorization(cfg.Endpoint, goodId, goodDepositId, badToken)
   if status != expected {
      t.Fatalf("Expected %v, got %v\n", expected, status)
   }
}

func TestFulfillBadDepositId(t *testing.T) {
   expected := http.StatusBadRequest
   status := client.FulfillDepositAuthorization(cfg.Endpoint, goodId, badDepositId, goodToken)
   if status != expected {
      t.Fatalf("Expected %v, got %v\n", expected, status)
   }
}

func getExistingAuthorization() *api.Authorization {

   status, details := client.SearchDepositAuthorizationById(cfg.Endpoint, "0", goodToken)
   if status == http.StatusOK {
      if len(details) != 0 {
         return details[0]
      }
   }

   return nil
}

func ensureValidAuthorizations(t *testing.T, details []*api.Authorization) {

   for _, e := range details {
      if emptyField(e.Id) ||
         emptyField(e.EmployeeId) ||
         emptyField(e.ComputingId) ||
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
         //emptyField( e.LibraId ) ||
         emptyField(e.Status) ||
         //emptyField( e.ApprovedAt ) ||
         //emptyField( e.AcceptedAt ) ||
         //emptyField( e.ExportedAt ) ||
         //emptyField( e.UpdatedAt ) ||
         emptyField(e.Status) {
         log.Printf("%t", e)
         t.Fatalf("Expected non-empty field but one is empty\n")
      }
   }
}

func emptyField(field string) bool {
   return len(strings.TrimSpace(field)) == 0
}

func loadConfig() TestConfig {

   data, err := ioutil.ReadFile("service_test.yml")
   if err != nil {
      log.Fatal(err)
   }

   var c TestConfig
   if err := yaml.Unmarshal(data, &c); err != nil {
      log.Fatal(err)
   }

   log.Printf("Test config; endpoint   [%s]\n", c.Endpoint)
   log.Printf("Test config; auth token [%s]\n", c.Token)

   return c
}
