package main

import (
    "io/ioutil"
    "log"
    "testing"
    "strings"
    "depositauthws/client"
    "depositauthws/api"
    "gopkg.in/yaml.v2"
    "net/http"
)

type TestConfig struct {
    Endpoint  string
    Token     string
}

var cfg = loadConfig( )

var goodId = "1"
var notFoundId = "x"
var goodToken = cfg.Token
var badToken = "badness"
var empty = " "

//
// healthcheck tests
//

func TestHealthCheck( t *testing.T ) {
    expected := http.StatusOK
    status := client.HealthCheck( cfg.Endpoint )
    if status != expected {
        t.Fatalf( "Expected %v, got %v\n", expected, status )
    }
}

//
// options tests
//

func TestOptionsHappyDay( t *testing.T ) {
    expected := http.StatusOK
    status, options := client.Options( cfg.Endpoint )
    if status != expected {
        t.Fatalf( "Expected %v, got %v\n", expected, status )
    }
    ensureValidOptions( t, options )
}

//
// get tests
//

func TestGetHappyDay( t *testing.T ) {
    expected := http.StatusOK
    status, details := client.GetDepositRequest( cfg.Endpoint, goodId, goodToken )
    if status != expected {
        t.Fatalf( "Expected %v, got %v\n", expected, status )
    }
    ensureValidRegistrations( t, details )
}

func TestGetEmptyId( t *testing.T ) {
    expected := http.StatusBadRequest
    status, _ := client.GetDepositRequest( cfg.Endpoint, empty, goodToken )
    if status != expected {
        t.Fatalf( "Expected %v, got %v\n", expected, status )
    }
}

func TestGetNotFoundId( t *testing.T ) {
    expected := http.StatusNotFound
    status, _ := client.GetDepositRequest( cfg.Endpoint, notFoundId, goodToken )
    if status != expected {
        t.Fatalf( "Expected %v, got %v\n", expected, status )
    }
}

func TestGetBadToken( t *testing.T ) {
    expected := http.StatusForbidden
    status, _ := client.GetDepositRequest( cfg.Endpoint, goodId, badToken )
    if status != expected {
        t.Fatalf( "Expected %v, got %v\n", expected, status )
    }
}

//
// search tests
//

func TestSearchHappyDay( t *testing.T ) {
    expected := http.StatusOK
    status, details := client.SearchDepositRequest( cfg.Endpoint, "0", goodToken )
    if status != expected {
        t.Fatalf( "Expected %v, got %v\n", expected, status )
    }
    ensureValidRegistrations( t, details )
}

func TestSearchEmptyId( t *testing.T ) {
    expected := http.StatusBadRequest
    status, _ := client.SearchDepositRequest( cfg.Endpoint, empty, goodToken )
    if status != expected {
        t.Fatalf( "Expected %v, got %v\n", expected, status )
    }
}

func TestSearchBadToken( t *testing.T ) {
    expected := http.StatusForbidden
    status, _ := client.SearchDepositRequest( cfg.Endpoint, goodId, badToken )
    if status != expected {
        t.Fatalf( "Expected %v, got %v\n", expected, status )
    }
}

//
// create tests
//

func TestSingleCreate( t *testing.T ) {
    reg := makeSingleRegistration( )
    expected := http.StatusOK
    status, details := client.CreateDepositRequest( cfg.Endpoint, reg, goodToken )
    if status != expected {
        t.Fatalf( "Expected %v, got %v\n", expected, status )
    }

    if details == nil || len( details ) != 1 {
        t.Fatalf( "Incomplete registration details returned" )
    }

    ensureValidRegistrations( t, details )
}

func TestMultiCreate( t *testing.T ) {
    reg := makeMultiRegistration( )
    expected := http.StatusOK
    status, details := client.CreateDepositRequest( cfg.Endpoint, reg, goodToken )
    if status != expected {
        t.Fatalf( "Expected %v, got %v\n", expected, status )
    }

    if details == nil || len( details ) != 2 {
        t.Fatalf( "Incomplete registration details returned" )
    }

    ensureValidRegistrations( t, details )
}

func TestCreateBadRegistration( t *testing.T ) {
    expected := http.StatusBadRequest
    status, _ := client.CreateDepositRequest( cfg.Endpoint, api.Registration{ }, goodToken )
    if status != expected {
        t.Fatalf( "Expected %v, got %v\n", expected, status )
    }
}

func TestCreateBadToken( t *testing.T ) {
    reg := makeSingleRegistration( )
    expected := http.StatusForbidden
    status, _ := client.CreateDepositRequest( cfg.Endpoint, reg, badToken )
    if status != expected {
        t.Fatalf( "Expected %v, got %v\n", expected, status )
    }
}

//
// update tests
//

//
// delete tests
//

func TestDeleteHappyDay( t *testing.T ) {
    newId := createNewReg( t )
    expected := http.StatusOK
    status := client.DeleteDepositRequest( cfg.Endpoint, newId, goodToken )
    if status != expected {
        t.Fatalf( "Expected %v, got %v\n", expected, status )
    }
}

func TestDeleteEmptyId( t *testing.T ) {
    expected := http.StatusBadRequest
    status := client.DeleteDepositRequest( cfg.Endpoint, empty, goodToken )
    if status != expected {
        t.Fatalf( "Expected %v, got %v\n", expected, status )
    }
}

func TestDeleteNotFoundId( t *testing.T ) {
    expected := http.StatusNotFound
    status := client.DeleteDepositRequest( cfg.Endpoint, notFoundId, goodToken )
    if status != expected {
        t.Fatalf( "Expected %v, got %v\n", expected, status )
    }
}

func TestDeleteBadToken( t *testing.T ) {
    expected := http.StatusForbidden
    status := client.DeleteDepositRequest( cfg.Endpoint, goodId, badToken )
    if status != expected {
        t.Fatalf( "Expected %v, got %v\n", expected, status )
    }
}

func ensureValidRegistrations( t *testing.T, details [] * api.Registration ) {

    for _, e := range details {
        if emptyField( e.Id ) ||
           emptyField( e.Requester ) ||
           emptyField( e.For ) ||
           emptyField( e.Department ) ||
           emptyField( e.Degree ) {
           //emptyField( e.RequestDate ) ||
           //emptyField( e.Status ) {
            t.Fatalf( "Expected non-empty field but one is empty\n" )
        }
    }
}

func ensureValidOptions( t *testing.T, options * api.Options ) {

    for _, f := range options.Department {
        if emptyField( f ) {
            t.Fatalf( "Expected non-empty department field but one is empty\n" )
        }
    }
    for _, f := range options.Degree {
        if emptyField( f ) {
            t.Fatalf( "Expected non-empty degree field but one is empty\n" )
        }
    }
}

func createNewReg( t *testing.T ) string {
    reg := makeSingleRegistration( )
    expected := http.StatusOK
    status, results := client.CreateDepositRequest( cfg.Endpoint, reg, goodToken )
    if status != expected {
        t.Fatalf( "Expected %v, got %v\n", expected, status )
    }
    if results == nil || len( results ) != 1 {
        t.Fatalf( "Incomplete registration details returned" )
    }

    return results[ 0 ].Id
}

func emptyField( field string ) bool {
    return len( strings.TrimSpace( field ) ) == 0
}

func makeSingleRegistration( ) api.Registration {
    return api.Registration{
        For: "dpg3k",
        Requester: "dpg3k",
        Department: "Engineering",
        Degree: "Ph.D" }
}

func makeMultiRegistration( ) api.Registration {
    return api.Registration{
        For: "dpg3k, tss6n",
        Requester: "dpg3k",
        Department: "Engineering",
        Degree: "Ph.D" }
}

func loadConfig( ) TestConfig {

    data, err := ioutil.ReadFile( "service_test.yml" )
    if err != nil {
        log.Fatal( err )
    }

    var c TestConfig
    if err := yaml.Unmarshal( data, &c ); err != nil {
        log.Fatal( err )
    }

    log.Printf( "Test config; endpoint   [%s]\n", c.Endpoint )
    log.Printf( "Test config; auth token [%s]\n", c.Token )

    return c
}