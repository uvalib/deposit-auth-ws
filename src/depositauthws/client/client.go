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

func GetDepositAuthorization( endpoint string, id string, token string ) ( int, [] * api.Authorization ) {

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

func SearchDepositAuthorization( endpoint string, id string, token string ) ( int, [] * api.Authorization ) {

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

func ImportDepositAuthorization( endpoint string, token string ) ( int, int ) {

    url := fmt.Sprintf( "%s/import?auth=%s", endpoint, token )
    //fmt.Printf( "%s\n", url )

    resp, body, errs := gorequest.New( ).
       SetDebug( debugHttp ).
       Post( url  ).
       Timeout( time.Duration( 5 ) * time.Second ).
       End( )

    if errs != nil {
        return http.StatusInternalServerError, 0
    }

    defer resp.Body.Close( )

    r := api.ImportExportResponse{ }
    err := json.Unmarshal( []byte( body ), &r )
    if err != nil {
        return http.StatusInternalServerError, 0
    }

    return resp.StatusCode, r.Count
}

func ExportDepositAuthorization( endpoint string, token string ) ( int, int ) {

    url := fmt.Sprintf( "%s/export?auth=%s", endpoint, token )
    //fmt.Printf( "%s\n", url )

    resp, body, errs := gorequest.New( ).
       SetDebug( debugHttp ).
       Post( url  ).
       Timeout( time.Duration( 5 ) * time.Second ).
       End( )

    if errs != nil {
        return http.StatusInternalServerError, 0
    }

    defer resp.Body.Close( )

    r := api.ImportExportResponse{ }
    err := json.Unmarshal( []byte( body ), &r )
    if err != nil {
        return http.StatusInternalServerError, 0
    }

    return resp.StatusCode, r.Count
}

func UpdateDepositAuthorization( endpoint string, reg api.Authorization, token string ) ( int, * api.Authorization ) {
    return http.StatusInternalServerError, nil
}

func DeleteDepositAuthorization( endpoint string, id string, token string ) int {

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