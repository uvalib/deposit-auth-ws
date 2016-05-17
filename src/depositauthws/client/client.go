package client

import (
    "time"
    "fmt"
    "github.com/parnurzeal/gorequest"
    "net/http"
    "depositauthws/api"
    "encoding/json"
)

var debugHttp = false

func HealthCheck( endpoint string ) int {

    url := fmt.Sprintf( "%s/healthcheck", endpoint )
    //fmt.Printf( "%s\n", url )

    resp, _, errs := gorequest.New( ).
       SetDebug( debugHttp ).
       Get( url ).
       Timeout( time.Duration( 5 ) * time.Second ).
       End( )

    if errs != nil {
        return http.StatusInternalServerError
    }

    defer resp.Body.Close( )

    return resp.StatusCode
}

func Options( endpoint string ) ( int, * api.Options ) {

    url := fmt.Sprintf( "%s/options", endpoint )
    //fmt.Printf( "%s\n", url )

    resp, body, errs := gorequest.New( ).
       SetDebug( debugHttp ).
       Get( url ).
       Timeout( time.Duration( 5 ) * time.Second ).
       End( )

    if errs != nil {
        return http.StatusInternalServerError, nil
    }

    defer resp.Body.Close( )

    r := api.OptionsResponse{ }
    err := json.Unmarshal( []byte( body ), &r )
    if err != nil {
        return http.StatusInternalServerError, nil
    }

    return resp.StatusCode, r.Options
}

func GetDepositRequest( endpoint string, id string, token string ) ( int, [] * api.Registration ) {

    url := fmt.Sprintf( "%s/%s?auth=%s", endpoint, id, token )
    //fmt.Printf( "%s\n", url )

    resp, body, errs := gorequest.New( ).
       SetDebug( debugHttp ).
       Get( url  ).
       Timeout( time.Duration( 5 ) * time.Second ).
       End( )

    if errs != nil {
       return http.StatusInternalServerError, nil
    }

    defer resp.Body.Close( )

    r := api.StandardResponse{ }
    err := json.Unmarshal( []byte( body ), &r )
    if err != nil {
        return http.StatusInternalServerError, nil
    }

    return resp.StatusCode, r.Details
}

func SearchDepositRequest( endpoint string, id string, token string ) ( int, [] * api.Registration ) {

    url := fmt.Sprintf( "%s?auth=%s&later=%s", endpoint, token, id )
    //fmt.Printf( "%s\n", url )

    resp, body, errs := gorequest.New( ).
       SetDebug( debugHttp ).
       Get( url  ).
       Timeout( time.Duration( 5 ) * time.Second ).
       End( )

    if errs != nil {
        return http.StatusInternalServerError, nil
    }

    defer resp.Body.Close( )

    r := api.StandardResponse{ }
    err := json.Unmarshal( []byte( body ), &r )
    if err != nil {
        return http.StatusInternalServerError, nil
    }

    return resp.StatusCode, r.Details
}

func CreateDepositRequest( endpoint string, reg api.Registration, token string ) ( int, [] * api.Registration ) {

    url := fmt.Sprintf( "%s?auth=%s", endpoint, token )
    //fmt.Printf( "%s\n", url )

    resp, body, errs := gorequest.New( ).
       SetDebug( debugHttp ).
       Post( url  ).
       Send( reg ).
       Timeout( time.Duration( 5 ) * time.Second ).
       Set( "Content-Type", "application/json" ).
       End( )

    if errs != nil {
        return http.StatusInternalServerError, nil
    }

    defer resp.Body.Close( )

    r := api.StandardResponse{ }
    err := json.Unmarshal( []byte( body ), &r )
    if err != nil {
        return http.StatusInternalServerError, nil
    }

    //if resp.StatusCode == http.StatusOK {
    //    return http.StatusOK, r.Details[ 0 ]
    //}

    return resp.StatusCode, r.Details
}

func UpdateDepositRequest( endpoint string, reg api.Registration, token string ) ( int, * api.Registration ) {
    return http.StatusInternalServerError, nil
}

func DeleteDepositRequest( endpoint string, id string, token string ) int {

    url := fmt.Sprintf( "%s/%s?auth=%s", endpoint, id, token )
    //fmt.Printf( "%s\n", url )

    resp, body, errs := gorequest.New( ).
       SetDebug( debugHttp ).
       Delete( url  ).
       Timeout( time.Duration( 5 ) * time.Second ).
       End( )

    if errs != nil {
        return http.StatusInternalServerError
    }

    defer resp.Body.Close( )

    r := api.StandardResponse{ }
    err := json.Unmarshal( []byte( body ), &r )
    if err != nil {
        return http.StatusInternalServerError
    }

    return resp.StatusCode
}