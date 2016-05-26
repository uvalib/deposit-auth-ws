package handlers

import (
    "log"
    "strings"
    "encoding/json"
    "net/http"
    "depositauthws/api"
    "depositauthws/mapper"
)

func EncodeStandardResponse( w http.ResponseWriter, status int, message string, details [] * api.Authorization ) {

    log.Printf( "Status: %d (%s)\n", status, message )
    jsonAttributes( w )
    coorsAttributes( w )
    w.WriteHeader( status )
    if err := json.NewEncoder(w).Encode( api.StandardResponse{ Status: status, Message: message, Details: details } ); err != nil {
        log.Fatal( err )
    }
}

func EncodeImportExportResponse( w http.ResponseWriter, status int, message string, count int ) {

    log.Printf( "Status: %d (%s)\n", status, message )
    jsonAttributes( w )
    w.WriteHeader( status )
    if err := json.NewEncoder(w).Encode( api.ImportExportResponse { Status: status, Message: message, Count: count } ); err != nil {
        log.Fatal( err )
    }
}

func EncodeHealthCheckResponse( w http.ResponseWriter, status int, message string ) {
    healthy := status == http.StatusOK
    jsonAttributes( w )
    w.WriteHeader( status )
    if err := json.NewEncoder(w).Encode( api.HealthCheckResponse { CheckType: api.HealthCheckResult{ Healthy: healthy, Message: message } } ); err != nil {
        log.Fatal( err )
    }
}

func jsonAttributes( w http.ResponseWriter ) {
    w.Header( ).Set( "Content-Type", "application/json; charset=UTF-8" )
}

func coorsAttributes( w http.ResponseWriter ) {
    w.Header( ).Set( "Access-Control-Allow-Origin", "*" )
    w.Header( ).Set( "Access-Control-Allow-Headers", "Content-Type" )
}

func NotEmpty( param string ) bool {
    return len( strings.TrimSpace( param ) ) != 0
}

// map any field values as necessary
func MapResultsFieldValues( details [] * api.Authorization ) {
    for _, d := range details {
        d.Department, _ = mapper.MapField( "department", d.Plan )
        d.Degree, _ = mapper.MapField( "degree", d.Degree )
    }
}