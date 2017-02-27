package handlers

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
    "depositauthws/authtoken"
    "depositauthws/config"
    "depositauthws/dao"
    "depositauthws/logger"
)

func AuthorizationDelete( w http.ResponseWriter, r *http.Request ) {

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
    if authtoken.Validate( config.Configuration.AuthTokenEndpoint, token, config.Configuration.Timeout ) == false {
        status := http.StatusForbidden
        EncodeStandardResponse( w, status, http.StatusText( status ), nil )
        return
    }

    // get the request details
    count, err := dao.Database.DeleteDepositAuthorizationById( id )
    if err != nil {
        logger.Log( fmt.Sprintf( "ERROR: %s\n", err.Error( ) ) )
        status := http.StatusInternalServerError
        EncodeStandardResponse( w, status,
            fmt.Sprintf( "%s (%s)", http.StatusText( status ), err ),
            nil )
        return
    }

    if count == 0 {
        status := http.StatusNotFound
        EncodeStandardResponse( w, status, http.StatusText( status ), nil )
        return
    }

    status := http.StatusOK
    EncodeStandardResponse( w, status, http.StatusText( status ), nil )
}