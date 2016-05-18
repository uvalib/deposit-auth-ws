package handlers

import (
    "log"
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
    "depositauthws/authtoken"
    "depositauthws/config"
    "depositauthws/dao"
)

func AuthorizationGet( w http.ResponseWriter, r *http.Request ) {

    vars := mux.Vars( r )
    id := vars[ "id" ]
    token := r.URL.Query( ).Get( "auth" )

    // parameters OK ?
    if NotEmpty( id ) == false || NotEmpty( token ) == false {
        status := http.StatusBadRequest
        EncodeStandardResponse( w, status, http.StatusText( status ), nil )
        return
    }

    // validate the token
    if authtoken.Validate( config.Configuration.AuthTokenEndpoint, token ) == false {
        status := http.StatusForbidden
        EncodeStandardResponse( w, status, http.StatusText( status ), nil )
        return
    }

    // get the request details
    reqs, err := dao.Database.GetDepositAuthorization( id )
    if err != nil {
        log.Println( err )
        status := http.StatusInternalServerError
        EncodeStandardResponse( w, status,
            fmt.Sprintf( "%s (%s)", http.StatusText( status ), err ),
            nil )
        return
    }

    if reqs == nil || len( reqs ) == 0 {
        status := http.StatusNotFound
        EncodeStandardResponse( w, status, http.StatusText( status ), nil )
        return
    }

    status := http.StatusOK
    EncodeStandardResponse( w, status, http.StatusText( status ), reqs )
}