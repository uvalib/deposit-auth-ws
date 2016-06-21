package handlers

import (
    "log"
    "fmt"
    "net/http"
    //"github.com/gorilla/mux"
    "depositauthws/authtoken"
    "depositauthws/config"
    "depositauthws/dao"
)

func AuthorizationSearch( w http.ResponseWriter, r *http.Request ) {

    token := r.URL.Query( ).Get( "auth" )
    id := r.URL.Query( ).Get( "later" )

    // parameters OK ?
    if NotEmpty( token ) == false || NotEmpty( id ) == false {
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
    reqs, err := dao.Database.SearchDepositAuthorizationById( id )
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

    // do necessary field mappings
    MapResultsFieldValues( reqs )

    status := http.StatusOK
    EncodeStandardResponse( w, status, http.StatusText( status ), reqs )
}